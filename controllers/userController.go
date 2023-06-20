package controllers

import (
	"net/http"

	"github.com/drsimplegraffiti/ref/initializers"
	"github.com/drsimplegraffiti/ref/models"
	"github.com/drsimplegraffiti/ref/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context){ // c is gin's context, which is a wrapper around http.Request and http.ResponseWriter
	var body struct{ // we didnt use type User struct because we dont want to expose the password
		Email string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are empty"})
		return
}

//check if user exists
var count int64
initializers.DB.Model(&models.User{}).Where("email = ?", body.Email).Count(&count)
if count > 0 {
	c.JSON(http.StatusConflict, gin.H{
		"error": "User already exists",
		"status": http.StatusConflict,
	})
	return
}

//hash password
hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 8)
if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Error hashing password"})
	return
}

//create user
user := models.User{
	Email: body.Email,
	Password: string(hashedPassword),
}
result := initializers.DB.Create(&user)
if result.Error != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Error creating user"})
	return
}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}


func Login(c *gin.Context) {
	var body struct{
		Email string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are empty"})
		return
	}

	//check if user exists
	var user models.User
	result := initializers.DB.Where("email = ?", body.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	}

	//check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	//generate token
	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}


func GetCurrentUser(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var user models.User
	initializers.DB.Where("id = ?", userID).First(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {
	var body struct{
		FirstName string
		LastName string
		Age int
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are empty"})
		return
	}

	// check if the body parameters are empty
	if body.FirstName == "" || body.LastName == "" || body.Age == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are empty"})
		return
	}

	userID := c.MustGet("userID").(uint)
	var user models.User
	initializers.DB.Where("id = ?", userID).First(&user)

	//check if user exists
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	//update user
	user.FirstName = body.FirstName
	user.LastName = body.LastName
	user.Age = body.Age
	initializers.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"user": user})

}