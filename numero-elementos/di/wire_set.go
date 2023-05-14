//

//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"hackerrank/numero-elementos/ctx"
	"hackerrank/repositories"
)

var stdSet = wire.NewSet(
	repositories.NewCleanList,
	wire.Bind(new(ctx.CleanList), new(*repositories.CleanList)),
	ctx.NewHandler,
)
