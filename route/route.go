package route

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"result": false, "error": "Method Not Allowed"})
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"result": false, "error": "Endpoint Not Found"})
	})
	r.Use(AccessControl())
	r.Static("/css", "./templates/assets/css")
	r.Static("/js", "./templates/assets/js")
	r.Static("/fonts", "./templates/assets/fonts")
	r.Static("/img", "./templates/assets/img")

	r.GET("/", func(c *gin.Context) {
		r.LoadHTMLFiles("templates/views/index.html")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/dashboard", func(c *gin.Context) {
		r.LoadHTMLFiles("templates/views/user/index.html")
		c.HTML(http.StatusOK, "user/index.html", gin.H{
			"title": "This Is Dashboard",
		})
	})
	return r
}
