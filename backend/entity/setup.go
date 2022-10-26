package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Booking{},
		&Employee{},
		&User{},
		&FoodSet{},
		&FoodOrderedFoodSet{},
		&FoodOrdered{},
		&PaymentType{},
		&Bill{},
	)
	db = database
	password, err := bcrypt.GenerateFromPassword([]byte("45191"), 14)
	if err != nil {
		panic("failed")
	}

	db.Model(&User{}).Create(&User{
		Email:    "moon@gmil.com",
		Name:     "Kiadtisak J",
		Password: string(password),
	})

	db.Model(&User{}).Create(&User{
		Name:     "Fei",
		Email:    "man@gmil.com",
		Password: string(password),
	})

	var jo User
	var fei User

	db.Raw("SELECT * FROM users WHERE email = ?", "moon@gmil.com").Scan(&jo)
	db.Raw("SELECT * FROM users WHERE email = ?", "man@gmil.com").Scan(&fei)

	//Booking
	db.Model(&Employee{}).Create(&Employee{
		Name:     "Komson",
		Password: string(password),
		Email:    "komson@gmai.com",
	})

	db.Model(&Employee{}).Create(&Employee{
		Name:     "wichi",
		Password: string(password),
		Email:    "wichai@gmai.com",
	})

	var a User
	var b User

	db.Raw("SELECT * FROM employees WHERE email = ?", "komson@gmai.com").Scan(&a)
	db.Raw("SELECT * FROM employees WHERE email = ?", "wichai@gmai.com").Scan(&b)

	// Food sets
	FoodSet1 := FoodSet{
		Name:   "Set A",
		Detail: "Detail1",
		Price:  150,
	}
	FoodSet2 := FoodSet{
		Name:   "Set B",
		Detail: "Detail2",
		Price:  120,
	}
	FoodSet3 := FoodSet{
		Name:   "Set C",
		Detail: "Detail3",
		Price:  250,
	}
	db.Model(&FoodSet{}).Create(&FoodSet1)
	db.Model(&FoodSet{}).Create(&FoodSet2)
	db.Model(&FoodSet{}).Create(&FoodSet3)

	//FoodOrdereds
	Foodordered1 := FoodOrdered{
		FoodTime:   time.Date(2022, 10, 16, 12, 00, 00, 00, time.Local),
		TotalPrice: 420,
		FoodOrderedFoodSets: []FoodOrderedFoodSet{ { FoodSet: FoodSet1, Quantity: 2 }, { FoodSet: FoodSet2, Quantity: 1 } },
	}
	db.Model(&FoodOrdered{}).Create(&Foodordered1)

	Foodordered2 := FoodOrdered{
		FoodTime:   time.Date(2022, 10, 16, 12, 00, 00, 00, time.Local),
		TotalPrice: 490,
		FoodOrderedFoodSets: []FoodOrderedFoodSet{ { FoodSet: FoodSet2, Quantity: 2 }, { FoodSet: FoodSet3, Quantity: 1 } },
	}
	db.Model(&FoodOrdered{}).Create(&Foodordered2)

	Foodordered3 := FoodOrdered{
		FoodTime:   time.Date(2022, 10, 16, 12, 00, 00, 00, time.Local),
		TotalPrice: 150,
		FoodOrderedFoodSets: []FoodOrderedFoodSet{ { FoodSet: FoodSet1, Quantity: 1 } },
	}
	db.Model(&FoodOrdered{}).Create(&Foodordered3)

	// FoodOrderedFoodsets

	// Booking
	Booking1 := Booking{
		BookingTimeStart: time.Date(2022, 10, 16, 12, 00, 00, 00, time.Local),
		BookingTimeStop:  time.Date(2022, 10, 16, 12, 00, 00, 00, time.Local),
		Room:             "1012",
		TotalPrice:       3000,
		User:             jo,
		FoodOrdereds:     []FoodOrdered{Foodordered1, Foodordered3},
	}
	db.Model(&Booking{}).Create(&Booking1)

	Booking2 := Booking{
		BookingTimeStart: time.Date(2022, 10, 16, 12, 00, 00, 00, time.Local),
		BookingTimeStop:  time.Date(2022, 10, 16, 12, 00, 00, 00, time.Local),
		Room:             "1112",
		TotalPrice:       2000,
		User:             fei,
		FoodOrdereds:     []FoodOrdered{Foodordered2},
	}
	db.Model(&Booking{}).Create(&Booking2)

	//PaymentType
	PaymentType1 := PaymentType{
		Name: "เงินสด",
	}
	db.Model(&PaymentType{}).Create(&PaymentType1)
	PaymentType2 := PaymentType{
		Name: "ธนาคาร A",
	}
	db.Model(&PaymentType{}).Create(&PaymentType2)
	PaymentType3 := PaymentType{
		Name: "ธนาคาร B",
	}
	db.Model(&PaymentType{}).Create(&PaymentType3)

	//Bill1
	/*db.Model(&Bill{}).Create(&Bill{

		Booking: Booking1,
		FoodOrdered:        Place1,
		PaymentType:  PaymentType1,
		BillTime:     time.Now(),
		TotalPrice:   (Booking1.TotalCredit * 800),
	})*/

	/*//Bill2
	db.Model(&Bill{}).Create(&Bill{

		Booking: Booking2,
		FoodOrdered:        Place2,
		PaymentType:  PaymentType2,
		BillTime:     time.Now(),
		TotalPrice:   (Booking2.TotalCredit * 800),
	})

	//Bill3
	db.Model(&Bill{}).Create(&Bill{

		Booking: Booking3,
		FoodOrdered:        Place3,
		PaymentType:  PaymentType3,
		BillTime:     time.Now(),
		TotalPrice:   (Booking2.TotalCredit * 800),
	})*/

}
