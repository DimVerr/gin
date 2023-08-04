package main

import (
	"bytes"
	"encoding/json"
	"gogin/handlers"
	"gogin/config"
	"gogin/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	setup()
	exitCode := m.Run()
	teardown()

	os.Exit(exitCode)
}
func setup() {
	db := config.ConnectToTestDB()
	db.AutoMigrate(&models.User{}, models.Credential{})
	
}

func teardown() {
	db := config.ConnectToTestDB()
	migrator := db.Migrator()
	migrator.DropTable(&models.User{})
	migrator.DropTable(&models.Credential{})
}

func router() *gin.Engine{
	app := gin.Default()

	public := app.Group("/api")
	public.POST("/login", handlers.LoginTest)
	public.POST("/signup", handlers.SignUpTest)

	protected:= app.Group("/api/creds")
	protected.Use(config.JwtAuthMiddleware())
	protected.GET("/all", handlers.GetAllTest)
	protected.GET("/:cred_id", handlers.GetOneTest)	
	protected.POST("/new", handlers.CreateCredsTest)
	protected.PUT("/:cred_id", handlers.UpdateCredsTest)
	protected.DELETE("/:cred_id", handlers.DeleteCredsTest)

	return app
}

func makeRequest(method, url string, body interface{}, isAuthenticatedRequest bool) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if isAuthenticatedRequest {
			request.Header.Add("Authorization", "Bearer "+bearerToken())
	}
	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}

func bearerToken() string {
	user := models.LoginRequest{
			Name: "yemiwebby",
			Password: "test",
	}

	writer := makeRequest("POST", "/api/login", user, false)
	var response models.LoginResponse
	json.Unmarshal(writer.Body.Bytes(), &response)
	return response.Token

}

func TestRegister(t *testing.T) {
	newUser := models.LoginRequest{
			Name: "yemiwebby",
			Password: "test",
	}
	writer := makeRequest("POST", "/api/signup", newUser, false)
	assert.Equal(t, http.StatusCreated, writer.Code)
	var response models.UserResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.Equal(t, models.UserResponse{ID: response.ID, Name: response.Name, Password: response.Password}, response)
}

func TestLogin(t *testing.T) {
	user := models.LoginRequest{
			Name: "yemiwebby",
			Password: "test",
	}

	writer := makeRequest("POST", "/api/login", user, false)

	assert.Equal(t, http.StatusOK, writer.Code)

	var response models.LoginResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.Equal(t, models.LoginResponse{Token: response.Token, ID: response.ID}, response)
}

func TestAddEntry(t *testing.T) {
	newEntry := models.Credential{
			CredName: "test",
			Domain: "test",
			Login: "test",
			Password: "test",
	}
	writer := makeRequest("POST", "/api/creds/new", newEntry, true)
	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestGetAllEntries(t *testing.T) {
	writer := makeRequest("GET", "/api/creds/all", nil, true)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestGetOneEntry(t *testing.T) {
	writer := makeRequest("GET", "/api/creds/1", nil, true)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestUpdateEntry(t *testing.T) {
	newEntry := models.Credential{
		CredName: "update",
		Domain: "update",
		Login: "update",
		Password: "update",
	}

	writer := makeRequest("PUT", "/api/creds/1", newEntry, true)
	assert.Equal(t, http.StatusOK, writer.Code)
	var response models.Response
	json.Unmarshal(writer.Body.Bytes(), &response)
	assert.Equal(t, models.Response{ID: response.ID, CredName: response.CredName, Domain: response.CredName,Login: response.Login, Password: response.Password}, response)

}

func TestDelete(t *testing.T) {
	writer := makeRequest("DELETE", "/api/creds/1", nil, true)
	assert.Equal(t, http.StatusOK, writer.Code)
}