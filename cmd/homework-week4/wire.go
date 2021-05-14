// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"homework-week4/internal/biz"
	"homework-week4/internal/conf"
	"homework-week4/internal/data"
	"homework-week4/internal/server"
	"homework-week4/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
