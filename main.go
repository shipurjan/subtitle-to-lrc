package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/shipurjan/subtitle-to-lrc/converter"
	"github.com/shipurjan/subtitle-to-lrc/converter/shared"
	"github.com/shipurjan/subtitle-to-lrc/utils"
	"github.com/urfave/cli/v2"
)

var (
	version string
)

func main() {

	app := &cli.App{
		Name:    "subtitle-to-lrc",
		Version: version,
		Usage:   "Convert subtitle files to .lrc format",
		UsageText: "subtitle-to-lrc [options] <input-file> [output-file]\n" +
			"<input-file> - the file must have an allowed subtitle extension (e.g. .srt, .vtt)\n" +
			"[output-file] - if not provided the program will use <input-file> filename with its extension replaced by .lrc",
		Compiled:               time.Now(),
		EnableBashCompletion:   true,
		HideHelpCommand:        true,
		Authors:                []*cli.Author{{Name: "Cyprian Zdebski", Email: "cyprianz5mail@gmail.com"}},
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "separator",
				Value:   "  ",
				Aliases: []string{"s"},
				Usage: "Separator to use to join lines when input subtitle file has multiple lines;\n" +
					".lrc files can only have one subtitle line for each timestamp",
			},
			&cli.BoolFlag{
				Name:    "no-length-limit",
				Aliases: []string{"n"},
				Usage: "Disables the length limit of a .lrc file;\n" +
					"by default a .lrc file can only have a maximum length of 59:59.99\n" +
					"(some players may not support longer durations)",
			},
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.NArg() < 1 {
				return errors.New("no input file provided, see --help for more information")
			}

			args := shared.UserArgs{
				Separator:     cCtx.String("separator"),
				NoLengthLimit: cCtx.Bool("no-length-limit"),
			}

			println(args.NoLengthLimit)

			// Read the argument
			var input_path string
			if len(cCtx.Args().Get(0)) == 0 {
				return errors.New("no input file provided")
			} else {
				input_path = cCtx.Args().Get(0)
			}

			// Convert to absolute path
			input_abs_path, err := filepath.Abs(input_path)
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Reading", input_abs_path)
			// Verify if the provided subtitle file is valid for conversion
			subtitle_file, subtitle_extension, err := utils.ReadSubtitleFile(input_abs_path)
			if err != nil {
				log.Fatal(err)
			}

			var output_path string
			if len(cCtx.Args().Get(1)) == 0 {
				output_path = strings.TrimSuffix(input_path, filepath.Ext(input_path)) + ".lrc"
			} else {
				output_path = cCtx.Args().Get(1)
			}

			output_abs_path, err := filepath.Abs(output_path)
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Converting to a .lrc file format")
			// Convert subtitle file to lrc
			lyrics_file, err := converter.ConvertToLyricsFile(subtitle_file, subtitle_extension, args)
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Writing converted subtitle to", output_abs_path)
			if err := writeLyricsFile(lyrics_file, output_abs_path); err != nil {
				log.Fatal(err)
			}
			log.Println("Successfully finished")

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func writeLyricsFile(lyricsFile []string, outputFilename string) error {
	// Join the lyrics lines into a single string
	lyricsContent := strings.Join(lyricsFile, "\n")

	// Write the lyrics content to the output file
	err := os.WriteFile(outputFilename, []byte(lyricsContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
