// package config

// import (
// 	"log"
// 	"os"

// 	"github.com/go-pg/pg"
// 	"github.com/go-pg/pg/orm"
// 	"github.com/go-pg/pg/v10"
// )

// //Database Connection
// func Connect() *pg.DB{
// 	options := &pg.Options{
// 		User:	os.Getenv("DB_USER"),
// 		Password: os.Getenv("DB_PASSWORD"),
// 		Address: os.Getenv("DB_DATABASE"),
// 	}

// 	var db *pg.DB = pg.Connect(options)
// 	if db == nil{
// 		log.Fatal("Database Connection Failed.......\n")
// 		os.Exit(100)
// 	}

// 	log.Printf("Database Connected Successfullly......\n")

// 	if err := createSchema(db); err !=nil{
// 		log.Fatal(err)
// 	}

// 	return db
// }

// func createSchema(db *pg.DB) error{
// 	models:=[]interface{}{
// 		(*Book)(nil),
// 	}

// 	for _, model := range models {
// 		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
// 			IfNotExists: true,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// import (
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
//   )

//   dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
//   db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Conntect() {

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbUri)

	d, err := gorm.Open("postgres", dbUri)

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
