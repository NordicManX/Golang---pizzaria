package main

import (
	"github.com/gin-gonic/gin"
)

type Pizza struct {
	ID    int
	Nome  string
	preco float64
}

func main() {
	r := gin.Default()
	r.GET("/pizzas", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"pizzas": "Toscana, Calabresa sem cebola",
		})
	})

	// var pizzas = []Pizza{

	// 	{ID: 1, Nome: "Toscana", preco: 49.50},
	// 	{ID: 2, Nome: "Marguerita", preco: 79.50},
	// 	{ID: 3, Nome: "Atum com queijo", preco: 69.50},
	// }
	// fmt.Println(pizzas)
	r.Run()
}
