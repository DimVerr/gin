package utils

import (
	"fmt"
	"net/http"
	"os"

	"fiber/token"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex"`
	Password string
	Credential []Credential 
  }
  
  type UserResponse struct {
	ID uint
	Name string
	Password string
  }
  
  type Credential struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	UserID      uint `gorm:"foreignKey"`
	CredName     string
	Domain       string
	Login        string
	Password     string
  }

  type CredentialSwag struct {
	CredName     string
	Domain       string
	Login        string
	Password     string
  }

  type Response struct {
	ID uint
	CredName string
	Domain string
	Login string
	Password string
  }

  type UpdateRequest struct {
	ID uint 
	CredName string
	Domain string
	Login string
	Password string
  }

  type LoginRequest struct {
	Name string `json:"name"`
	Password string `json:"password"`
  }

  type LoginResponse struct {
	ID uint `json:"user_id"`
	Token string `json:"token"`
  }

  type HTTPError struct {
	Status  string
	Message string
	}

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

	db.AutoMigrate(&User{}, &Credential{})
	return db
}

func LoginCheck(name string, password string) (string,error,uint) {
	db := ConnectToDB()
	var err error

	u := User{}

	err = db.Model(User{}).Where("name = ?", name).Take(&u).Error

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