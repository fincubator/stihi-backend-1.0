package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetConfigs() map[string]string {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found, will use environment variables from os")
	}

	keys := map[string]string{}
	configFilePath, _   := os.LookupEnv("CONFIG")
	configDbPath, _ 	:= os.LookupEnv("DB_CONFIG")
	configMongoPath, _  := os.LookupEnv("MONGO_DB_CONFIG")
	configRedisPath, _  := os.LookupEnv("REDIS_CONFIG")

	keys["configFilePath"] = configFilePath
	keys["configDbPath"] = configDbPath
	keys["configMongoPath"] = configMongoPath
	keys["configRedisPath"] = configRedisPath

	return keys
}