package main

import (
	"../controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())

	router.LoadHTMLGlob("templates/index.tmpl.html")
	router.Static("/css", "templates/css")
	router.GET("/", controller.ShowHome)

	router.Run()
