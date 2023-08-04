package handlers

import (
	"gogin/token"
	"gogin/models"
	"gogin/config"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// Create gofoc
// @Summary Create credentials.
// @Description Create credentials of user.
// @Accept json
// @Param cred body models.CredentialSwag true "creds"
// @Success 201 {object} models.Response
// @Failure 400 {int} uint
// @Failure 422 {int} uint
// @Security ApiKeyAuth
// @Router /api/creds/new [post]
func CreateCreds(c *gin.Context) {
	var response models.Response
	db := config.ConnectToDB()
	body := models.Credential{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest ,"Unable to extract id from token")
		return
	}

	creds:= models.Credential{UserID: body.UserID ,CredName: body.CredName, Domain: body.Domain, Login: body.Login, Password: body.Password}
	creds.UserID = uint(user_id)

	result := db.Create(&creds)
	if result.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return 
	}
	copier.Copy(&response, &creds)
	c.JSON(http.StatusCreated, response)
}

func CreateCredsTest(c *gin.Context) {
	var response models.Response
	db := config.ConnectToTestDB()
	body := models.Credential{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest ,"Unable to extract id from token")
		return
	}

	creds:= models.Credential{UserID: body.UserID ,CredName: body.CredName, Domain: body.Domain, Login: body.Login, Password: body.Password}
	creds.UserID = uint(user_id)

	result := db.Create(&creds)
	if result.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return 
	}
	copier.Copy(&response, &creds)
	c.JSON(http.StatusCreated, response)
}