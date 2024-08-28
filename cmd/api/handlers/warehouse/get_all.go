package warehouse

import (
	"strconv"

	"github.com/labstack/echo/v4"

	response_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/response"
	warehouse_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/warehouse"
)

func (wh *WarehouseHandler) GetAll(c echo.Context) error {

	//Get the rol the Middleware
	rol := c.Get("rol").(int)
	if rol != 1 && rol != 2 {
		return c.JSON(401, &response_model.Response{
			Error: response_model.Error{
				Code:   40526,
				Detail: "invalid rol: available for Admin and CoAdmin",
			},
			Data: []*warehouse_model.Warehouse{}})
	}

	//Get the filters from the client
	id_business := c.Request().URL.Query().Get("idBusiness")
	name := c.Request().URL.Query().Get("name")

	limit_string := c.Request().URL.Query().Get("limit")
	offset_string := c.Request().URL.Query().Get("offset")
	limit, _ := strconv.Atoi(limit_string)
	offset, _ := strconv.Atoi(offset_string)

	if id_business == "" || limit_string == "" || offset_string == "" {
		return c.JSON(401, &response_model.Response{
			Error: response_model.Error{
				Code:   4052,
				Detail: "idBusiness must be sent, limit must be sent, offset must be sent",
			},
			Data: []*warehouse_model.Warehouse{}})
	}

	//Send to the service
	code_err, list_warehouse, err := wh.WarehouseService.GetAll(id_business, name, limit, offset)
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
			Data: list_warehouse})
	}

	//OK
	return c.JSON(200, &response_model.Response{
		Error: response_model.Error{
			Code:   0,
			Detail: "",
		},
		Data: list_warehouse})
}
