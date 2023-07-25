package handlers

import (
	"fiber/token"
	"fiber/utils"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/copier"
)

// GetOne godoc
// @Summary Get credentials by id.
// @Description Get credentials by id.
// @Param cred_id path uint true "cred_id"
// @Success 200 {object} utils.Credential
// @Failure 404 {int} uint
// @Security ApiKeyAuth
// @Router /creds/{cred_id} [get]
func GetOne(c *gin.Context) {
	var cred utils.Credential
	db := utils.ConnectToDB()
	
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
	var response utils.Response
	copier.Copy(&response, &cred)
	
	c.JSON(http.StatusOK, response)
}

// func GetOne(c *fiber.Ctx) error{
// 	var cred utils.Credential
// 	db := utils.ConnectToDB()
// 	user := c.Locals("user").(*jtoken.Token)
// 	claims := user.Claims.(jtoken.MapClaims)
// 	user_id := claims["ID"].(float64)

// 	cred_id := c.Params("cred_id")

// 	result := db.Where("user_id = ? AND id = ?", user_id , cred_id).Find(&cred)
// 	if result.RowsAffected == 0 {
// 		return c.SendStatus(404)
// 	}
// 	var response utils.Response
// 	copier.Copy(&response, &cred)
	
// 	return c.Status(200).JSON(response)
// }