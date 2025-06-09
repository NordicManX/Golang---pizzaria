package main

import (
	"pizzaria/models"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/pizzas", getPizzas)
	r.Run()
}
func getPizzas(c *gin.Context) {
	var pizzas = []models.Pizza{

		{ID: 1, Nome: "Toscana", Preco: 49.50},
		{ID: 2, Nome: "Marguerita", Preco: 79.50},
		{ID: 3, Nome: "Atum com queijo", Preco: 69.50},
	}

	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}
