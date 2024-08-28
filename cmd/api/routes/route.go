package route

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"

	kardex_supply_handler "github.com/e-lua/demo-api-inventory-clean-architecture/cmd/api/handlers/kardex_supply"
	measure_handler "github.com/e-lua/demo-api-inventory-clean-architecture/cmd/api/handlers/measure"
	provider_handler "github.com/e-lua/demo-api-inventory-clean-architecture/cmd/api/handlers/provider"
	supply_handler "github.com/e-lua/demo-api-inventory-clean-architecture/cmd/api/handlers/supply"
	warehouse_handler "github.com/e-lua/demo-api-inventory-clean-architecture/cmd/api/handlers/warehouse"
)

type RouteConfig struct {
	App                 *echo.Echo
	KardexSupplyHandler *kardex_supply_handler.KardexSupplyHandler
	MeasureHandler      *measure_handler.MeasureHandler
	ProviderHandler     *provider_handler.ProviderHandler
	SupplyHandler       *supply_handler.SupplyHandler
	WarehouseHandler    *warehouse_handler.WarehouseHandler
}

func (router *RouteConfig) Start() {

	router.App.Use(middleware.Logger())
	router.App.Use(middleware.Recover())

	router.V1()

	PORT := "6500"

	handler := cors.AllowAll().Handler(router.App)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
