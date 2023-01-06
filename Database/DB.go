package Database

import (
	"database/sql"
	"log"
)



var db *sql.DB

func InitDB(urlDB string) *sql.DB{
	var err error
	//urlDB = "postgres://grzuanbs:J_Q4hKYnrgJmSBu8UnJPoxK85vmGhgLq@john.db.elephantsql.com/grzuanbs"
	db, err = sql.Open("postgres", urlDB)
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	//defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
	`
	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}
 
	return db

}




