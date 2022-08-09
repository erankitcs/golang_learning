package main

import (
	"net/http"
	"os"
	"strconv"
	"vacationtracker/employee"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "3000")
	}
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
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
