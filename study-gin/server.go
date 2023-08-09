package main

import (
	"io"
	"net/http"
	"os"
	"study-gin/controller"
	"study-gin/middlewares"
	"study-gin/service"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func SetupLogOutput() {
	f, _ := os.Create("gin.Log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	SetupLogOutput()

	server := gin.New()

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")
	
	server.Use(
		gin.Recovery(),
		middlewares.Logger(),
		middlewares.BasicAuth(),
		gindump.Dump(),
	)

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Hello World",
			})
		})

		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		// apiRoutes.POST("/videos", func(ctx *gin.Context) {
		// 	ctx.JSON(200, videoController.Save(ctx))
		// })

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"Message": "omaigah"})
			}
		})
	}

	
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}


	server.Run(":8080")
}
