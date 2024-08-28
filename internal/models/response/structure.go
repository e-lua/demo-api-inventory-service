package response

type Error struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

type Jwt struct {
	JWT       string `json:"jwt"`
	Rol       int    `json:"rol"`
	Full_name string `json:"fullName"`
	Color     string `json:"color"`
	Email     string `json:"email"`
}

type Auth struct {
	Id        string `json:"id"`
	Full_name string `json:"fullName"`
	Rol       int    `json:"rol"`
	Country   int    `json:"country"`
}
