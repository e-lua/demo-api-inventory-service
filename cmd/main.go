package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	routes "github.com/e-lua/demo-api-inventory-clean-architecture/cmd/api/routes"

	jobs "github.com/e-lua/demo-api-inventory-clean-architecture/cmd/jobs"
	config "github.com/e-lua/demo-api-inventory-clean-architecture/internal/infrastructure/config"

	kardex_supply_repository "github.com/e-lua/demo-api-inventory-clean-architecture/internal/repositories/postgres/kardex_supply"
	measure_repository "github.com/e-lua/demo-api-inventory-clean-architecture/internal/repositories/postgres/measure"
	provider_repository "github.com/e-lua/demo-api-inventory-clean-architecture/internal/repositories/postgres/provider"
	supply_repository "github.com/e-lua/demo-api-inventory-clean-architecture/internal/repositories/postgres/supply"
	warehouse_repository "github.com/e-lua/demo-api-inventory-clean-architecture/internal/repositories/postgres/warehouse"
	job_repository "github.com/e-lua/demo-api-inventory-clean-architecture/internal/repositories/redis/job"

	automated_service "github.com/e-lua/demo-api-inventory-clean-architecture/internal/services/automated"
	kardex_supply_service "github.com/e-lua/demo-api-inventory-clean-architecture/internal/services/kardex_supply"
	measure_service "github.com/e-lua/demo-api-inventory-clean-architecture/internal/services/measure"
	provider_service "github.com/e-lua/demo-api-inventory-clean-architecture/internal/services/provider"
	supply_service "github.com/e-lua/demo-api-inventory-clean-architecture/internal/services/supply"
	warehouse_service "github.com/e-lua/demo-api-inventory-clean-architecture/internal/services/warehouse"

	kardex_supply_handler "github.com/e-lua/demo-api-inventory-clean-architecture/cmd/api/handlers/kardex_supply"
	measure_handler "github.com/e-lua/demo-api-inventory-clean-architecture/cmd/api/handlers/measure"
	provider_handler "github.com/e-lua/demo-api-inventory-clean-architecture/cmd/api/handlers/provider"
	supply_handler "github.com/e-lua/demo-api-inventory-clean-architecture/cmd/api/handlers/supply"
	warehouse_handler "github.com/e-lua/demo-api-inventory-clean-architecture/cmd/api/handlers/warehouse"
)

func main() {

	//Read variable of the environment .env.dev
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatalf("Error loading file .env.dev")
	}

	//Setup Repositories
	conn_master_postgres := config.ConnMasterPostgres(os.Getenv("URL_PG_DATABASE_MASTER"))
	defer conn_master_postgres.Close()

	conn_master_redis := config.ConnMasterRedis(os.Getenv("URL_REDIS_SERVER"))
	defer conn_master_redis.Close()

	kardex_supply_repository := kardex_supply_repository.NewKardexSupplyRepository(conn_master_postgres)
	measure_repository := measure_repository.NewMeasureRepository(conn_master_postgres)
	provider_repository := provider_repository.NewMeasureRepository(conn_master_postgres)
	supply_repository := supply_repository.NewSupplyRepository(conn_master_postgres)
	warehouse_repository := warehouse_repository.NewWarehouseRepository(conn_master_postgres)
	job_repository := job_repository.NewJobRedisRepository(conn_master_redis)

	//Setup Services
	kardex_supply_service := kardex_supply_service.NewKardexSupplyService(kardex_supply_repository)
	measure_service := measure_service.NewMeasureService(measure_repository)
	provider_service := provider_service.NewProviderService(provider_repository, supply_repository)
	supply_service := supply_service.NewSupplyService(supply_repository)
	warehouse_service := warehouse_service.NewWarehouseService(warehouse_repository, supply_repository)
	automated_service := automated_service.NewAutomatedService(warehouse_repository, supply_repository, provider_repository, job_repository)

	//Setup Handler
	kardex_supply_handler := kardex_supply_handler.NewKardexSupplyHandler(kardex_supply_service)
	measure_handler := measure_handler.NewMeasureHandler(measure_service)
	provider_handler := provider_handler.NewProviderHandler(provider_service)
	supply_handler := supply_handler.NewSupplyHandler(supply_service)
	warehouse_handler := warehouse_handler.NewWarehouseHandler(warehouse_service)

	//Start Jobs
	job := jobs.JobConfig{
		AutomatedService: automated_service,
	}
	job.Start()

	//Start Api
	route_config := routes.RouteConfig{
		App:                 echo.New(),
		KardexSupplyHandler: kardex_supply_handler,
		MeasureHandler:      measure_handler,
		ProviderHandler:     provider_handler,
		SupplyHandler:       supply_handler,
		WarehouseHandler:    warehouse_handler,
	}
	route_config.Start()

}
