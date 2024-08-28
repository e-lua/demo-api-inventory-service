package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	response_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/response"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		//Read variable of the environment .env.dev
		err := godotenv.Load(".env.dev")
		if err != nil {
			log.Println("Error loading file .env.dev")
		}

		//Sending the request
		resquest_http, _ := http.NewRequest("GET", os.Getenv("URL_AUTH_ENDPOINT"), nil)
		resquest_http.Header.Add("auth-token", c.Request().Header.Get("auth-token"))
		client := &http.Client{}
		response_http, _ := client.Do(resquest_http)

		//Decoding the response
		var response_auth *response_model.ResponseAuth
		error_decode := json.NewDecoder(response_http.Body).Decode(&response_auth)
		if error_decode != nil {
			return c.JSON(403, &response_model.Response{
				Error: response_model.Error{
					Code:   9459,
					Detail: "error in the auth, details: " + error_decode.Error(),
				},
				Data: ""})
		}
		if response_auth.Data.Id == "" {
			return c.JSON(403, &response_model.Response{
				Error: response_model.Error{
					Code:   9459,
					Detail: "error in the auth, details: This user does not exist",
				},
				Data: ""})
		}

		//Assigning the values
		c.Set("id", response_auth.Data.Id)
		c.Set("fullName", response_auth.Data.Full_name)
		c.Set("rol", response_auth.Data.Rol)
		c.Set("country", response_auth.Data.Country)

		//OK
		return next(c)
	}

}
