package handlers

import (
	_ "gogin/docs"
	"gogin/config"
	"gogin/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Log in.
// @Description Log in the app.
// @Accept json
// @Produce json
// @Param user body models.LoginRequest true "login"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {int} uint
// @Failure 404 {int} uint
// @Failure 500 {int} uint
// @Router /api/login [post]
func Login(c *gin.Context) {
    body := models.User{}

    if err := c.ShouldBindJSON(&body); err != nil {
        c.AbortWithError(http.StatusUnprocessableEntity, err)
        return
    }
	user := models.User{Name: body.Name , Password: body.Password}

	token, err, user_id := config.LoginCheck(user.Name, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}


	c.JSON(http.StatusOK, gin.H{"token":token, "user_id": user_id})

}

func LoginTest(c *gin.Context) {
    body := models.User{}

    if err := c.ShouldBindJSON(&body); err != nil {
        c.AbortWithError(http.StatusUnprocessableEntity, err)
        return
    }
	user := models.User{Name: body.Name , Password: body.Password}

	token, user_id, err := config.LoginCheckTest(user.Name, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}


	c.JSON(http.StatusOK, gin.H{"token":token, "user_id": user_id})

}