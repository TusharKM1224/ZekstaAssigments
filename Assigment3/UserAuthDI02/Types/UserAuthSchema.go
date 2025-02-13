package Types

type User struct{
	User_id uint `gorm:primarykey`
	User_Name string `gorm:size:100; not null`
	Email string `gorm:size:100;not null`
	Phone string `gorm:not null`
	Password string `gorm:size:100;not null`
	Role_name string `gorm:not null`
}
type admin struct{
	admin_id uint `gorm:primarykey`
	User_name string `gorm:size:100`
	Password string `gorm:size:100`
}

type ValidateOtp struct{
	Email string 
	Otp string
}
