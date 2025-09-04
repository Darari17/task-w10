package main

import (
	"log"

	"github.com/Darari17/task-w10/day-2/internal/configs"
	"github.com/Darari17/task-w10/day-2/internal/routers"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load env\nCause: ", err.Error())
		return
	}

	db, err := configs.InitDB()
	if err != nil {
		log.Println("Failed to connect to database\nCause: ", err.Error())
		return
	}
	defer db.Close()

	if err := configs.TestDB(db); err != nil {
		log.Println("Ping to DB failed\nCause: ", err.Error())
		return
	}
	log.Println("DB Connected")

	router := routers.InitRouter(db)

	router.Run("localhost:8080")
}
