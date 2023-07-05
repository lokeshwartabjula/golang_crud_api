package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
}

var users = []User{
	{
		ID:        "0F8JIqi4zwvb77FGz6Wt",
		FirstName: "Heinz-Georg",
		Email:     "heinz-georg.fiedler@example.com",
	},
	{
		ID:        "0P6E1d4nr0L1ntW8cjGU",
		Email:     "katie.hughes@example.com",
		FirstName: "Katie",
	},
	{
		ID:        "1Lkk06cOUCkiAsUXFLMN",
		FirstName: "Vetle",
		Email:     "vetle.aasland@example.com",
	},
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Users retrieved",
		"success": true,
		"users":   users,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var updateUser User
	err := c.BindJSON(&updateUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse request body",
			"success": false,
		})
		return
	}

	for i, user := range users {
		if user.ID == id {
			users[i].Email = updateUser.Email
			users[i].FirstName = updateUser.FirstName
			c.JSON(http.StatusOK, gin.H{
				"message": "User updated",
				"success": true,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "User not found",
		"success": false,
	})
}

func AddUser(c *gin.Context) {
	var newUser User
	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse request body",
			"success": false,
		})
		return
	}

	newUser.ID = generateUserID()
	users = append(users, newUser)

	c.JSON(http.StatusOK, gin.H{
		"message": "User added",
		"success": true,
	})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"user":    user,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "User not found",
		"success": false,
	})
}

func generateUserID() string {
	rand.Seed(time.Now().UnixNano())

	// Generate a random 8-character ID
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	id := make([]byte, 8)
	for i := range id {
		id[i] = charset[rand.Intn(len(charset))]
	}

	return string(id)
}

func main() {
	r := gin.Default()

	r.GET("/users", GetUsers)
	r.PUT("/update/:id", UpdateUser)
	r.POST("/add", AddUser)
	r.GET("/user/:id", GetUserByID)

	r.Run(":8080")
}
