package configs

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/moquayson/task-api-assignment/database"
	"log"
	"os"
)

var DBContext *sql.DB

func init() {
	LoadEnvFile()
	ConnectToDatabase()
}

func ConnectToDatabase() {
	dbUrl := os.Getenv("DATABASE_URL")
	if len(dbUrl) == 0 {
		dbUrl = "./payment.db"
	}
	// Initialize the database
	var err error
	DBContext, err = sql.Open("sqlite3", dbUrl)
	if err != nil {
		log.Fatalln(err)
	}
	//defer DBContext.Close()

	log.Println("connected to database")

	// migrate tables
	if err = database.MigrateTables(DBContext); err != nil {
		log.Fatalln(err)
	}

	//seed data
	if err = database.SeedAdminUser(DBContext); err != nil {
		log.Fatalln(err)
	}

}
