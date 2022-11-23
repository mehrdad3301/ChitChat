package db 

import ( 

	"log"
	"time"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
) 

var db *sql.DB

func init() { 

	d, err := sql.Open("sqlite3", "./database.db") 
	if err != nil { 
		log.Fatal("db: init: ", err) 
	}
	db = d 
} 

func getTime() { 
	time.Now().Format(time.ANSIC))
}
