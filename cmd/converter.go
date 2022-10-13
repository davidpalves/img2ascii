package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/davidpalves/img2ascii/converter"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(converterCmd)
	converterCmd.Flags().Int("width", 80, "Width of the output image")
	converterCmd.Flags().String("output", "", "Output ASCII to file")
}

var converterCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts a given image to ASCII Art",
	Run: func(cmd *cobra.Command, args []string) {
		width, _ := cmd.Flags().GetInt("width")

		imageSize := converter.ImageSize{
			Width: width,
		}

		filePath, _ := cmd.Flags().GetString("path")
		urlPath, _ := cmd.Flags().GetString("url")
		outputPath, _ := cmd.Flags().GetString("output")

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

		if strings.TrimSpace(outputPath) != "" {
			converter.OutputASCIIToFile([]byte(result), outputPath)
		} else {
			fmt.Print(string(result))
		}

	},
}
