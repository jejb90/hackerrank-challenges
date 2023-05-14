//

//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"hackerrank/robot-en-marte/ctx"
)

func Initialize() (*ctx.Handler, error) {
	wire.Build(stdSet)

	return &ctx.Handler{}, nil
}
