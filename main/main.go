package main

import (
	"github.com/gin-gonic/gin"
	"sample_web/shellcaller"
)


func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Any("/:action/:param", shellcaller.DoActionFunc)
	router.Run(":9527")
}