package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed public/*
var f embed.FS

func main() {
	router := gin.Default()
	// Basic Gin Test
	/* 	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!!")
	}) */
	// Serving Files from Gin
	/* router.StaticFile("/", "./public/index.html")
	router.Static("/public", "./public")
	router.StaticFS("/fs", http.FileSystem(http.FS(f))) */

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
		ctx.String(http.StatusOK, "New Request POSTed Successfully.")
	})

	adminGp := router.Group("/admin")
	adminGp.GET("/roles", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "This is Admin Roles page...")
	})

	adminGp.GET("/policies", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "This is Admin Policies page...")
	})

	log.Fatal(router.Run(":3000"))
}
