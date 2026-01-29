package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConnectDatabase = ""

	Port = 0

	SecretKey []byte
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: .env não encontrado, usando variáveis de ambiente do Render")
	}

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = os.Getenv("API_PORT")
		if portStr == "" {
			portStr = "9000"
		}
	}
	Port, err = strconv.Atoi(portStr)
	if err != nil {
		Port = 9000
	}
		StringConnectDatabase = os.Getenv("DATABASE_URL")
		if StringConnectDatabase == "" {
			StringConnectDatabase = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require",
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_NAME"),
			)
		}
	
			SecretKey = []byte(os.Getenv("SECRET_KEY"))
		}
