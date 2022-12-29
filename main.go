package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var todos = []todo{
	{Id: "1", Name: "Sam", Age: 29},
	{Id: "2", Name: "Kam", Age: 20},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func getIndex(context *gin.Context) {
	jsonData := []byte(`{"message":"Welcome Tanvir Islam Streame"}`)
	context.Data(http.StatusOK, gin.MIMEJSON, jsonData)
}

func main() {
	router := gin.Default()
	router.GET("/", getIndex)
	router.GET("/todos", getTodos)
	router.Run("localhost:6700")
}
