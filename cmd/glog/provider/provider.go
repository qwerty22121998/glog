package provider

import (
	"github.com/qwerty22121998/glog/cmd/glog/controller"
	"github.com/qwerty22121998/glog/pkg/repository"
	"github.com/qwerty22121998/glog/pkg/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type Provider interface {
	Repository() repository.Provider
	Service() service.Provider
	Controller() controller.Provider
}

type provider struct {
	repositoryProvider repository.Provider
	serviceProvider    service.Provider
	controllerProvider controller.Provider
}

func (p *provider) Repository() repository.Provider {
	return p.repositoryProvider
}

func (p *provider) Service() service.Provider {
	return p.serviceProvider
}

func (p *provider) Controller() controller.Provider {
	return p.controllerProvider
}

func NewProvider(db *mongo.Database) Provider {
	repo := repository.NewProvider(db)
	services := service.NewProvider(repo)
	controllers := controller.NewProvider(services)
	return &provider{
		repositoryProvider: repo,
		serviceProvider:    services,
		controllerProvider: controllers,
	}
}
