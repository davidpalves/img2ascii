package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var result []byte
var err error

var rootCmd = &cobra.Command{
	Use:   "img2ascii",
	Short: "Convert images to ASCII art",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("url", "", "Image URL to be converted")
	rootCmd.PersistentFlags().String("path", "", "Path to image file to be converted")

	rootCmd.MarkFlagsMutuallyExclusive("url", "path")

}
