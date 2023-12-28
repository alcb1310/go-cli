package utils

import (
	// "fmt"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

func Create(c *cli.Context) error {
	err := createBranch(c.String("name"))
	if err != nil {
		if strings.Contains(err.Error(), "128") {
			fmt.Println("Branch already exists")
			return err
		}
		log.Fatal(err)
	}
	return err
}

func createBranch(name string) error {
	cmd := exec.Command("git", "checkout", "-b", name)
	err := cmd.Run()
	return err
}
