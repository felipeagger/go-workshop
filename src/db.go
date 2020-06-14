package src

import (
	"fmt"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	dbName = "go_workshop"
	dbUser = "admin"
	dbPort = "5432"
)

var (
	dbHost string = os.Getenv("DB_HOST")
	dbPass string = os.Getenv("DB_PASSWD")
)

func dbConnect() *gorm.DB {

	urlConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)

	db, err := gorm.Open("postgres", urlConn)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	if err != nil {
		fmt.Println("Error in conection with database!")
		panic(err)
	}

	return db
}

func selectAll(table interface{}, c *gin.Context) interface{} {

	db := dbConnect()

	if err := db.Find(reflect.ValueOf(table).Interface()).Error; err != nil {
		checkErr(err, c)
	}

	return table
}

func selectUserID(ID int, c *gin.Context) (user, *gorm.DB) {

	db := dbConnect()
	var user user

	if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
		checkErr(err, c)
	}

	return user, db
}

func selectDebtByID(ID string, c *gin.Context) (debt, *gorm.DB) {

	db := dbConnect()
	var debt debt

	if err := db.Where("id = ?", ID).First(&debt).Error; err != nil {
		checkErr(err, c)
	}

	return debt, db
}

func selectDebtsByUser(userID int, c *gin.Context) []debt {

	db := dbConnect()
	var debts []debt

	if err := db.Where("user_id = ?", userID).Find(&debts).Error; err != nil {
		checkErr(err, c)
	}

	return debts
}

func checkErr(err error, c *gin.Context) {
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
}
