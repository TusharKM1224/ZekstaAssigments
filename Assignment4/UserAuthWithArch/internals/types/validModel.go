package types


type ValidateModel struct{
	Email string `binding:"required"`
	Otp string `binding:"required"`
}