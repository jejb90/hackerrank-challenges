//

//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"hackerrank/repositories"
	"hackerrank/robot-en-marte/ctx"
)

var stdSet = wire.NewSet(
	repositories.NewIntruccionRepository,
	wire.Bind(new(ctx.Intruccion), new(*repositories.Intruccion)),
	ctx.NewHandler,
)
