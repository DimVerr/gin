package handlers

import (
	"gogin/config"
	"gogin/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// SignUp godoc
// @Summary Signup.
// @Description Log in the app.
// @Accept json
// @Param user body models.LoginRequest true "signup"
// @Success 200 {object} models.UserResponse
// @Failure 409 {int} uint
// @Failure 422 {int} uint
// @Router /api/signup [post]
func SignUp(c *gin.Context) {
	db := config.ConnectToDB()
    body := models.User{}

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusUnprocessableEntity, "Name is already in use")
        return
    }
	user := models.User{Name: body.Name , Password: body.Password}

    if result := db.Create(&user); result.Error != nil {
    	c.AbortWithError(http.StatusConflict, result.Error)
		return
	}
	var response models.UserResponse
	copier.Copy(&response, &user)
    c.JSON(http.StatusCreated, &response)
}


func SignUpTest(c *gin.Context) {
	db := config.ConnectToTestDB()
    body := models.User{}

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusUnprocessableEntity, "Name is already in use")
        return
    }
	user := models.User{Name: body.Name , Password: body.Password}

    if result := db.Create(&user); result.Error != nil {
    	c.AbortWithError(http.StatusConflict, result.Error)
		return
	}
	var response models.UserResponse
	copier.Copy(&response, &user)
    c.JSON(http.StatusCreated, &response)
}
