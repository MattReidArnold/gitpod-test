package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var greetingCmd = &cobra.Command{
	Use: "greeting",
	Short: "greet users",
	Run: runGreeting,
}

func init() {
	rootCmd.AddCommand(greetingCmd)
}

func runGreeting(cmd *cobra.Command, args []string) {
	fmt.Println("Hello, world")
	fmt.Printf("Good bye")
}