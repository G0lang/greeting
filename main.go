package main

import (
	"github.com/G0lang/greeting/app"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// VERSION keeps the version of application
	VERSION string
)

func main() {

	app := &app.App{Version: VERSION}
	app.InitializeDB()
	app.Run()
}
