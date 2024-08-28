package provider

import (
	"github.com/labstack/echo/v4"

	kardex_supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/kardex_supply"
	response_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/response"
)

func (ksh *KardexSupplyHandler) Add(c echo.Context) error {

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

	//Get the full_name form the Middleware
	full_name := c.Get("fullName").(string)

	//Inicilization
	var input_kardex_supply *kardex_supply_model.KardexSupply

	//Bind the model
	err := c.Bind(&input_kardex_supply)
	if err != nil {
		return c.JSON(400, &response_model.Response{
			Error: response_model.Error{
				Code:   9451,
				Detail: "check the structure or the type of the value, details: " + err.Error(),
			},
			Data: ""})
	}

	//Send to the service
	code_err, err := ksh.KardexSupplyService.Add(full_name, input_kardex_supply)
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
