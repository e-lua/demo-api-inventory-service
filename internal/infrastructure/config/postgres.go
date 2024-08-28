package config

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnMasterPostgres(env_db_url string) *pgxpool.Pool {

	//Create the context
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	//Handle conection with the master
	urlString := env_db_url
	config, _ := pgxpool.ParseConfig(urlString)

	//Set fields to pool the connections
	config.MaxConns = 80                       //Num max conn
	config.MinConns = 20                       //Num min conn
	config.MaxConnIdleTime = 5 * time.Minute   //Max time a conn can be inactive
	config.HealthCheckPeriod = 1 * time.Minute //Time of period to check conn health

	//Create pool of conn
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	return pool
}
