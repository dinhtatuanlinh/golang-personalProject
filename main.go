package main

// "github.com/gin-gonic/gin"

// func main() {

// }

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//the return type of the func is part of its overall type definition - specify string as it's return type to comply with example you have above

type Boo func(str string)

func pr(str string){
	fmt.Println(str)
}
func bus(b string, c string, pr Boo) {
	fmt.Println(b)
	pr(c)
}
func main() {

	bus("linh", "ta", pr)
		router := gin.Default()

	router.GET("/",
		func(c *gin.Context) {
			c.JSON(200, gin.H{ "message": "pong",})
		})
	router.GET("/name/:name",
		func(c *gin.Context) {
			name := c.Param("name")
			// http://localhost:3000/name/linh?firstname=abc
			firstname := c.DefaultQuery("firstname", "Guest")
			c.JSON(200, gin.H{ "message": name, "query": firstname})
		})

 	router.Run(":3000")

}