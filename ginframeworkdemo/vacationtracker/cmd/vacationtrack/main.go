package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"vacationtracker/employee"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "3000")
	}
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	// Built in middleware
	r.Use(gin.BasicAuth(gin.Accounts{"admin": "password"}))
	r.Use(MyErrorLogger)
	r.Use(gin.CustomRecovery(MyPanicRecovery))
	registerRoutes(r)

	r.Run()

}

func registerRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/employees")
	})

	r.GET("/employees", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", employee.GetAll())
	})

	r.GET("/employees/:employeeID", func(ctx *gin.Context) {
		empID := ctx.Param("employeeID")
		if emp, ok := tryToGetEmp(ctx, empID); ok {
			ctx.HTML(http.StatusOK, "employee.tmpl", emp)
		}
	})

	r.POST("/employees/:employeeID", func(ctx *gin.Context) {
		var timeoff employee.TimeOff
		err := ctx.ShouldBind(&timeoff)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		timeoff.Type = employee.TimeoffTypePTO
		timeoff.Status = employee.TimeoffStatusRequested
		empID := ctx.Param("employeeID")
		if emp, ok := tryToGetEmp(ctx, empID); ok {
			emp.TimeOff = append(emp.TimeOff, timeoff)
			ctx.Redirect(http.StatusFound, "/employees/"+empID)
		}

	})

	r.GET("/error", func(ctx *gin.Context) {
		err := &gin.Error{
			Err:  errors.New("something went wrong"),
			Type: gin.ErrorTypePublic | gin.ErrorTypeRender,
			Meta: "Error Page",
		}
		ctx.Error(err)
	})

	r.GET("/panic", func(ctx *gin.Context) {
		panic("our custom panic testing")
	})

	//adding custom middleware
	g := r.Group("/api/employees", Benchmark)
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, employee.GetAll())
	})

	g.GET("/:employeeID", func(ctx *gin.Context) {
		empID := ctx.Param("employeeID")
		if emp, ok := tryToGetEmp(ctx, empID); ok {
			ctx.JSON(http.StatusOK, emp)
		}
	})
	g.POST("/:employeeID", func(ctx *gin.Context) {
		var timeoff employee.TimeOff
		err := ctx.ShouldBindJSON(&timeoff)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		timeoff.Type = employee.TimeoffTypePTO
		timeoff.Status = employee.TimeoffStatusRequested
		empID := ctx.Param("employeeID")
		if emp, ok := tryToGetEmp(ctx, empID); ok {
			emp.TimeOff = append(emp.TimeOff, timeoff)
			ctx.JSON(http.StatusOK, *emp)
		}

	})

	r.Static("/public", "./public")
}

func tryToGetEmp(ctx *gin.Context, employeeID string) (*employee.Employee, bool) {
	empID, err := strconv.Atoi(employeeID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return nil, false
	}

	emp, err := employee.Get(empID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return nil, false
	}
	return emp, true

}

var Benchmark gin.HandlerFunc = func(ctx *gin.Context) {
	t := time.Now()
	ctx.Next()
	elapsed := time.Since(t)
	log.Print("Time taken in processing:", elapsed)

}

var MyErrorLogger gin.HandlerFunc = func(ctx *gin.Context) {
	ctx.Next()
	for _, err := range ctx.Errors {
		log.Print(map[string]interface{}{
			"Err":  err.Error(),
			"Type": err.Type,
			"Meta": err.Meta,
		})
	}
}

var MyPanicRecovery gin.RecoveryFunc = func(c *gin.Context, err interface{}) {
	log.Print("Recovery Process for panic with error :", err)
}
