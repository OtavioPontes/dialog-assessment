package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/otaviopontes/api-go/src/config"
	"github.com/redis/go-redis/v9"
)

func Connect() (*sql.DB, error) {
	db, err := sqlx.Connect("postgres", config.ConectionString)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db.DB, nil
}

func ConnectRedis() (*redis.Client, error) {
	var ctx = context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDb,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
