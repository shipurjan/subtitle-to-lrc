package main

import (
	"errors"
	"fmt"
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

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func main() {
	files, err := os.ReadDir("./converter")
	if err != nil {
		fmt.Println("[ERROR]", err)
	}
	converter_folders := filter(files, func(file os.DirEntry) bool {
		return file.IsDir() && file.Name() != "shared"
	})

	directory_names := make([]string, 0, len(converter_folders))
	for _, folder := range converter_folders {
		directory_names = append(directory_names, folder.Name())
	}
	allowed_extensions := strings.Join(directory_names, ", ")

	app := &cli.App{
		Name:    "subtitle-to-lrc",
		Version: version,
		Usage:   "Convert subtitle files to .lrc format",
		UsageText: "subtitle-to-lrc [options] <input-file> [output-file]\n" +
			"\n" +
			"<input-file> - the file must have an allowed subtitle extension (" + allowed_extensions + ")\n" +
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
				return err
			}

			// Verify if the provided subtitle file is valid for conversion
			subtitle_file, subtitle_extension, err := utils.ReadSubtitleFile(input_abs_path)
			if err != nil {
				return err
			}

			var output_path string
			if len(cCtx.Args().Get(1)) == 0 {
				output_path = strings.TrimSuffix(input_path, filepath.Ext(input_path)) + ".lrc"
			} else {
				output_path = cCtx.Args().Get(1)
			}

			output_abs_path, err := filepath.Abs(output_path)
			if err != nil {
				return err
			}

			// Convert subtitle file to lrc
			lyrics_file, err := converter.ConvertToLyricsFile(subtitle_file, subtitle_extension, args)
			if err != nil {
				return err
			}

			if err := writeLyricsFile(lyrics_file, output_abs_path); err != nil {
				return err
			}
			fmt.Println("[OK]", output_abs_path)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("[ERROR]", err)
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
