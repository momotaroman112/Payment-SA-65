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

	Bookings	[]Booking	 `gorm:"foreignKey:UserID"`
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

	//Member_id ทำหน้าที่เป็น fk
	// ไม่จำเป็นต้องใช้ เพราะ Member มีอยู่ใน Booking ซึ่งคนที่สั่งอาหารต้องทำการ Booking ก่อน
	// MemberID *uint
	// Member    User


	BookingID *uint
	Booking   Booking

	TotalPrice uint

	//คสพ 1 อยู่ฝั่ง FoodOrdered
	FoodOrderedFoodSets []FoodOrderedFoodSet `gorm:"foreignKey:FoodOrderedID"`

	//FFID *uint
	//FF    FoodOrdered_FoodSet
}


type PaymentType struct {
	gorm.Model
	Name string
	Bill []Bill `gorm:"foreignKey:PaymentTypeID"`
}


type Booking struct {
	gorm.Model
	BookingTimeStart  time.Time
	BookingTimeStop time.Time
	Room		  string
	TotalPrice     uint

	UserID     *uint
	User      	User

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

	FoodOrderedID *uint
	FoodOrdered   FoodOrdered

	TotalPrice uint
}
