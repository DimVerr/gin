package handlers

import (
	"fiber/token"
	"fiber/utils"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/copier"
)

// Create gofoc
// @Summary Create credentials.
// @Description Create credentials of user.
// @Accept json
// @Param cred body utils.CredentialSwag true "creds"
// @Success 201 {object} utils.Credential
// @Failure 400 {int} uint
// @Failure 422 {int} uint
// @Security ApiKeyAuth
// @Router /creds [post]
func CreateCreds(c *gin.Context) {
	var response utils.Response
	db := utils.ConnectToDB()
	body := utils.Credential{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest ,"Unable to extract id from token")
		return
	}

	creds:= utils.Credential{UserID: body.UserID ,CredName: body.CredName, Domain: body.Domain, Login: body.Login, Password: body.Password}
	creds.UserID = uint(user_id)

	result := db.Create(&creds)
	if result.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return 
	}
	copier.Copy(&response, &creds)
	c.JSON(http.StatusCreated, response)
}

// func CreateCreds(c *fiber.Ctx) error{
// 	var response utils.Response
// 	db := utils.ConnectToDB()
// 	creds := new(utils.Credential)
// 	if err := c.BodyParser(creds); err != nil {
// 		return c.Status(422).SendString(err.Error())
// 	}

// 	user := c.Locals("user").(*jtoken.Token)
// 	claims := user.Claims.(jtoken.MapClaims)
// 	user_id := claims["ID"].(float64)

// 	creds.UserID = uint(user_id)
// 	result := db.Create(&creds)
// 	if result.RowsAffected == 0 {
// 		return c.SendStatus(400)
// 	}
// 	copier.Copy(&response, &creds)
// 	return c.Status(201).JSON(response)
// }
