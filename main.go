package main

import (
	"encoding/json"
	"fmt"

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
	fmt.Println(ids[0].UserID)
	c.JSON(200, gin.H{
		"payload": ids,
	})
}
