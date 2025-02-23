package config

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	dsn := "root:erudent123@unix(/opt/homebrew/var/mysql/mysql.sock)/users_db?parseTime=true&tls=false"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Open connection
	// db, err := sql.Open("mysql", dsn)
	// if err != nil {
	// 	log.Fatal("Error opening DB:", err)
	// }
	// defer db.Close()

	// // Ping the database to check connection
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal("Error connecting to DB:", err)
	// }

	fmt.Println("Connected to MySQL database successfully!")

}
