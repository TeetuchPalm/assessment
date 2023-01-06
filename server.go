package main

import (
	"fmt"
	"os"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"database/sql"
	"log"
	"apiExpense/Database"

)


var db *sql.DB

func main() {

	db = Database.InitDB("postgres://grzuanbs:J_Q4hKYnrgJmSBu8UnJPoxK85vmGhgLq@john.db.elephantsql.com/grzuanbs")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))
	log.Fatal(e.Start(":2565"))
}
