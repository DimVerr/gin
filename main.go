package main

import (
	"gogin/handlers"
	"gogin/config"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gogin/docs"
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := gin.New()
	app.Use(CORSMiddleware())

	errEnv := godotenv.Load()
	if errEnv != nil {
	  panic("Failed to load .env file")
	}
	config.ConnectToDB()
	// url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition

	public := app.Group("/api")
	public.POST("/login", handlers.Login)
	public.POST("/signup", handlers.SignUp)

	
	protected:= app.Group("/api/creds")
	protected.Use(config.JwtAuthMiddleware())
	protected.GET("/all", handlers.GetAll)
	protected.GET("/:cred_id", handlers.GetOne)	
	protected.POST("/new", handlers.CreateCreds)
	protected.PUT("/:cred_id", handlers.UpdateCreds)
	protected.DELETE("/:cred_id", handlers.DeleteCreds)


	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err:= app.Run(); err != nil {
		log.Fatal(err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

