package types


type LoginModel struct{
	Email string `binding:"required"`
	Pass string	`binding:"required"`
}