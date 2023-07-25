package main

import (
	_ "fiber/docs"
	"fiber/handlers"
	"fiber/utils"
	"log"

	// swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	app := gin.Default()

	errEnv := godotenv.Load()
	if errEnv != nil {
	  panic("Failed to load .env file")
	}
	// secret := os.Getenv("SECRET")
	// jwt := utils.NewAuthMiddleware(secret)
	
	utils.ConnectToDB()
	// app.GET("/swagger/*", swagger.HandlerDefault) // default
	public := app.Group("/api")
	public.POST("/login", handlers.Login)
	public.POST("/signup", handlers.SignUp)

	protected:= app.Group("/api/creds")
	protected.Use(utils.JwtAuthMiddleware())
	protected.GET("/all", handlers.GetAll)
	protected.GET("/:cred_id", handlers.GetOne)	
	protected.POST("/new", handlers.CreateCreds)
	protected.PUT("/:cred_id", handlers.UpdateCreds)
	protected.DELETE("/:cred_id", handlers.DeleteCreds)



	if err:= app.Run("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}

// func main() {
// 	app := fiber.New()

// 	errEnv := godotenv.Load()
// 	if errEnv != nil {
// 	  panic("Failed to load .env file")
// 	}
// 	secret := os.Getenv("SECRET")
// 	jwt := utils.NewAuthMiddleware(secret)
	
// 	utils.ConnectToDB()
// 	app.Get("/swagger/*", swagger.HandlerDefault) // default
// 	app.Post("/login", handlers.Login)
// 	app.Post("/signup", handlers.SignUp)
// 	app.Get("/creds/all", jwt, handlers.GetAll)
// 	app.Get("/creds/:cred_id", jwt, handlers.GetOne)	
// 	app.Post("/creds", jwt, handlers.CreateCreds)
// 	app.Put("/creds/:cred_id", jwt, handlers.UpdateCreds)
// 	app.Delete("/creds/:cred_id", jwt, handlers.DeleteCreds)



// 	if err:= app.Listen(":8080"); err != nil {
// 		log.Fatal(err)
// 	}
// }

