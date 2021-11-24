package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/AirellJordan98/hacktiv_ecommerce/api"
	db "github.com/AirellJordan98/hacktiv_ecommerce/db/sqlc"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")

	fmt.Println("HOST:", os.Getenv("DB_HOST"))
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require TimeZone=Asia/Shanghai",
		host, dbPort, user, password, dbName)
	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
	// 	host, dbPort, user, password, dbName)

	// DBSource := "postgresql://sofyan:postgres@localhost:5432/hacktiv-ecommerce?sslmode=disable"
	conn, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalln("error connecting to db", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store, conn)

	err = server.Start(":" + port)

	if err != nil {
		log.Fatalln("error while starting server", err)
	}

}
