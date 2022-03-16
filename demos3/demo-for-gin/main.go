package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BirthdayTime time.Time

func (u *BirthdayTime) MarshalJSON() ([]byte, error) {
	t := time.Time(*u)
	return t.MarshalJSON()
}

func (u *BirthdayTime) UnmarshalJSON(data []byte) error {
	t := time.Time{}
	err := t.UnmarshalJSON(data)
	if err != nil {
		return err
	}

	*u = BirthdayTime(t)
	return nil
}

type Person struct {
	User string `form:"user" json:"user" xml:"user"  binding:"required"`
	// Password string `form:"password" json:"password" xml:"password" binding:"required"`
	Birthday BirthdayTime `form:"birthday" json:"birthday"`
}

func GAddPersion(c *gin.Context) {

	person := Person{}
	err := c.ShouldBind(&person)
	if err == nil {
		log.Println(person.User)
		// log.Println(person.Address)
		log.Println(person.Birthday)
		// log.Println(person.CreateTime)
		// log.Println(person.UnixTime)
	} else {
		log.Println(err.Error())
	}
	c.String(http.StatusOK, "Succeed")

	// firstname := c.DefaultQuery("firstname", "Guest")
	// lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

	// c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func main() {
	r := gin.Default()
	r.Static("/", "./html")
	r.POST("/addpersion", GAddPersion)
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
