package main

import (
	"dataxperience-server-side/controllers"
	"dataxperience-server-side/models"
	"net/http"
	"time"

	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	models.ConnectDB()
	corsMiddleware := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"}, // You might want to specify the actual origins here.
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders:   []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    })

    r.Use(func(c *gin.Context) {
        corsMiddleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
            c.Next()
        })).ServeHTTP(c.Writer, c.Request)
    })


	// Middleware JWT
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("your_secret_key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.Authors); ok {
				return jwt.MapClaims{
					"id":    v.ID,
					"email": v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			id, ok := claims["id"].(float64)
			if !ok {
				return nil
			}
			return &models.Authors{
				Model: gorm.Model{
					ID: uint(id),
				},
				Email: claims["email"].(string),
			}
		},
		
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"code": code, "message": message})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		panic("Error creating JWT middleware")
	}

	publicAPI := r.Group("/api")
	{
		publicAPI.POST("/author", controllers.CreateAuthor)
		publicAPI.GET("/news", controllers.IndexNews)
		publicAPI.PUT("/news/:id", controllers.UpdateNews)
		publicAPI.POST("/login", controllers.LoginHandler)
	}

	protectedAPI := r.Group("/api")
	protectedAPI.Use(authMiddleware.MiddlewareFunc())
	{
		protectedAPI.POST("/news", controllers.CreateNews)
		protectedAPI.GET("/news/:id", controllers.ShowNews)
		protectedAPI.DELETE("/news/:id", controllers.DeleteNews)

		protectedAPI.GET("/author/:id", controllers.ShowAuthor)
		protectedAPI.PUT("/author/:id", controllers.UpdateAuthor)
		protectedAPI.DELETE("/author/:id", controllers.DeleteAuthor)

		protectedAPI.GET("/categories/", controllers.IndexCategory)
		protectedAPI.POST("/categories/", controllers.CreateCategory)
		protectedAPI.GET("/categories/:id", controllers.ShowCategories)
		protectedAPI.PUT("/categories/:id", controllers.UpdateCategory)
		protectedAPI.DELETE("/categories/:id", controllers.DeleteCategory)

		protectedAPI.POST("/client", controllers.CreateClient)
		protectedAPI.GET("/client/:id", controllers.ShowClient)
		protectedAPI.PUT("/client/:id", controllers.UpdateClient)
		protectedAPI.DELETE("/client/:id", controllers.DeleteClient)
	}

	r.Run(":5000")
}
