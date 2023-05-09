package project

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdNew = &cobra.Command{
	Use:   "new",
	Short: "New project.",
	Long:  "",
	Run:   run,
}

func run(_ *cobra.Command, args []string) {

	fmt.Println("new")
	fmt.Println(args)
}
