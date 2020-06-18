// +build !wireinject

package wire

import (
	"github.com/google/wire"
)

var BookDaoSet = wire.NewSet()

func InitMongoConfig() {
}