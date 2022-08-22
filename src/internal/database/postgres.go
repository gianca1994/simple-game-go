package database

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"regexp"
	"simple-game-golang/src/model"
	"strconv"
)

func NewPostgreSQL() {
	projectName := regexp.MustCompile(`^(.*` + "simple-game-golang" + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	if err := godotenv.Load(string(rootPath) + `/.env`); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(0)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPass, dbName, dbSSLMode)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return
	}
	/*
		db.Create(&model.User{
			Name:       "Gianca",
			DamageMax:  10,
			DamageMin:  5,
			Health:     50,
			DefenseMax: 3,
			DefenseMin: 1,
		})
	*/

	/*
		userSearch := model.User{}
		db.Find(&test, 1)
		fmt.Println(test.Name)
	*/

	fmt.Println("Successfully migrated the database")
}
