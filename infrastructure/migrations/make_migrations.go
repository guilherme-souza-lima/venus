package migrations

import (
	"gorm.io/gorm"
	"ssVenus/infrastructure/model"
)

type DatabaseMakeMigrations struct {
	postgres *gorm.DB
}

func NewDatabaseMakeMigrations(postgres *gorm.DB) DatabaseMakeMigrations {
	return DatabaseMakeMigrations{postgres}
}

func (ref DatabaseMakeMigrations) MakeMigrations() {
	err := ref.postgres.Migrator().AutoMigrate(&model.GamesModel{})
	if err != nil {
		panic("Error Migrating Contacts Table Games")
	}
}
