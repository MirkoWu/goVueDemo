package main

import (
	"goDemo/vuedemo/midware/vue"
	"goDemo/vuedemo/static"

	"github.com/gin-gonic/gin"
)

/*
*参考 https://www.itfanr.cc/2022/08/18/golang-embed-static-files/
 */
func main() {
	r := gin.Default()

	r.Use(vue.Serve("/", vue.EmbedFolder(static.Static, ".")))

	r.Run(":8080")

}
