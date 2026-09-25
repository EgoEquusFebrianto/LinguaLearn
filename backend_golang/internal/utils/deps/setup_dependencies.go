package deps

import (
	"context"
	"database/sql"
	"time"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/config"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/delivery/http/handler"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/domain"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/infrastructure/database/mongodb"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/infrastructure/database/mysql"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/infrastructure/database/redisdb"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/repository"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/security"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/service"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"gorm.io/gorm"
)

type Dependencies struct {
	MySQL                *sql.DB
	MongoDB              *mongo.Client
	Redis                *redis.Client
	GORM                 *gorm.DB
	UserRepository       domain.UserRepository
	AuthService          *service.AuthService
	AuthHandler          *handler.AuthHandler
	JwtService           *security.JWTService
	DictionaryRepository domain.BankWordRepository
	BankWordService      *service.BankWordService
	BankWordHandler    *handler.BankWordHandler
}

func SetupDependencies(
	cfg *config.Config,
) (*Dependencies, error) {
	// DATABASES
	db, err := mysql.NewMySQL(&cfg.MySQL)
	if err != nil {
		return nil, err
	}

	mongoClient, err := mongodb.NewMongoDb(&cfg.MongoDb)
	if err != nil {
		db.Close()
		return nil, err
	}

	redisClient, err := redisdb.NewRedis(&cfg.Redis)
	if err != nil {
		db.Close()

		ctx, cancel := context.WithTimeout(
			context.Background(),
			5*time.Second,
		)
		defer cancel()

		mongoClient.Disconnect(ctx)

		return nil, err
	}

	gormDB, err := mysql.NewGorm(db)
	if err != nil {
		db.Close()
		redisClient.Close()

		ctx, cancel := context.WithTimeout(
			context.Background(),
			5*time.Second,
		)
		defer cancel()

		mongoClient.Disconnect(ctx)

		return nil, err
	}

	// REPOSITORY
	userRepository := repository.NewUserRepository(gormDB)

	bankWordRepository := repository.NewBankWordRepository(
		mongoClient,
		cfg.MongoDb.MongoDb,
	)

	// SECURITY
	jwtService := security.NewJWTService(
		cfg.JWT.AccessSecret,
		time.Duration(
			cfg.JWT.AccessExpiresMinute,
		)*time.Minute,
	)

	passwordHasher := security.NewPasswordHasher()

	// SERVICE
	refreshTokenService :=
		service.NewRefreshTokenService(redisClient)

	authService := service.NewAuthService(
		userRepository,
		passwordHasher,
		jwtService,
		refreshTokenService,
	)

	BankWord := service.NewBankWordServiceService(
		bankWordRepository,
	)

	// HANDLER
	authHandler := handler.NewAuthHandler(authService)

	bankWordHandler := handler.NewBankWordHandler(BankWord)

	// DEPENDENCIES
	dependencies := &Dependencies{
		MySQL:   db,
		MongoDB: mongoClient,
		Redis:   redisClient,
		GORM:    gormDB,

		UserRepository: userRepository,

		AuthService: authService,
		AuthHandler: authHandler,
		JwtService:  jwtService,

		DictionaryRepository:	bankWordRepository,
		BankWordHandler:		bankWordHandler,
	}

	return dependencies, nil
}
