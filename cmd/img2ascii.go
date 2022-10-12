package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/davidpalves/img2ascii/converter"
	"github.com/spf13/cobra"
)

var img2asciiCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts a given image to ASCII Art",
	Run: func(cmd *cobra.Command, args []string) {
		width, _ := cmd.Flags().GetInt("width")

		imageSize := converter.ImageSize{
			Width: width,
		}

		filePath, _ := cmd.Flags().GetString("path")
		urlPath, _ := cmd.Flags().GetString("url")

		if strings.TrimSpace(filePath) == "" && strings.TrimSpace(urlPath) == "" {
			cmd.Usage()
			return
		}

		if filePath != "" {
			img := converter.ImageFileSystem{
				FilePath: filePath,
				Image:    imageSize,
			}
			result, err = img.ConvertImageToASCII()
		} else if urlPath != "" {
			img := converter.ImageURL{
				UrlPath: urlPath,
				Image:   imageSize,
			}
			result, err = img.ConvertImageToASCII()
		}

		if err != nil {
			log.Panicf("Could not convert image to ASCII")
		}

		fmt.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(img2asciiCmd)
	img2asciiCmd.Flags().Int("width", 80, "Width of the output image")
}
