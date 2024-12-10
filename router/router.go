package router

import (
	"github.com/gin-gonic/gin"
	"studyGo/controllers"
	"studyGo/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)

	}

	api := r.Group("/api")
	api.GET("/auth", controllers.TokenVerify)
	api.GET("/exchangeRates", controllers.GetExchangeRates)
	api.GET("/qcs", controllers.QCSLottery)
	api.GET("/get_wishes", controllers.GetWish)
	api.Use(middleware.AuthMiddle())
	{

		api.POST("/exchangeRates", controllers.CreateExchangeRate)
		api.POST("/new_wish", controllers.NewWish)
		api.DELETE("/delete_wish/:id", controllers.DeleteWish)
		api.PUT("/wish_check/:id", controllers.WishCheck)
		api.PUT("/update_wish/:id", controllers.WishUpdate)
	}
	api.POST("/new_articles", controllers.CreatArticle)
	api.POST("/get_articles", controllers.GetArticles)
	return r

}
