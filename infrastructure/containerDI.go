package infrastructure

import (
	"gorm.io/gorm"
	"ssVenus/infrastructure/database"
	"ssVenus/infrastructure/database/postgresql"
	"ssVenus/infrastructure/migrations"
	"ssVenus/user_case/handler"
	"ssVenus/user_case/repository"
	"ssVenus/user_case/service"
)

type ContainerDI struct {
	Config                  Config
	DB                      *gorm.DB
	Migration               migrations.DatabaseMakeMigrations
	GamesRepository         repository.GamesRepository
	ParticipatingRepository repository.ParticipatingRepository
	GamesService            service.GamesService
	ParticipatingService    service.ParticipatingService
	GamesHandler            handler.GamesHandler
	ParticipatingHandler    handler.ParticipatingHandler
}

func NewContainerDI(config Config) *ContainerDI {
	container := &ContainerDI{
		Config: config,
	}

	configDB := database.Config{
		Hostname: container.Config.Host,
		Port:     container.Config.Port,
		UserName: container.Config.User,
		Password: container.Config.Password,
		Database: container.Config.Database,
	}

	container.DB = postgresql.InitGorm(&configDB)
	container.Migration = migrations.NewDatabaseMakeMigrations(container.DB)

	container.buildRepository()
	container.buildService()
	container.buildHandler()
	container.build()
	return container
}

func (c *ContainerDI) buildRepository() {
	c.GamesRepository = repository.NewUserRepository(c.DB)
	c.ParticipatingRepository = repository.NewParticipatingRepository(c.DB)
}
func (c *ContainerDI) buildService() {
	c.GamesService = service.NewGamesService(c.GamesRepository)
	c.ParticipatingService = service.NewParticipatingService(c.ParticipatingRepository)
}
func (c *ContainerDI) buildHandler() {
	c.GamesHandler = handler.NewGamesHandler(c.GamesService)
	c.ParticipatingHandler = handler.NewParticipatingHandler(c.ParticipatingService)
}

func (c *ContainerDI) build()    {}
func (c *ContainerDI) ShutDown() {}
