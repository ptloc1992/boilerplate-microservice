package cmd

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/urfave/cli"
)

func newLogger() *log.Logger {
	logger := log.New(os.Stdout, "", 0)
	return logger
}

func newDB(appContext *cli.Context) *gorm.DB {
	db, err := gorm.Open("mysql", appContext.GlobalString("database"))
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return db
}

func newRabbitMQ() {

}

func newReddis() {

}
