package handlers

import (
	_ "fiber/docs"
	"fiber/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Log in.
// @Description Log in the app.
// @Accept json
// @Produce json
// @Param user body utils.LoginRequest true "login"
// @Success 200 {object} utils.LoginResponse
// @Failure 400 {int} uint
// @Failure 404 {int} uint
// @Failure 500 {int} uint
// @Router /login [post]
func Login(c *gin.Context) {
    body := utils.User{}

    if err := c.ShouldBindJSON(&body); err != nil {
        c.AbortWithError(http.StatusUnprocessableEntity, err)
        return
    }
	user := utils.User{Name: body.Name , Password: body.Password}

	token, err, user_id := utils.LoginCheck(user.Name, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}


	c.JSON(http.StatusOK, gin.H{"token":token, "user_id": user_id})

}




// func Login(c *fiber.Ctx) error{

// 	loginRequest := new(utils.LoginRequest)
// 	if err := c.BodyParser(loginRequest); err != nil{
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": err.Error(),
// 		}) 
// 	}

// 	user, err := FindByCreds(loginRequest.Name, loginRequest.Password)
// 	if err != nil {
// 		return c.Status(401).JSON(fiber.Map{
// 			"error" : err.Error(),
// 		})
// 	}

// 	day := time.Hour * 24

// 	claims := jtoken.MapClaims{
// 		"ID": user.ID,
// 		"name": user.Name,
// 		"exp": time.Now().Add(day * 1).Unix(),
// 	}

// 	errEnv := godotenv.Load()
// 	if errEnv != nil {
// 	  panic("Failed to load .env file")
// 	}

// 	secret := os.Getenv("SECRET")

// 	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

// 	t, err := token.SignedString([]byte(secret))
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"error": err.Error(), 
// 		})
// 	}
// 	return c.Status(200).JSON(utils.LoginResponse{
// 	Token: t,
// 	ID: user.ID,
// 	})
// }

// func FindByCreds(name, password string) (*utils.User, error){
// 	var users utils.User
// 	db := utils.ConnectToDB()

// 	result := db.First(&users, "name = ?", name).Where("password = ?", password)
// 	if result.Error != nil{
// 		return nil, errors.New("user not found")
// 	}else  {
// 		return &utils.User{
// 			ID: users.ID,
// 			Name: users.Name,
// 			Password: users.Password, 
// 		}, nil
// 	}
// }