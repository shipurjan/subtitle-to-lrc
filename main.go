package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/shipurjan/subtitle-to-lrc/converter"
	"github.com/shipurjan/subtitle-to-lrc/utils"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			// &cli.StringFlag{
			// 	Name:  "lang",
			// 	Value: "english",
			// 	Usage: "language for the greeting",
			// },
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.NArg() > 0 {
				// Read the argument
				filename := cCtx.Args().Get(0)

				// Convert to absolute path
				filepath, err := filepath.Abs(filename)
				if err != nil {
					log.Fatal(err)
				}

				// Verify if the provided subtitle file is valid for conversion
				subtitle_file, subtitle_extension, err := utils.ReadSubtitleFile(filepath)
				if err != nil {
					log.Fatal(err)
				}

				// Convert subtitle file to lrc
				lyrics_file, err := converter.ConvertToLyricsFile(subtitle_file, subtitle_extension)
				if err != nil {
					log.Fatal(err)
				}
				log.Println(lyrics_file)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
