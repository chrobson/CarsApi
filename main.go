package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type car struct {
	ID    string  `json:"id"`
	Color string  `json:"color"`
	Price float64 `json:"price"`
}

var cars = []car{
	{ID: "1", Color: "Blue", Price: 25355.25},
	{ID: "2", Color: "Red", Price: 55122.99},
	{ID: "3", Color: "Green", Price: 62332.33},
}

func getCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cars)
}

func getCarById(c *gin.Context) {
	id := c.Param("id")

	for _, car := range cars {
		if car.ID == id {
			c.IndentedJSON(http.StatusOK, car)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car with id " + id + " not found"})
}

func postCars(c *gin.Context) {
	var newCar car
	if err := c.BindJSON(&newCar); err != nil {
		return
	}
	cars = append(cars, newCar)
	c.IndentedJSON(http.StatusCreated, newCar)
}

func main() {
	router := gin.Default()
	router.GET("/cars", getCars)
	router.GET("/cars/:id", getCarById)
	router.POST("/cars", postCars)
	router.Run("localhost:8080")
}
