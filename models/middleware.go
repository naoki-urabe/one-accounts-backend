package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"one-accounts/config"
	"time"
)

type Detail struct {
	Uuid        string    `json:"uuid"`
	TradingDay  time.Time `json:"trading_day,string"`
	TradingName string    `json:"trading_name"`
	Payment     int       `json:"payment,string"`
	Balance     int       `json:"balance,string"`
	Note        string    `json:"note"`
	Bank        string    `json:"bank"`
}

var Db *gorm.DB

func GetAccountDetails(details *[]Detail) {
	Db.Find(&details)
}

func InsertDetail(detail *Detail) {
	Db.NewRecord(detail)
	Db.Create(&detail)
}

func init() {
	var err error
	dbConnectInfo := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		config.Config.DbUserName,
		config.Config.DbUserPassword,
		config.Config.DbHost,
		config.Config.DbPort,
		config.Config.DbName,
	)
	//fmt.Println(dbConnectInfo)
	Db, err = gorm.Open(config.Config.DbDriverName, dbConnectInfo)
	// Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Successfully connect database..")
	}
	Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Detail{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created table...")
	}
}
