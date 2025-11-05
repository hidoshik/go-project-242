package goproject242

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

func PathSize() {
	(&cli.Command{}).Run(context.Background(), os.Args)
}