package cmd

import (
	"fmt"
	"log"

	"github.com/davidpalves/img2ascii/converter"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(Image2ASCII())

	rootCmd.PersistentFlags().String("url", "", "Image URL to be converted")
	rootCmd.PersistentFlags().String("path", "", "Path to image file to be converted")
	rootCmd.LocalFlags().Int("width", 80, "Width of the output image")

	rootCmd.MarkFlagsMutuallyExclusive("url", "path")
}
func Image2ASCII() *cobra.Command {
	return &cobra.Command{
		Use:   "convert",
		Short: "Converts a given image to ASCII Art",
		Run: func(cmd *cobra.Command, args []string) {
			width, _ := cmd.Flags().GetInt("width")

			imageSize := converter.ImageSize{
				Width: width,
			}

			filePath, _ := cmd.Flags().GetString("path")
			urlPath, _ := cmd.Flags().GetString("url")

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
}
