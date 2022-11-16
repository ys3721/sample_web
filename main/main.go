package main

import (
	"github.com/gin-gonic/gin"
	"sample_web/shellcaller"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Any("/:action/:param/:args", shellcaller.DoActionFunc)
	router.Run(":9527")
}
