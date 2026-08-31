package app

import (
	"context"
	"time"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/config"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/database/mongodb"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/database/mysql"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/database/redis"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/deps"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/handler"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/repository"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/security"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/service"
)

func SetupDependencies(
	cfg *config.Config,
) (*deps.Dependencies, error) {
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

	redisClient, err := redis.NewRedis(&cfg.Redis)
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

	dictionaryService := service.NewDictionaryService(
		bankWordRepository,
	)

	// HANDLER
	authHandler := handler.NewAuthHandler(authService)

	dictionaryHandler := handler.NewDictionaryHandler(dictionaryService)

	// DEPENDENCIES
	dependencies := &deps.Dependencies{
		MySQL:          db,
		MongoDB:        mongoClient,
		Redis:          redisClient,
		GORM:           gormDB,
		
		UserRepository: userRepository,

		AuthService:    authService,
		AuthHandler:    authHandler,
		JwtService:     jwtService,

		DictionaryRepository: bankWordRepository,
		DictionaryService:    dictionaryService,
		DictionaryHandler:    dictionaryHandler,
	}

	return dependencies, nil
}