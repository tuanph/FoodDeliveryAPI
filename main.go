package main

import (
	"FoodDelivery/modules/restaurant/restauranttransport/ginrestaurant"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// connecto to DB
	dsn := "root:blogic@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(db))
		restaurants.GET("", func(c *gin.Context) {
			var data []Restaurant
			newDB := db
			if err := newDB.Table("restaurants").Find(&data).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
			}
			c.JSON(http.StatusOK, data)
		})

	}
	return r.Run()
}

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}
