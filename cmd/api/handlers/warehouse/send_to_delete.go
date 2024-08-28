package warehouse

import (
	"github.com/labstack/echo/v4"

	response_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/response"
)

func (wh *WarehouseHandler) SendToDelete(c echo.Context) error {

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

	//Send to the service
	code_err, err := wh.WarehouseService.SendToDelete(idwarehouse)
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
