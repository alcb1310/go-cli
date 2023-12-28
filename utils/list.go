package utils

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func List(c *cli.Context) error {
	_ = c
	fmt.Println("list all branches")
	return nil
}
