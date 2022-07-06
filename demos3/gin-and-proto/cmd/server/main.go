package main

import (
	// "gin-and-proto/goproto"

	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

// type loginParam struct {
// 	Username string `json:"username"  form:"username"  binding:"required"`
// 	Password string `json:"password"  form:"password"  binding:"required"`
// }

type loginResp struct {
	ErrorCode int
	UserID    int64
	Session   string
}

func Login(c *gin.Context) {

	var resp loginResp

	resp.ErrorCode = 2
	resp.UserID = 1
	resp.Session = "session"

	test := &Student{
		Name: "geektutu",
		Male: true,
		// Scores: []int32{98, 85, 88},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println(data)

	c.ProtoBuf(http.StatusOK, test)

	// c.IndentedJSON(http.StatusOK, resp)

}
func main() {
	r := gin.Default()
	r.GET("/login", Login)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
