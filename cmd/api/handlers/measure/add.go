package provider

import (
	"github.com/labstack/echo/v4"

	measure_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/measure"
	response_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/response"
)

func (mh *MeasureHandler) Add(c echo.Context) error {

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

	//Inicilization
	var input_measure *measure_model.Measure

	//Bind the model
	err := c.Bind(&input_measure)
	if err != nil {
		return c.JSON(400, &response_model.Response{
			Error: response_model.Error{
				Code:   9451,
				Detail: "check the structure or the type of the value, details: " + err.Error(),
			},
			Data: ""})
	}

	//Send to the service
	code_err, err := mh.MeasureService.Add(input_measure)
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
