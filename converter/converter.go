package converter

import (
	"errors"

	"github.com/shipurjan/subtitle-to-lrc/converter/srt"
	"github.com/shipurjan/subtitle-to-lrc/converter/vtt"
)

func ConvertToLyricsFile(subtitle_file []byte, extension string) ([]byte, error) {
	converters := map[string]func([]byte) ([]byte, error){
		"vtt": vtt.ConvertToLrc,
		"srt": srt.ConvertToLrc,
	}

	if converter, ok := converters[extension]; ok {
		return ConvertToLrcFile(converter, subtitle_file)
	}
	return nil, errors.New("unsupported subtitle file extension: " + extension)
}

func ConvertToLrcFile(convert_function func([]byte) ([]byte, error), subtitle_file []byte) ([]byte, error) {
	return convert_function(subtitle_file)
}
