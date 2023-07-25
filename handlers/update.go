package handlers

import (
	"fiber/token"
	"fiber/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// UpdateCreds godoc
// @Summary Update Credentials.
// @Description Update Credentials of your user.
// @Param creds body utils.UpdateRequest true "update request"
// @Accept json
// @Success 200 {object} utils.Credential
// @Failure 404 {int} uint
// @Failure 422 {int} uint
// @Security ApiKeyAuth
// @Router /creds/{cred_id} [put]
func UpdateCreds(c *gin.Context) {
	db := utils.ConnectToDB()

	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest ,"Unable to extract id from token")
		return
	}

	cred_id := c.Param("cred_id")

	body := utils.Credential{}
	if err := c.ShouldBindJSON(&body); err != nil{
		c.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	creds:= utils.Credential{CredName: body.CredName, Domain: body.Domain, Login: body.Login, Password: body.Password}
	creds.UserID = uint(user_id)

	result := db.Where("user_id = ? AND id = ?", user_id , cred_id).Updates(&creds)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, "No credentials were found")
	}

	var response utils.Response
	copier.Copy(&response, &creds)
	c.JSON(http.StatusOK ,response)
}