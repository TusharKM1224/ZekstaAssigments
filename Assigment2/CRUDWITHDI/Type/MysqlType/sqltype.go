package Type

type TableSchema struct {
	ID  uint16 `gorm:primarykey`
	Name string `gorm:size:110;not null`
	Email string `gorm:Unique; not null`

}