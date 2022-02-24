//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeEvent() (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}


 (see (https://github.com/golang/tools/blob/master/gopls/doc/settings.md#buildflags-string).
Otherwise, see the troubleshooting guidelines for help investigating (https://github.com/golang/tools/blob/master/gopls/doc/troubleshooting.md).
