package handlers

import (
	"gogin/token"
	"gogin/config"
	"gogin/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// GetOne godoc
// @Summary Get credentials by id.
// @Description Get credentials by id.
// @Param cred_id path uint true "cred_id"
// @Success 200 {object} models.Response
// @Failure 404 {int} uint
// @Security ApiKeyAuth
// @Router /api/creds/{cred_id} [get]
func GetOne(c *gin.Context) {
	var cred models.Credential
	db := config.ConnectToDB()
	
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest ,"Unable to extract id from token")
		return
	}
	cred_id := c.Param("cred_id")

	result := db.Where("user_id = ? AND id = ?", user_id , cred_id).Find(&cred)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, "No credentials were found")
		return
	}
	var response models.Response
	copier.Copy(&response, &cred)
	
	c.JSON(http.StatusOK, response)
}

func GetOneTest(c *gin.Context) {
	var cred models.Credential
	db := config.ConnectToTestDB()
	
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest ,"Unable to extract id from token")
		return
	}
	cred_id := c.Param("cred_id")

	result := db.Where("user_id = ? AND id = ?", user_id , cred_id).Find(&cred)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, "No credentials were found")
		return
	}
	var response models.Response
	copier.Copy(&response, &cred)
	
	c.JSON(http.StatusOK, response)
}
