package api

import (
	"real_project/api/handlers"
	log "real_project/pkg/logger"
	"real_project/storage"

	"github.com/gin-gonic/gin"
)

type Options struct {
	Storage storage.StorageI
	Log     log.Log
}

func Api(o Options) *gin.Engine {
	h := handlers.NewHandler(handlers.Handlers{o.Storage,o.Log})

	engine := gin.Default()

	api := engine.Group("/api")

	api.GET("/ping", h.Ping)

	own := api.Group("/own")
	{
		// own.POST("/log-in")
		// own.POST("/log-out")
		own.POST("/category", h.CreateCategory)
		// own.PUT("/category:id")
		// own.POST("/category")
		// own.PUT("/category:id")
		// own.DELETE("/category/:id")
	}

	// vw := api.Group("/vw")
	// {
	// 	vw.POST("/log-out")
	// 	vw.POST("/comment/:article_id")
	// }

	pb := api.Group("/pb")
	// {
	// 	pb.POST("/check_user")
	// 	pb.POST("/sign-up")
	// 	pb.POST("/login-in")
	pb.GET("/categories", h.GetCategoriesList)
	pb.GET("/categories/:id", h.GetCategory)
	// 	pb.GET("/sub-categories")
	// 	pb.GET("/sub-categories/:id")
	// 	pb.GET("/articles")
	// 	pb.GET("/articles/:id")
	// }

	return engine
}
