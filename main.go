package main

import (
	"FoodDelivery/component"
	"FoodDelivery/component/uploadprovider"
	"FoodDelivery/middleware"
	"FoodDelivery/modules/restaurant/restauranttransport/ginrestaurant"
	"FoodDelivery/modules/upload/uploadtransport/ginupload"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// connecto to DB
	dsn := "root:blogic@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"

	s3BucketName := "food-delivery-golang"
	s3Region := "ap-southeast-1"
	s3APIKey := "AKIAYTGE2NRUZ4V5L62E"
	s3SecretKey := "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"
	s3Domain := ""
	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	if err := runService(db, s3Provider); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider) error { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	appCtx := component.NewAppContext(db, upProvider)
	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping pong",
		})
	})

	r.POST("/upload", ginupload.Upload(appCtx))

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.PUT("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

	}
	return r.Run()
}

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}
