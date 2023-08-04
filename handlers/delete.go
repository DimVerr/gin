package handlers

import (
	"gogin/token"
	"gogin/config"
	"gogin/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// DeleteCreds godoc
// @Summary Delete credentials by id.
// @Description Delete credentials by id.
// @Param cred_id path uint true "cred_id"
// @Success 200  {int} uint
// @Failure 404 {int} uint
// @Security ApiKeyAuth
// @Router /api/creds/{cred_id} [delete]
func DeleteCreds(c *gin.Context) {
	var creds models.Credential
	db := config.ConnectToDB()
	
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest ,"Unable to extract id from token")
		return
	}
	cred_id := c.Param("cred_id")

	result := db.Where("user_id = ? AND id = ?", user_id , cred_id).Delete(&creds)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, "No credentials were found")
		return
	}

	c.JSON(http.StatusOK, "Credentials were deleted successfully")
}

func DeleteCredsTest(c *gin.Context) {
	var creds models.Credential
	db := config.ConnectToTestDB()
	
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest ,"Unable to extract id from token")
		return
	}
	cred_id := c.Param("cred_id")

	result := db.Where("user_id = ? AND id = ?", user_id , cred_id).Delete(&creds)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, "No credentials were found")
		return
	}

	c.JSON(http.StatusOK, "Credentials were deleted successfully")
}