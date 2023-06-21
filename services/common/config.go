package common

import (
	"github.com/urfave/cli/v2"
)

type Config interface {
	Parse(ctx *cli.Context) ([]byte, error)
	Empty() (string, error)
	String() string
}
