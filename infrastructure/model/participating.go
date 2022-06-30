package model

type ParticipatingModel struct {
	ID   int    `gorm:"PrimaryKey"`
	User string `gorm:"column:user_id"`
	Game string `gorm:"column:game_id"`
	Nick string `gorm:"column:nick"`
}

func (ref *ParticipatingModel) TableName() string {
	return "participating"
}
