package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "helloの短い説明",
	Long:  `helloの長い説明`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ハロー")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
