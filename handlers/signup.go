package handlers

import (
	"fiber/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// SignUp godoc
// @Summary Signup.
// @Description Log in the app.
// @Accept json
// @Param user body utils.LoginRequest true "signup"
// @Success 200 {object} utils.User
// @Failure 409 {int} uint
// @Failure 422 {int} uint
// @Router /signup [post]
func SignUp(c *gin.Context) {
	db := utils.ConnectToDB()
    body := utils.User{}

    if err := c.ShouldBindJSON(&body); err != nil {
        c.AbortWithError(http.StatusUnprocessableEntity, err)
        return
    }
	user := utils.User{Name: body.Name , Password: body.Password}

    if result := db.Create(&user); result.Error != nil {
    	c.AbortWithError(http.StatusConflict, result.Error)
		return
	}
	var response utils.UserResponse
	copier.Copy(&response, &user)
    c.JSON(http.StatusCreated, &response)
}

// func SignUp(c *fiber.Ctx) error{
// 	db := utils.ConnectToDB()
// 	user := new(utils.User)

// 	if err := c.BodyParser(&user); err != nil {
// 		return c.Status(422).SendString(err.Error())
// 	}

// 	if result := db.Create(user); result.Error != nil {
// 		return c.Status(409).SendString("User with this name is already existing")
// 	}
// 	var response utils.UserResponse
// 	copier.Copy(&response, &user)
// 	return c.Status(200).JSON(response)
// }