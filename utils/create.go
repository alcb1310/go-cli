package utils

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Create(c *cli.Context) error {
	fmt.Printf("create a new branch called %s\n", c.String("name"))
	return nil
}
