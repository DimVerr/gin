package main

import (
	"bytes"
	"encoding/json"
	"fiber/handlers"
	"fiber/utils"

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
	// teardown()

	os.Exit(exitCode)
}
func setup() {
	db := utils.ConnectToDB()
	db.AutoMigrate(&utils.User{}, utils.Credential{})
	
}

// func teardown() {
// 	db := utils.ConnectToDB()
// 	migrator := db.Migrator()
// 	migrator.DropTable(&utils.User{})
// 	migrator.DropTable(&utils.Credential{})
// }

func router() *gin.Engine{
	app := gin.Default()

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
	user := utils.LoginRequest{
			Name: "yemiwebby",
			Password: "test",
	}

	writer := makeRequest("POST", "/api/login", user, false)
	var response utils.LoginResponse
	json.Unmarshal(writer.Body.Bytes(), &response)
	return response.Token

}

func TestRegister(t *testing.T) {
	newUser := utils.LoginRequest{
			Name: "yemiwebby",
			Password: "test",
	}
	writer := makeRequest("POST", "/api/signup", newUser, false)
	assert.Equal(t, http.StatusCreated, writer.Code)
	var response utils.UserResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.Equal(t, utils.UserResponse{ID: response.ID, Name: response.Name, Password: response.Password}, response)
}

func TestLogin(t *testing.T) {
	user := utils.LoginRequest{
			Name: "yemiwebby",
			Password: "test",
	}

	writer := makeRequest("POST", "/api/login", user, false)

	assert.Equal(t, http.StatusOK, writer.Code)

	var response utils.LoginResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	assert.Equal(t, utils.LoginResponse{Token: response.Token, ID: response.ID}, response)
}

func TestAddEntry(t *testing.T) {
	newEntry := utils.Credential{
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
	writer := makeRequest("GET", "/api/creds/7", nil, true)
	assert.Equal(t, http.StatusOK, writer.Code)
}
func TestUpdateEntry(t *testing.T) {
	newEntry := utils.Credential{
		CredName: "update",
		Domain: "update",
		Login: "update",
		Password: "update",
	}

	writer := makeRequest("PUT", "/api/creds/7", newEntry, true)
	assert.Equal(t, http.StatusOK, writer.Code)
	var response utils.Response
	json.Unmarshal(writer.Body.Bytes(), &response)
	assert.Equal(t, utils.Response{ID: response.ID, CredName: response.CredName, Domain: response.CredName,Login: response.Login, Password: response.Password}, response)

}

func TestDelete(t *testing.T) {
	writer := makeRequest("DELETE", "/api/creds/7", nil, true)
	assert.Equal(t, http.StatusOK, writer.Code)
}