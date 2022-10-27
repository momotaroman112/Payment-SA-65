package entity

import (
	"time"

	"gorm.io/gorm"
)

//การกำหนด schema
type User struct {
	gorm.Model
	Name string
	Email		string			`gorm:"uniqueIndex"`
	Password     string

	Bookings	[]Booking	 `gorm:"foreignKey:MemberID"`
}

type Employee struct {
	gorm.Model
	Name string
	//StudentCode
	Email		string			`gorm:"uniqueIndex"`
	Password     string

	Bills	[]Bill	 `gorm:"foreignKey:EmployeeID"`
}

type FoodSet struct {
	gorm.Model //เอาไว้ hold พวก Pk (FoodSetID)

	Name   string
	Detail string
	Price  int

	FoodOrderedFoodSets []FoodOrderedFoodSet `gorm:"foreignKey:FoodSetID"`
}

type FoodOrderedFoodSet struct {
	gorm.Model //เอาไว้ hold พวก Pk (FF_id)

	// FoodSets       []FoodSet gorm:"foreignKey:FF_id"
	FoodOrderedID *uint
	FoodOrdered   FoodOrdered `gorm:"references:ID"`

	FoodSetID *uint
	FoodSet   FoodSet `gorm:"references:ID"`

	Quantity uint
}

type FoodOrdered struct {
	gorm.Model //เอาไว้ hold พวก Pk (FoodOrderedID)
	FoodTime   time.Time

	BookingID *uint
	Booking   Booking

	TotalPrice uint

	FoodOrderedFoodSets []FoodOrderedFoodSet `gorm:"foreignKey:FoodOrderedID"`

}


type PaymentType struct {
	gorm.Model
	Name string
	Bill []Bill `gorm:"foreignKey:PaymentTypeID"`
}


type Type struct {
	gorm.Model
	Name  string
	Price int
	Rooms []Room `gorm:"foreignKey:TypeID"`
}

type Room struct {
	gorm.Model
	Number string
	Name   string

	//TypeID  ทำหน้าที่เป็น FK
	TypeID *uint
	Type   Type
	//EmployeeID ทำหน้าที่เป็น FK
	EmployeeID *uint
	Employee   Employee
}

type Booking struct {
	gorm.Model
	BookingTimeStart  time.Time
	BookingTimeStop time.Time

	RoomID 			*uint
	Room			Room

	MemberID     	*uint
	Member      	User

	FoodOrdereds []FoodOrdered `gorm:"foreignKey:BookingID"`
}


type Bill struct {
	gorm.Model
	
	BillTime time.Time

	EmployeeID     *uint
	Employee        Employee

	PaymentTypeID *uint
	PaymentType   PaymentType

	BookingID *uint    `gorm:"uniqueIndex"`
	Booking    Booking `gorm:"constraint:OnDelete:CASCADE"` //belong to ลบใบลงทะเบียน บิลหาย

	// FoodOrderedID *uint
	// FoodOrdered   FoodOrdered `gorm:"foreignKey:BillID"`

	TotalPrice uint
}
