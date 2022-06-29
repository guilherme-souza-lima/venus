package model

type ParticipatingModel struct {
	ID   int    `gorm:"PrimaryKey"`
	UUID string `gorm:"unique"`
	User string `gorm:"column:user_id"`
	Game string `gorm:"column:game_id"`
	Nick string `gorm:"column:nick"`
}

func (ref *ParticipatingModel) TableName() string {
	return "participating"
}
