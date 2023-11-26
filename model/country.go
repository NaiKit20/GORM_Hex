package model

type Country struct {
	ID   uint   `gorm:"primaryKey;column:idx"`
	Name string `gorm:"type:varchar(255)"`
}

// TableName specifies the custom table name for the Country struct.
func (Country) TableName() string {
	// Table name in database
	return "country"
}