package main

import (
	"github.com/momotaroman112/sa-65-file/controller/payment"
	"github.com/momotaroman112/sa-65-file/entity"
	"github.com/momotaroman112/sa-65-file/middlewares"
	"github.com/gin-gonic/gin"


)
//Payment
func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{

			protected.GET("/employee/:id", controller.GetEmployee)
			protected.GET("/employees", controller.ListEmployee)
			protected.POST("/employees", controller.CreateEmployee)

			// registation Routes
			protected.GET("/bookings", controller.ListBooking)
			protected.GET("/booking/:id", controller.GetBooking)
			protected.GET("/booking/user/:id", controller.GetBookingByUser)


			protected.GET("/food_ordereds", controller.GetFoodOrdered)

	
			protected.GET("/payment_types", controller.GetPaymentType)

			protected.GET("/bills", controller.ListBill)
			protected.POST("/bills", controller.Createbill)

			// borrow list
			r.POST("/createborrowlist", controller.CreateBorrowList)
			r.GET("/borrowlist", controller.ListBorrow)
			r.GET("/equipment/:id", controller.GetEquipment)
			r.GET("/equipment", controller.ListEquipment)
			r.GET("/user/:id", controller.GetUser)
			r.GET("/user", controller.ListUser)
			r.POST("/createroom", controller.CreateRoom) //
			r.GET("/room", controller.ListRoom)
			r.GET("/type", controller.ListType)
		
			//Room
			r.GET("/building", controller.ListBuilding)
			r.GET("/serviceday", controller.ListServiceDay)
			r.GET("/period", controller.ListPeriod)
			r.POST("/login", controller.Login)


			//foodordered
	
			protected.GET("/users", controller.ListUsers)
			protected.GET("/user/:id", controller.GetUser)
			protected.POST("/users", controller.CreateUser)
			protected.PATCH("/users", controller.UpdateUser)
			protected.DELETE("/users/:id", controller.DeleteUser)

			protected.GET("/bookings", controller.ListBookings)

			protected.GET("/foodsets", controller.ListFoodSets)
			protected.GET("/foodset/:id", controller.GetFoodSet)
			protected.POST("/foodsets", controller.CreateFoodSet)
		

			protected.GET("/foodpayment_types", controller.ListFoodPaymentTypes)
			protected.GET("/foodpayment_type/:id", controller.GetFoodPaymentType)
			protected.POST("/foodpayment_types", controller.CreateFoodPaymentType)
			
	
			protected.GET("/foodordereds", controller.ListFoodOrdereds)
			protected.GET("/foodordered/:id", controller.GetFoodOrdered)
			protected.POST("/foodordereds", controller.CreateFoodOrdered)

			//equipment
			r.GET("/employees", controller.ListEmployees)
			r.PATCH("/employees", controller.UpdateEmployee)
			r.DELETE("/employees/:id", controller.DeleteEmployee)
			// Category Routes
			r.GET("/catagories", controller.ListCategory)
			r.GET("/catagory/:id", controller.GetCategory)
			r.POST("/catagories", controller.CreateCategory)
			r.PATCH("/catagories", controller.UpdateCategory)
			r.DELETE("/catagories/:id", controller.DeleteCategory)
			// Unit Routes
			r.GET("/units", controller.ListUnits)
			r.GET("/unit/:id", controller.GetUnit)
			r.POST("/units", controller.CreateUnit)
			r.PATCH("/units", controller.UpdateUnit)
			r.DELETE("units/:id", controller.DeleteUnit)
			// Equipment Routes
			r.GET("/equipments", controller.ListEquipments)
			r.GET("/equipment/:id", controller.GetEquipment)
			r.POST("/equipment", controller.CreateEquipment)
			r.PATCH("/equipments", controller.UpdateEquipment)
			r.DELETE("/equipments/:id", controller.DeleteEquipment)


			//booking
			protected.POST("/createbooking", controller.CreateBooking)
			protected.GET("/usage", controller.ListUsage)
			protected.GET("/member/:id", controller.GetMember)

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
