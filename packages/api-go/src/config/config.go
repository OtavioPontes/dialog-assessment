package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConectionString = ""
	Port            = 0
	FrontEndUrl     = ""
	RedisAddr       = ""
	RedisPassword   = ""
	RedisDb         = 0
)
var SecretKey []byte

func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	RedisAddr = fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	RedisPassword = os.Getenv("REDIS_PASSWORD")
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		RedisDb = 16
	}
	RedisDb = db

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 9000
	}

	FrontEndUrl = os.Getenv("FRONTEND_URL")

	ConectionString = fmt.Sprintf(
		"user=%s dbname=%s sslmode=disable password=%s host=%s port=%s",

		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}

func LoadTest() {
	var err error

	if err = godotenv.Load("../../.env"); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 9000
	}

	FrontEndUrl = os.Getenv("FRONTEND_URL")

	ConectionString = fmt.Sprintf(
		"user=%s dbname=%s sslmode=disable password=%s host=%s port=%s",

		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
