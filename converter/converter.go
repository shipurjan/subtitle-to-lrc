package converter

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/shipurjan/subtitle-to-lrc/converter/shared"
	"github.com/shipurjan/subtitle-to-lrc/converter/vtt"
)

func ConvertToLyricsFile(subtitle_file []string, extension string, args shared.UserArgs) ([]string, error) {
	converters := map[string]func([]string) ([]shared.SubtitleChunk, error){
		"vtt": vtt.ConvertToChunks,
	}

	if Converter, ok := converters[extension]; ok {
		chunks, err := Converter(subtitle_file)
		if err != nil {
			return nil, err
		}
		validated_chunks, err := ValidateAndPrettifyChunks(chunks)
		if err != nil {
			return nil, err
		}
		return ConvertSubtitleChunksToLyricsFile(validated_chunks, args)
	}
	return nil, errors.New("unsupported subtitle file extension: " + extension)
}

func ValidateAndPrettifyChunks(chunks []shared.SubtitleChunk) ([]shared.SubtitleChunk, error) {
	// Sort in chronological order
	sort.Slice(chunks, func(i, j int) bool {
		return chunks[i].StartTimeMs < chunks[j].StartTimeMs
	})

	// Compare the end time of the current chunk with the start time of the next chunk
	for i := 0; i < len(chunks)-1; i++ {
		currentChunk := &chunks[i]
		nextChunk := &chunks[i+1]

		if nextChunk.StartTimeMs <= currentChunk.EndTimeMs {
			currentChunk.EndTimeMs = nextChunk.StartTimeMs
		}
	}

	last_chunk := chunks[len(chunks)-1]
	if last_chunk.EndTimeMs > 3599999 {
		return nil, errors.New("subtitle file is too long; maximum supported length of a .lrc file is 59:59.99")
	}

	return chunks, nil
}

func ConvertSubtitleChunksToLyricsFile(chunks []shared.SubtitleChunk, args shared.UserArgs) ([]string, error) {
	lrc_lines := make([]string, 0)

	for i := 0; i < len(chunks)-1; i++ {
		currentChunk := &chunks[i]
		nextChunk := &chunks[i+1]

		if currentChunk.EndTimeMs == nextChunk.StartTimeMs {
			lrc_lines = append(lrc_lines, GetLrcLine(currentChunk, args))
		} else {
			lrc_lines = append(lrc_lines, GetLrcLine(currentChunk, args))
			lrc_lines = append(lrc_lines, MillisecondsToTimestamp(currentChunk.EndTimeMs))
		}
	}

	// Last chunk edge case
	currentChunk := &chunks[len(chunks)-1]
	lrc_lines = append(lrc_lines, GetLrcLine(currentChunk, args))
	lrc_lines = append(lrc_lines, MillisecondsToTimestamp(currentChunk.EndTimeMs))

	return lrc_lines, nil
}

func MillisecondsToTimestamp(milliseconds int) string {
	minutes := milliseconds / 60000
	seconds := (milliseconds / 1000) % 60
	millis := (milliseconds % 1000) / 10

	return fmt.Sprintf("[%02d:%02d.%02d]", minutes, seconds, millis)
}

func GetLrcLine(chunk *shared.SubtitleChunk, args shared.UserArgs) string {
	return MillisecondsToTimestamp(chunk.StartTimeMs) + strings.Join(chunk.Lines, args.Separator)
}
