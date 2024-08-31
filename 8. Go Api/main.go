package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)
type Person struct {
	Name string `form:"name"`
	Address string `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}
func main () {
	r := gin.Default();
	r.GET("/",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message":"hello go api",
		})
	})
	r.GET("/welcome/:id", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
		id := c.Param("id")
		fmt.Println(id)
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	r.GET("/testing",startPage)
	r.Run();
}

func startPage (ctx *gin.Context) {
	var person Person

	if ctx.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}
	ctx.String(200, "Success")
}