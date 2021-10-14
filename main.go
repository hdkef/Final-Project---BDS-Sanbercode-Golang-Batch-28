package main

import (
	"bloggo/config"
	"bloggo/controllers"
	docs "bloggo/docs"
	"bloggo/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Bloggo API
// @version 1.0
// @description Bloggo is a REST Api for blogging

// @contact.name Hadekha Erfadila Fitra
// @contact.email hdkef11@gmail.com

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.apikey AuthToken
// @in header
// @name authorization

// @x-extension-openapi {"example": "value on a json format"}

// gin-swagger middleware
// swagger embed files

func main() {
	docs.SwaggerInfo.BasePath = "/api/v1"

	db, err := config.DBConn()
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := gin.Default()

	//middleware
	r.Use(middleware.DBMiddleware(db))
	r.Use(middleware.JWTMiddleware)

	authCtl := &controllers.AuthCtl{}
	articleCtl := &controllers.ArticleCtl{}
	commentCtl := &controllers.CommentCtl{}
	draftCtl := &controllers.DraftCtl{}
	inboxCtl := &controllers.InboxCtl{}
	mediaCtl := &controllers.MediaCtl{}
	ratingCtl := &controllers.RatingCtl{}
	userCtl := &controllers.UserCtl{}

	v1 := r.Group("/api/v1")
	{
		articles := v1.Group("/articles")
		{
			articles.GET("", articleCtl.GetAll())
			articles.GET(":id", articleCtl.GetOne())
			articles.POST("", articleCtl.Post())
			articles.PUT(":id", articleCtl.Put())
			articles.DELETE(":id", articleCtl.Delete())
		}
		comments := v1.Group("/comments")
		{
			comments.GET(":article-id", commentCtl.GetAll())
			comments.POST(":article-id", commentCtl.Post())
			comments.PUT(":id", commentCtl.Put())
			comments.DELETE(":id", commentCtl.Delete())
		}
		drafts := v1.Group("/drafts")
		{
			drafts.GET("", draftCtl.GetAll())
			drafts.GET(":id", draftCtl.GetOne())
			drafts.POST("", draftCtl.Post())
			drafts.PUT(":id", draftCtl.Put())
			drafts.DELETE(":id", draftCtl.Delete())
		}
		inboxes := v1.Group("/inboxes")
		{
			inboxes.GET(":receiver-id", inboxCtl.GetAll())
			inboxes.POST(":receiver-id", inboxCtl.Post())
			inboxes.PUT(":id", inboxCtl.Put())
			inboxes.DELETE(":id", inboxCtl.Delete())
		}
		media := v1.Group("/media")
		{
			media.GET("", mediaCtl.GetAll())
			media.POST("", mediaCtl.Post())
			media.PUT(":id", mediaCtl.Put())
			media.DELETE(":id", mediaCtl.Delete())
		}
		rating := v1.Group("/ratings")
		{
			rating.GET(":article-id", ratingCtl.GetAll())
			rating.GET(":article-id/sum", ratingCtl.GetSum())
			rating.POST(":article-id", ratingCtl.Post())
			rating.PUT(":id", ratingCtl.Put())
			rating.DELETE(":id", ratingCtl.Delete())
		}
		auth := v1.Group("/auth")
		{
			auth.POST("login", authCtl.Login())
			auth.POST("register", authCtl.Register())
			auth.POST("ch-pwd", authCtl.ChangePWD())
		}
		user := v1.Group("/users")
		{
			user.GET("", userCtl.GetAll())
			user.GET(":id", userCtl.GetOne())
			user.GET(":id/public", userCtl.GetOnePublic())
			user.POST("", userCtl.Post())
			user.PUT(":id", userCtl.Put())
			user.DELETE(":id", userCtl.Delete())
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
