package main

import (
	"github.com/momotaroman112/sa-65-file/controller"
	"github.com/momotaroman112/sa-65-file/entity"
	"github.com/momotaroman112/sa-65-file/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{

			// // User Routes
			// protected.GET("/users/:id", controller.GetUser)
			// protected.POST("/users", controller.CreateUser)

			protected.GET("/employee/:id", controller.GetEmployee)
			protected.GET("/employees", controller.ListEmployee)
			protected.POST("/employees", controller.CreateEmployee)

			// registation Routes
			protected.GET("/bookings", controller.ListBooking)
			protected.GET("/booking/:id", controller.GetBooking)
			protected.GET("/booking/user/:id", controller.GetBookingByUser)

			// foodordered Routes
			protected.GET("/food_ordereds", controller.GetFoodOrdered)

			// payment_types Routes
			protected.GET("/payment_types", controller.GetPaymentType)

			// bills Routes
			protected.GET("/bills", controller.ListBill)
			protected.POST("/bills", controller.Createbill)

		}
	}

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
