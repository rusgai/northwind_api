package main

import (
	_ "github.com/go-playground/validator/v10"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "github.com/redis/go-redis/v9"
	_ "github.com/rs/zerolog/log"
	_ "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/swag"
)
