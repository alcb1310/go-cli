package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	// "github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/nexidian/gocliselect"
	"github.com/urfave/cli/v2"
)

func List(c *cli.Context) error {
	_ = c
	out, err := exec.Command("git", "branch", "-a").Output()
	if err != nil {
		log.Fatal(err)
	}

	b := strings.Split(string(out), "\n")
	// branches := []string{}

	menu := gocliselect.NewMenu("Select branch")
	menu.AddItem("Cancel", "cancel")
	menu.AddItem("Create new branch", "create")

	for _, branch := range b {
		if branch == "" {
			continue
		}
		// branches = append(branches, branch)
		menu.AddItem(branch, branch)
	}

	cmd := exec.Command("clear") // Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
	choice := menu.Display()

	switch choice {
	case "cancel":
		return nil
	case "create":
		name := getBranchName()
		err := createBranch(name)
		if err != nil {
			if strings.Contains(err.Error(), "128") {
				fmt.Println("Branch already exists")
				return err
			}
			log.Fatal(err)
		}
		return err
	default:
		fmt.Println(choice)
		return nil
	}
}

func getBranchName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter branch name: ")
	text, _ := reader.ReadString('\n')

	return strings.TrimSpace(text)
}
