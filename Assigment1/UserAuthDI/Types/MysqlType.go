
package Types



type UserAuthstruct struct{
	ID uint16 `gorm:primarykey`
	Name string `gorm:size:100;not null`
	Email string `gorm:unique;not null`
}