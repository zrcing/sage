package main

import (
	"github.com/spf13/cobra"
	"github.com/zrcing/sage/cmd/sage/internal/project"
	"log"
	"os"
)

var cmd = &cobra.Command{
	Use:   "sage",
	Short: "Sage: An toolkit for Go microservices.",
	Long:  `Sage: An toolkit for Go microservices.`,
	//Version: release,
}

func init() {
	cmd.AddCommand(project.CmdNew)
}

func main() {

	// Initialize default logger to log to stderr
	log.SetFlags(0)
	log.SetPrefix("sage: ")
	log.SetOutput(os.Stderr)

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
