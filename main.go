package main

import (
	"encoding/json"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Identity struct {
	Email     string `json:"email"`
	UserID    string `json:"user_id"`
	SessionId string `json:"session_id"`
}

func main() {

	r := gin.Default()
	r.POST("/test", handleTest)
	r.Run()
}
func handleTest(c *gin.Context) {
	fmt.Println(c.PostForm("Grant_Type"))
	// fmt.Println(c.PostFormArray("Payload"))
	var ids []Identity
	json.Unmarshal([]byte(c.PostForm("Payload")), &ids)
	mm, _ := json.Marshal(ids[0])
	var cc jwt.MapClaims
	json.Unmarshal(mm, &cc)
	fmt.Println(cc)

	fmt.Println(ids[0].Email)
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
