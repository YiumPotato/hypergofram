package routes

import (
	"net/http"

	"github.com/YiumPotato/hypergofram/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default() // Logger and Recovery middlewares already attached

	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Serve static files as a File System, easy access @ /assets
	r.StaticFS("/assets", http.Dir("./static"))

	// Serve these favicon files directly
	r.StaticFile("/favicon.ico", "./static/favicon/favicon.ico")
	r.StaticFile("/android-chrome-192x192.png", "./static/favicon/android-chrome-192x192.png")
	r.StaticFile("/android-chrome-512x512.png", "./static/favicon/android-chrome-512x512.png")

	// use TEMPL renderer in gin
	r.HTMLRender = &TemplRender{}

	// Routes
	r.GET("/", controllers.Home)
	r.GET("/about", controllers.About)

	return r
}
