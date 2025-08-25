package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Table  , Ingredient ,Soup

// Table , Function ทำให้ Return โต๊ะใหญ่ เล็ก ถ้า Total <= 2 เล็ก >= 3 ใหญ่ -->GetTable Input เป็น TableID
// Student struct
type Table struct {
	ID    string `json:"id"`
	Total int    `json:"total"`
}

type Ingredient struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Type   string `json:"type"`
}

// In-memory database (ในโปรเจคจริงใช้ database)
var TableInShop = []Table{
	{ID: "1", Total: 4},
	{ID: "2", Total: 2},
	{ID: "3", Total: 1},
	{ID: "4", Total: 5},
	{ID: "5", Total: 3},
}

var IngreInStock = []Ingredient{
	{ID: "1", Name: "Instance Noodle", Amount: 32, Type: "Noodle"},
	{ID: "2", Name: "Cheese Tofu", Amount: 42, Type: "Tofu"},
	{ID: "3", Name: "Egg Tofu", Amount: 45, Type: "Tofu"},
	{ID: "4", Name: "Vermicelli", Amount: 26, Type: "Noodle"},
	{ID: "5", Name: "Sliced Pork Belly", Amount: 86, Type: "Meat"},
	{ID: "6", Name: "Beef Slices", Amount: 78, Type: "Meat"},
	{ID: "7", Name: "Marinated pork", Amount: 95, Type: "Meat"},
	{ID: "8", Name: "Corn", Amount: 43, Type: "Vegetable"},
	{ID: "9", Name: "Wakame Seaweed", Amount: 67, Type: "Vegetable"},
	{ID: "10", Name: "Jade Noodle", Amount: 44, Type: "Noodle"},
	{ID: "11", Name: "Cabbage", Amount: 58, Type: "Vegetable"},
	{ID: "12", Name: "Bok Choy", Amount: 21, Type: "Vegetable"},
}

func getTable(c *gin.Context) {
	TableId := c.Query("ID")
	if TableId != "" {
		filter := []Table{}
		var tableSize int = 0
		for _, table := range TableInShop {
			if fmt.Sprint(table.ID) == TableId {
				filter = append(filter, table)
				tableSize = table.Total
			}
		}
		if tableSize > 0 && tableSize < 3 {
			c.JSON(http.StatusOK, gin.H{"Table": "Small"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Table": "Big"})
		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, TableInShop)
}

func getIngredients(c *gin.Context) {
	IngreID := c.Query("ID")
	if IngreID != "" {
		filter := []Ingredient{}
		for _, Ingre := range IngreInStock {
			if fmt.Sprint(Ingre.ID) == IngreID {
				filter = append(filter, Ingre)
			}
		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, IngreInStock)
}

func main() {
	r := gin.Default()
	r.GET("/Soup", func(c *gin.Context) {
		c.JSON(200, gin.H{"Soup": "Mala"})
		c.JSON(200, gin.H{"Soup": "Shabu"})
		c.JSON(200, gin.H{"Soup": "Tomyum"})

	})
	api := r.Group("/api/v1")
	{
		api.GET("/MalaTable", getTable)
		api.GET("/MalaIngre", getIngredients)
		fmt.Println()
	}

	r.Run(":8080")
}
