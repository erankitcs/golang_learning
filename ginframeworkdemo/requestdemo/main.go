package main

import (
	"embed"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

//go:embed public/*
var f embed.FS

type TimeofRequest struct {
	// use - for no validation
	Date   time.Time `json:"data" form:"data" binding:"required,future" time_format:"2006-01-02"`
	Amount float64   `json:"amount" form:"amount" binding:"required,gt=0"`
}

var ValidationFuture validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		return date.After(time.Now())
	}
	return true
}

func main() {
	router := gin.Default()
	// Basic Gin Test
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!!")
	})
	// Serving Files from Gin
	router.StaticFile("/", "./public/index.html")
	router.Static("/public", "./public")
	router.StaticFS("/fs", http.FileSystem(http.FS(f)))

	router.GET("/employee", func(ctx *gin.Context) {
		ctx.File("public/employee.html")
	})

	router.GET("/employee/:username/*rest", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"username": ctx.Param("username"),
			"rest":     ctx.Param("rest"),
		})
	})

	router.POST("/employee", func(ctx *gin.Context) {
		//ctx.String(http.StatusOK, "New Request POSTed Successfully.")

		//Normal way to access and handle form data
		/* date := ctx.PostForm("date")
		amount := ctx.PostForm("amount")
		username := ctx.DefaultPostForm("username", "ankit")

		ctx.IndentedJSON(http.StatusOK, gin.H{
			"date":     date,
			"amount":   amount,
			"username": username,
		}) */

		// Form data binding using Struc
		var timeofRequest TimeofRequest
		if err := ctx.ShouldBind(&timeofRequest); err == nil {
			ctx.JSON(http.StatusOK, timeofRequest)
		} else {
			ctx.String(http.StatusInternalServerError, err.Error())
		}

	})

	//Binding Validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("future", ValidationFuture)
	}

	apiGrp := router.Group("/api")
	// JSON data binding using Struc
	//localhost:3000/api/timeoff
	// body - {
	//"data" : "2022-12-03T00:00:00Z",
	//"amount" : 100
	// }
	apiGrp.POST("/timeoff", func(ctx *gin.Context) {
		var timeofRequest TimeofRequest
		if err := ctx.ShouldBindJSON(&timeofRequest); err == nil {
			ctx.JSON(http.StatusOK, timeofRequest)
		} else {
			ctx.String(http.StatusInternalServerError, err.Error())
		}
	})

	adminGp := router.Group("/admin")
	adminGp.GET("/roles", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "This is Admin Roles page...")
	})

	adminGp.GET("/policies", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "This is Admin Policies page...")
	})

	//http://localhost:3000/request/mydata
	router.GET("/request/*rest", func(ctx *gin.Context) {
		url := ctx.Request.URL.String()
		headers := ctx.Request.Header
		cookies := ctx.Request.Cookies()

		ctx.IndentedJSON(http.StatusOK, gin.H{
			"url":     url,
			"header":  headers,
			"cookies": cookies,
		},
		)
	})

	//http://localhost:3000/query?username=ankit&year=2021&month=Jun&month=Jul
	//http://localhost:3000/query?username=ankit&month=Jun
	router.GET("/query", func(ctx *gin.Context) {

		username := ctx.Query("username")
		year := ctx.DefaultQuery("year", strconv.Itoa(time.Now().Year()))
		months := ctx.QueryArray("month")

		ctx.IndentedJSON(http.StatusOK, gin.H{
			"username": username,
			"year":     year,
			"months":   months,
		},
		)
	})

	log.Fatal(router.Run(":3000"))
}
