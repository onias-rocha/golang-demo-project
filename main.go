package main

import (
	"fmt"
	"golang-softplayer-apply/http"
	. "golang-softplayer-apply/model"
	"time"
)

var familiaRocha = []Person{
	{
		ID:        "1",
		Name:      "Teste Silva",
		Gender:    "Fluid",
		Email:     "teste.teste@teste.com",
		BirthDate: time.Date(1995, 03, 30, 0, 0, 0, 0, time.UTC),
		City:      "Testelandia",
		Country:   "Brazil",
		Cpf:       "22255588878",
	},
}

func main() {
	app := http.NewApp()

	app.UrlMapping()

	if err := app.Run("localhost:8080"); err != nil {
		fmt.Println("deu erro ", err)
	}

	//gin.Engine{}
	//router := gin.Default()
	//router.GET("/ping", ping)
	//router.GET("/person", getPeople)
	//router.GET("/person/:id", getPersonById)
	//router.DELETE("/person/:id", deletePersonById)
	//router.POST("/person", createPerson)
	//router.Run("localhost:8080")
}

//func ping(c *gin.Context) {
//	c.IndentedJSON(http.StatusOK, "pong")
//}
//
//func getPeople(c *gin.Context) {
//	c.IndentedJSON(http.StatusOK, familiaRocha)
//}
//
//func createPerson(c *gin.Context) {
//	var newPerson person
//
//	if err := c.BindJSON(&newPerson); err != nil {
//		return
//	}
//
//	familiaRocha = append(familiaRocha, newPerson)
//	c.IndentedJSON(http.StatusCreated, newPerson)
//}
//
//func getPersonById(c *gin.Context) {
//	var id = c.Param("id")
//
//	for _, a := range familiaRocha {
//		if a.ID == id {
//			c.IndentedJSON(http.StatusOK, a)
//			return
//		}
//	}
//	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
//}
//
//func deletePersonById(c *gin.Context) {
//	id := c.Param("id")
//
//	for i, a := range familiaRocha {
//		if a.ID == id {
//			familiaRocha = append(familiaRocha[:i], familiaRocha[i+1:]...)
//			c.IndentedJSON(http.StatusOK, gin.H{"message": "item removed"})
//			return
//		}
//	}
//
//	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
//}
