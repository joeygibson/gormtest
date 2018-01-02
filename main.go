package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
	Description string
}

func main() {
	fmt.Println("1")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect to db")
	}

	defer db.Close()

	fmt.Println("2")
	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "L1212", Price: 1000, Description: "A fancy product"})

	var product Product

	db.First(&product, 1)
	fmt.Printf("Prod: %v\n", product)

	db.First(&product, "code = ?", "L1212")
	fmt.Printf("Prod: %v\n", product)

	db.Model(&product).Update("Price", 2000)

	db.First(&product, "code = ?", "L1212")
	fmt.Printf("Prod: %v\n", product)

	db.Create(&Product{Code: "B900", Price: 123})
	if db.Error != nil {
		fmt.Printf("ERROR!: %v\n", db.Error)
	}
	db.First(&product, "code = ?", "B900")
	if db.Error != nil {
		fmt.Printf("ERROR!: %v\n", db.Error)
	}
	fmt.Printf("Prod: %v\n", product)

	fmt.Println("23")
	db.Delete(&product)
}
