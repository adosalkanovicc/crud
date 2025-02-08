package config

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var DB *sql.DB

func InitDB() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read the environment variables using viper
	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbName := viper.GetString("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Can't reach the database:", err)
	}
	fmt.Println("Connected to the database successfully!")

	runMigration()
}

func runMigration() {

	migrationFile := "migrations/db_schema.sql"
	fileContent, err := ioutil.ReadFile(migrationFile)
	if err != nil {
		log.Fatalf("Error reading migration file: %v", err)
	}

	_, err = DB.Exec(string(fileContent))
	if err != nil {
		log.Fatalf("Error running migration: %v", err)
	}
	fmt.Println("Database migration applied successfully!")
}
