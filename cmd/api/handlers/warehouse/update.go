package warehouse

import (
	"github.com/labstack/echo/v4"

	response_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/response"
	warehouse_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/warehouse"
)

func (wh *WarehouseHandler) Update(c echo.Context) error {

	//Get the rol the Middleware
	rol := c.Get("rol").(int)
	if rol != 1 && rol != 2 {
		return c.JSON(401, &response_model.Response{
			Error: response_model.Error{
				Code:   40526,
				Detail: "invalid rol: available for Admin and CoAdmin",
			},
			Data: ""})
	}

	//Get the idwarehouse
	idwarehouse := c.Param("idwarehouse")

	//Inicilization
	var input_warehouse *warehouse_model.Warehouse

	//Bind the model
	err := c.Bind(&input_warehouse)
	if err != nil {
		return c.JSON(400, &response_model.Response{
			Error: response_model.Error{
				Code:   9451,
				Detail: "check the structure or the type of the value, details: " + err.Error(),
			},
			Data: ""})
	}

	//Send to the service
	code_err, err := wh.WarehouseService.Update(idwarehouse, input_warehouse)
	if err != nil {

		var code_http int

		switch {
		case code_err < 4999:
			code_http = 400
		case code_err > 4999:
			code_http = 500
		}

		return c.JSON(code_http, &response_model.Response{
			Error: response_model.Error{
				Code:   code_err,
				Detail: err.Error(),
			},
			Data: ""})
	}

	//OK
	return c.JSON(200, &response_model.Response{
		Error: response_model.Error{
			Code:   0,
			Detail: "",
		},
		Data: "OK"})
}
