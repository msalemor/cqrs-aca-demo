package configs

import (
	"errors"
	"fmt"
	"goservices/dataaccess"
	"log"
	"os"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	FiberApp   *fiber.App
	Collection *mongo.Collection
	mongoURI   string
	SPort      string
	Validate   *validator.Validate
}

var App Config

func loadEnv(envVariable string, required bool, defaultValue string) (string, error) {
	val := os.Getenv(envVariable)
	if val != "" {
		return val, nil
	} else if val == "" && defaultValue != "" {
		return defaultValue, nil
	} else if required && val == "" {
		msg := fmt.Sprintf("Environment variable: %s is required.", envVariable)
		return "", errors.New(msg)
	}
	return "", nil
}

func Init(env string) {
	err := godotenv.Load(env)
	if err != nil {
		log.Fatalf("Unable to load the %s file. Err: %s\n", env, err)
	}

	mongoURI, _ := loadEnv("MONGO_URI", true, "")

	sport := os.Getenv("PORT")
	port := 3000
	if sport != "" {
		port, _ = strconv.Atoi(sport)
	}
	sport = fmt.Sprintf(":%v", port)

	databaseName, err := loadEnv("DATABASE_NAME", true, "vendorsDb")
	if err != nil {
		log.Fatalln(err)
	}

	collectionName, _ := loadEnv("COLLECTION_NAME", true, "vendorsCol")
	if mongoURI == "" {
		log.Fatalln("You need to set the MONGO_URI environment variable")
	}
	validate := validator.New()
	config := Config{FiberApp: fiber.New(),
		Collection: dataaccess.ConnectDB(mongoURI, databaseName, collectionName),
		mongoURI:   mongoURI,
		SPort:      sport,
		Validate:   validate}

	App = config

}
