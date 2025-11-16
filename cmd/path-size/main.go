package main

import (
	"context"
	"log"
	"os"
	"code"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name: "path-size",
		Usage: "print size of a file or directory;",
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
	code.GetSize("../go-project-242/Makefile")
}