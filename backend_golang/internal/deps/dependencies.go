package deps

import (
	"database/sql"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/handler"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/repository"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/security"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/service"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"gorm.io/gorm"
)

type Dependencies struct {
	MySQL 			*sql.DB
	MongoDB 		*mongo.Client
	Redis 			*redis.Client
	GORM			*gorm.DB
	UserRepository 	repository.UserRepository
	AuthService    	*service.AuthService
	AuthHandler    	*handler.AuthHandler
	JwtService		*security.JWTService
	DictionaryRepository repository.BankWordRepository
	DictionaryService    *service.DictionaryService
	DictionaryHandler    *handler.DictionaryHandler
}