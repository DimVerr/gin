package handlers

import (
	"gogin/token"
	"gogin/config"
	"gogin/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// UpdateCreds godoc
// @Summary Update Credentials.
// @Description Update Credentials of your user.
// @Param cred_id path uint true "cred_id"
// @Param creds body models.CredentialSwag true "update request"
// @Accept json
// @Success 200 {object} models.ResponseUpdate
// @Failure 404 {int} uint
// @Failure 422 {int} uint
// @Security ApiKeyAuth
// @Router /api/creds/{cred_id} [put]
func UpdateCreds(c *gin.Context) {
	db := config.ConnectToDB()

	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest ,"Unable to extract id from token")
		return
	}

	cred_id := c.Param("cred_id")

	body := models.Credential{}
	if err := c.ShouldBindJSON(&body); err != nil{
		c.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	creds:= models.Credential{CredName: body.CredName, Domain: body.Domain, Login: body.Login, Password: body.Password}
	creds.UserID = uint(user_id)

	result := db.Where("id = ? AND user_id = ?", cred_id , user_id).Updates(&creds)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, "No credentials were found")
	}

	var response models.ResponseUpdate
	copier.Copy(&response, &creds)
	c.JSON(http.StatusOK , response)
}

func UpdateCredsTest(c *gin.Context) {
	db := config.ConnectToTestDB()

	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest ,"Unable to extract id from token")
		return
	}

	cred_id := c.Param("cred_id")

	body := models.Credential{}
	if err := c.ShouldBindJSON(&body); err != nil{
		c.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}
	creds:= models.Credential{CredName: body.CredName, Domain: body.Domain, Login: body.Login, Password: body.Password}
	creds.UserID = uint(user_id)

	result := db.Where("id = ? AND user_id = ?", cred_id , user_id).Updates(&creds)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, "No credentials were found")
	}

	var response models.Response
	copier.Copy(&response, &creds)
	c.JSON(http.StatusOK ,creds)
}