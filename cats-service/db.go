package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SetupDB : initializing mysql database
func SetupDB() *gorm.DB {
	URL := initFromEnvs()

	db, err := gorm.Open(mysql.Open(URL), &gorm.Config{})
	if err != nil {
		log.Panic(err.Error())
	}

	db.AutoMigrate(&Cat{})

	initSeed(db)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func initFromEnvs() string {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	db := os.Getenv("MYSQL_DATABASE")

	if port == "0" || port == "" {
		port = "3306"
	}

	if host == "" {
		log.Println("warning: MYSQL_HOST is not defined, using localhost")
		host = "localhost"
	}

	if db == "" {
		db = "catCRUD"
	}

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, db)
	return URL
}

func initSeed(db *gorm.DB) {
	newCats := []Cat{
		{Name: "Kwat",
			Art: `
__
\ \______  ^-^
 \/      \(-x-)
 |________  _ \
          || ||
          "" ""`,
		},
		{Name: "Kwot",
			Art: `
    /\       /\
   /  \_____/  \
  /             \
 (   ()    ()    )
 [              ]
(        Y       )
 \    \_/\_/    /
  \____________/`,
		},

		{Name: "Night cot",
			Art: `
              a          a
             aaa        aaa
            aaaaaaaaaaaaaaaa
           aaaaaaaaaaaaaaaaaa
          aaaaafaaaaaaafaaaaaa
          aaaaaaaaaaaaaaaaaaaa
           aaaaaaaaaaaaaaaaaa
            aaaaaaa  aaaaaaa
             aaaaaaaaaaaaaa
  a         aaaaaaaaaaaaaaaa
 aaa       aaaaaaaaaaaaaaaaaa
 aaa      aaaaaaaaaaaaaaaaaaaa
 aaa     aaaaaaaaaaaaaaaaaaaaaa
 aaa    aaaaaaaaaaaaaaaaaaaaaaaa
  aaa   aaaaaaaaaaaaaaaaaaaaaaaa
  aaa   aaaaaaaaaaaaaaaaaaaaaaaa
  aaa    aaaaaaaaaaaaaaaaaaaaaa
   aaa    aaaaaaaaaaaaaaaaaaaa
    aaaaaaaaaaaaaaaaaaaaaaaaaa
     aaaaaaaaaaaaaaaaaaaaaaaaa
	`,
		},

		{
			Name: "Kitty",
			Art: `
	/\___/\
   (  o o  )
   /   *   \
   \__\_/__/ meow!
     /   \
    / ___ \
    \/___\/
			`,
		},
	}

	// var cats []Cat
	// result := db.Find(&cats)

	// if int(result.RowsAffected) != len(newCats) {
	log.Println("Updating cats...")

	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&newCats)

	if db.Error != nil {
		log.Println("db error: ", db.Error.Error())
	}
	// }

}
