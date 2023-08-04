package models

import "gorm.io/gorm"


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


type ResponseUpdate struct {
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