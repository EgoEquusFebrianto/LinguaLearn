package helper

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"gorm.io/gorm"
)

type Dependencies struct {
	MySQL 	*sql.DB
	MongoDB *mongo.Client
	Redis 	*redis.Client
	GORM	*gorm.DB
}