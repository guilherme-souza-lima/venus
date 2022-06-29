package model

type StatesModel struct {
	ID   int    `gorm:"PrimaryKey"`
	Name string `gorm:"column:name"`
}

func (ref *StatesModel) TableName() string {
	return "states"
}
