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
	configFilePath, _   := os.LookupEnv("CONFIG_FILE_PATH")
	configDbPath, _ 	:= os.LookupEnv("CONFIG_DB_PATH")
	configMongoPath, _  := os.LookupEnv("CONFIG_MONGO_DB_PATH")
	configRedisPath, _  := os.LookupEnv("CONFIG_REDIS_PATH")

	keys["configFilePath"] = configFilePath
	keys["configDbPath"] = configDbPath
	keys["configMongoPath"] = configMongoPath
	keys["configRedisPath"] = configRedisPath

	return keys
}