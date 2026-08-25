package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type MySQLConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

type MongoConfig struct {
	MongoUrl	string
	MongoUser	string
	MongoPwd	string
	MongoDb		string
}

type RedisConfig struct {
	RedisADDR 	string
	RedisPwd	string
	RedisDb		int
}

type ServerConfig struct {
	Address		string
}

type Config struct {
	MySQL 	MySQLConfig
	MongoDb MongoConfig
	Redis 	RedisConfig
	Server	ServerConfig
}

func Load() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	redisDb, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	return &Config{
		MySQL: MySQLConfig{
			DBHost: os.Getenv("DB_HOST"),
			DBPort: os.Getenv("DB_PORT"),
			DBUser: os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName: os.Getenv("DB_NAME"),
		},
		MongoDb: MongoConfig{
			MongoUrl: os.Getenv("MONGO_URL"),
			MongoUser: os.Getenv("MONGO_USER"),
			MongoPwd: os.Getenv("MONGO_PASSWORD"),
			MongoDb: os.Getenv("MONGO_NAME"),
		},
		Redis: RedisConfig{
			RedisADDR: os.Getenv("REDIS_ADDR"),
			RedisPwd: os.Getenv("REDIS_PASSWORD"),
			RedisDb: redisDb,
		},
		Server: ServerConfig{
			Address: os.Getenv("SERVER_ADDRESS"),
		},
	}, nil
}