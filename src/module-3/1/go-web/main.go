package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id        int     `json:"id"`
	Name      string  `json:"name" binding:"required"`
	LastName  string  `json:"lastName" binding:"required"`
	Email     string  `json:"email" binding:"required"`
	Age       int     `json:"age" binding:"required"`
	Height    float64 `json:"height" binding:"required"`
	IsActive  bool    `json:"isActive"`
	CreatedAt string  `json:"createdAt"`
}

func main() {
	r := gin.Default()

	users := r.Group("/users")
	{
		users.GET("/", GetAll)
		users.GET("/:id", GetById)
		users.POST("/", Create)

	}

	r.Run()
}

func GetAll(c *gin.Context) {
	file, _ := ioutil.ReadFile("./users.json")
	data := []User{}
	_ = json.Unmarshal([]byte(file), &data)

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func GetById(c *gin.Context) {
	file, _ := ioutil.ReadFile("./users.json")
	data := []User{}
	_ = json.Unmarshal([]byte(file), &data)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	result := User{}
	for _, user := range data {
		if user.Id == id {
			result = user
			break
		}
	}

	if result.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

var newUsers []User

func Create(c *gin.Context) {
	if token := c.Request.Header.Get("Authorization"); token != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
		})
		return
	}

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user.Id = len(newUsers)
	user.IsActive = true
	user.CreatedAt = time.Now().UTC().Format(time.RFC3339Nano)

	newUsers = append(newUsers, user)

	c.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}
