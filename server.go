package main

import (
	"apiEx/Database"
	"apiEx/handler"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


var db *sql.DB

func main() {

	db = Database.InitDB("postgres://grzuanbs:J_Q4hKYnrgJmSBu8UnJPoxK85vmGhgLq@john.db.elephantsql.com/grzuanbs")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(handler.GetDB(db))

	e.POST("/expenses", handler.CreateExpensesHandler)
	e.GET("/expenses/:id", handler.GetExpensesHandler)
	e.GET("/expenses", handler.GetAllExpensesHandler)
	e.PUT("/expenses/:id", handler.UpdateExpensesHandler)

	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))
	log.Fatal(e.Start(os.Getenv("PORT")))

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

