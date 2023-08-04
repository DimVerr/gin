package handlers

import (
	"gogin/token"
	"gogin/config"
	"gogin/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// GetAll godoc
// @Summary Get all credentials.
// @Description Get all credentials.
// @Success 200 {array} models.Response
// @Failure 404 {int} uint
// @Security ApiKeyAuth
// @Router /api/creds/all [get]
func GetAll(c *gin.Context) {
	var creds []models.Credential

	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest ,"Unable to extract id from token")
		return
	}

	db := config.ConnectToDB()

	result := db.Where("user_id=?", user_id).Find(&creds)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, "No credentials were found")
		return 
	}
	var response []models.Response
	copier.Copy(&response, &creds)
	
	c.JSON(http.StatusOK, response)
}

func GetAllTest(c *gin.Context) {
	var creds []models.Credential

	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest ,"Unable to extract id from token")
		return
	}

	db := config.ConnectToTestDB()

	result := db.Where("user_id=?", user_id).Find(&creds)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, "No credentials were found")
		return 
	}
	var response []models.Response
	copier.Copy(&response, &creds)
	
	c.JSON(http.StatusOK, response)
}