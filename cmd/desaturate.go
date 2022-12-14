package cmd

import (
	"log"
	"strings"

	"github.com/davidpalves/img2ascii/converter"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(desaturateCmd)

	desaturateCmd.Flags().String("output", "grayscale.png", "Path to the desaturated image. E.g.: samples/image.png")

}

var desaturateCmd = &cobra.Command{
	Use:   "desaturate",
	Short: "Creates a desaturated version of an image",
	Run: func(cmd *cobra.Command, args []string) {
		imageSize := converter.ImageSize{}

		filePath, _ := cmd.Flags().GetString("path")
		urlPath, _ := cmd.Flags().GetString("url")
		outputPath, _ := cmd.Flags().GetString("output")

		if strings.TrimSpace(filePath) == "" || strings.TrimSpace(urlPath) == "" {
			cmd.Usage()
			return
		}

		if filePath != "" {
			img := converter.ImageFileSystem{
				FilePath: filePath,
				Image:    imageSize,
			}
			_, err = img.DesaturateImage(outputPath)
		} else if urlPath != "" {
			img := converter.ImageURL{
				UrlPath: urlPath,
				Image:   imageSize,
			}
			_, err = img.DesaturateImage(outputPath)
		}

		if err != nil {
			log.Fatal("Could not process image")
		}

	},
}
