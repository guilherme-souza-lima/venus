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
	errFunc(err)
	err = ref.postgres.Migrator().AutoMigrate(&model.StatesModel{})
	errFunc(err)
	err = ref.postgres.Migrator().AutoMigrate(&model.ParticipatingModel{})
	errFunc(err)

	ref.postgres.FirstOrCreate(&model.StatesModel{}, &model.StatesModel{ID: 1, Name: "Green"})
	ref.postgres.FirstOrCreate(&model.StatesModel{}, &model.StatesModel{ID: 2, Name: "Yellow"})
	ref.postgres.FirstOrCreate(&model.StatesModel{}, &model.StatesModel{ID: 3, Name: "Red"})
	ref.postgres.FirstOrCreate(&model.StatesModel{}, &model.StatesModel{ID: 4, Name: "White"})
	ref.postgres.FirstOrCreate(&model.StatesModel{}, &model.StatesModel{ID: 5, Name: "Black"})
	ref.postgres.FirstOrCreate(&model.StatesModel{}, &model.StatesModel{ID: 6, Name: "Brown"})
}

func errFunc(err error) {
	if err != nil {
		panic("Error Migrating Contacts Table")
	}
}
