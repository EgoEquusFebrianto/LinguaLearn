package config

import (
	"fmt"
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
	MongoName	string
}

type RedisConfig struct {
	RedisADDR 	string
	RedisPwd	string
	RedisDb		int
}

type ServerConfig struct {
	Address		string
}

type JwtConfig struct {
	AccessSecret 		string
	AccessExpiresMinute int
	RefreshExpiresdays 	int
}

type Config struct {
	MySQL 	MySQLConfig
	MongoDb MongoConfig
	Redis 	RedisConfig
	Server	ServerConfig
	JWT		JwtConfig
}

func (cfg *MySQLConfig) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
}

func Load() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	redisDb, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	minute, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_EXPIRES_MINUTES"))
	expiration, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRES_DAYS"))

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
			MongoName: os.Getenv("MONGO_NAME"),
		},
		Redis: RedisConfig{
			RedisADDR: os.Getenv("REDIS_ADDR"),
			RedisPwd: os.Getenv("REDIS_PASSWORD"),
			RedisDb: redisDb,
		},
		Server: ServerConfig{
			Address: os.Getenv("SERVER_ADDRESS"),
		},
		JWT: JwtConfig{
			AccessSecret: os.Getenv("JWT_ACCESS_SECRET"),
			AccessExpiresMinute: minute,
			RefreshExpiresdays: expiration,
		},
	}, nil
}