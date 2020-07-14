package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/arammikayelyan/bookstore_users-api/domain/users"
	"github.com/arammikayelyan/bookstore_users-api/services"

	"github.com/gin-gonic/gin"
)

// CreateUser func
func CreateUser(c *gin.Context) {
	var user users.User

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GetUser func
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement")
}
