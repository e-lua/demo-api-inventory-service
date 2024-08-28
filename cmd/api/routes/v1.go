package route

import (
	middlewares "github.com/e-lua/demo-api-inventory-clean-architecture/cmd/api/middlewares"
)

func (router *RouteConfig) V1() {

	version_1 := router.App.Group("/v1")

	//WAREHOUSE
	router_warehouse := version_1.Group("/warehouse", middlewares.Auth)
	router_warehouse.POST("", router.WarehouseHandler.Add)
	router_warehouse.GET("", router.WarehouseHandler.GetAll)
	router_warehouse.GET("/trash", router.WarehouseHandler.GetTrash)
	router_warehouse.PUT("/:idwarehouse", router.WarehouseHandler.Update)
	router_warehouse.PUT("/recover/:idwarehouse", router.WarehouseHandler.RecoverFromTrash)
	router_warehouse.DELETE("/:idwarehouse", router.WarehouseHandler.SendToDelete)

	//PROVIDER
	router_provider := version_1.Group("/provider", middlewares.Auth)
	router_provider.POST("", router.ProviderHandler.Add)
	router_provider.GET("", router.ProviderHandler.GetAll)
	router_provider.GET("/trash", router.ProviderHandler.GetTrash)
	router_provider.PUT("/:idprovider", router.ProviderHandler.Update)
	router_provider.PUT("/recover/:idprovider", router.ProviderHandler.RecoverFromTrash)
	router_provider.DELETE("/:idprovider", router.ProviderHandler.SendToDelete)

	//MEASURE
	router_measure := version_1.Group("/measure", middlewares.Auth)
	router_measure.POST("", router.MeasureHandler.Add)
	router_measure.GET("", router.MeasureHandler.GetAll)
	router_measure.DELETE("/:idmeasure", router.MeasureHandler.Delete)

	//SUPPLY
	router_supply := version_1.Group("/supply", middlewares.Auth)
	router_supply.POST("", router.SupplyHandler.Add)
	router_supply.GET("", router.SupplyHandler.GetAll)
	router_supply.GET("/check-sku", router.SupplyHandler.CheckSku)
	router_supply.GET("/trash", router.SupplyHandler.GetTrash)
	router_supply.PUT("/:idsupply", router.SupplyHandler.Update)
	router_supply.PUT("/recover/:idsupply", router.SupplyHandler.RecoverFromTrash)
	router_supply.DELETE("/:idsupply", router.SupplyHandler.SendToDelete)

	//KARDEX SUPPLY
	router_kardex_supply := version_1.Group("/kardex-supply", middlewares.Auth)
	router_kardex_supply.POST("", router.KardexSupplyHandler.Add)
	router_kardex_supply.GET("", router.KardexSupplyHandler.GetAll)
	router_kardex_supply.PUT("/:idkardex", router.KardexSupplyHandler.Update)
	router_kardex_supply.DELETE("/:idkardex", router.KardexSupplyHandler.Delete)
}
