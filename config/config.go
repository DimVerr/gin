package config

import (
	"fmt"
	"net/http"
	"os"
	"gogin/token"
	"gogin/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)




  func ConnectToDB() *gorm.DB{
	var err error
	errEnv := godotenv.Load()
	if errEnv != nil {
	  panic("Failed to load .env file")
	}
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD :=  os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	if DB_HOST == "" || DB_USER == "" || DB_PASSWORD == "" || DB_NAME =="" || DB_PORT =="" {
		fmt.Println("Please add your db data to .env file")
		os.Exit(1)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Credential{})
	return db
}

func ConnectToTestDB() *gorm.DB{
	var err error
	errEnv := godotenv.Load()
	if errEnv != nil {
	  panic("Failed to load .env file")
	}
	TEST_DB_HOST := os.Getenv("TEST_DB_HOST")
	TEST_DB_USER := os.Getenv("TEST_DB_USER")
	TEST_DB_PASSWORD :=  os.Getenv("TEST_DB_PASSWORD")
	TEST_DB_NAME := os.Getenv("TEST_DB_NAME")
	TEST_DB_PORT := os.Getenv("TEST_DB_PORT")
	if TEST_DB_HOST == "" || TEST_DB_USER == "" || TEST_DB_PASSWORD == "" || TEST_DB_NAME =="" || TEST_DB_PORT =="" {
		fmt.Println("Please add your db data to .env file")
		os.Exit(1)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",TEST_DB_HOST, TEST_DB_USER, TEST_DB_PASSWORD, TEST_DB_NAME, TEST_DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Credential{})
	return db
}

func LoginCheck(name string, password string) (string,error,uint) {
	db := ConnectToDB()
	var err error

	u := models.User{}

	err = db.Model(models.User{}).Where("name = ?", name).Take(&u).Error

	if err != nil {
		return "", err, 0
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err, 0
	}

	token,err := token.GenerateToken(u.ID)

	if err != nil {
		return "",err, 0
	}

	return token,nil, u.ID
	
}

func LoginCheckTest(name string, password string) (string,uint,error) {
	db := ConnectToTestDB()
	var err error

	u := models.User{}

	err = db.Model(models.User{}).Where("name = ?", name).Take(&u).Error

	if err != nil {
		return "", 0, err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", 0, err
	}

	token,err := token.GenerateToken(u.ID)

	if err != nil {
		return "", 0, err
	}

	return token, u.ID, nil
	
}

func VerifyPassword(password,hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}