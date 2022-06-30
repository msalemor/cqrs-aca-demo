package configs

import (
	"errors"
	"fmt"
	"goservices/dataaccess"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	FiberApp   *fiber.App
	Collection *mongo.Collection
	mongoURI   string
	SPort      string
}

var AppConfig Config

func loadEnv(envVariable string, required bool, defaultValue string) (string, error) {
	val := os.Getenv(envVariable)
	if required && defaultValue != "" {
		return defaultValue, nil
	}
	if required && val == "" {
		msg := fmt.Sprintf("Environment variable: %s is required.", envVariable)
		return "", errors.New(msg)
	}
	return val, nil
}

func New() Config {
	mongoURI, _ := loadEnv("MONGO_URI", true, "")
	databaseName, err := loadEnv("DATABASE_NAME", true, "vendors")
	if err != nil {
		log.Fatalln(err)
	}
	collectionName, _ := loadEnv("COLLECTION_NAME", true, "default")
	if mongoURI == "" {
		log.Fatalln("You need to set the MONGO_URI environment variable")
	}
	sport := os.Getenv("PORT")
	port := 3000
	if sport != "" {
		port, _ = strconv.Atoi(sport)
	}
	sport = fmt.Sprintf(":%v", port)
	config := Config{FiberApp: fiber.New(), Collection: dataaccess.ConnectDB(mongoURI, databaseName, collectionName), mongoURI: mongoURI, SPort: sport}
	return config
}
