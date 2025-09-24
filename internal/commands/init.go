package commands

import (
	"fmt"

	"labit/internal/core"
)

func Init() error {
	if err := core.InitRepository(); err != nil {
		return err
	}

	fmt.Println("Initialized empty Labit repository in .labit/")
	return nil
}
