package vtt

import (
	"errors"
	"slices"
	"strconv"
	"strings"

	"github.com/shipurjan/subtitle-to-lrc/converter/shared"
)

func ConvertToChunks(vtt_file []string) ([]shared.SubtitleChunk, error) {
	header_separator_index := slices.Index(vtt_file, "")
	header := vtt_file[:header_separator_index]
	// Header could be used for some metadata in the future
	_ = header

	body := vtt_file[header_separator_index+1:]
	body_chunks, err := SplitBodyIntoChunks(body)
	if err != nil {
		return nil, err
	}

	return body_chunks, nil
}

func SplitBodyIntoChunks(body []string) ([]shared.SubtitleChunk, error) {
	chunks := make([]shared.SubtitleChunk, 0)
	chunk_candidate := make([]string, 0)

	for _, line := range body {
		if line == "" {
			// Get rid of comments etc.
			if IsChunkCandidateValid(chunk_candidate) {
				chunk, err := PostProcessChunk(chunk_candidate)
				if err != nil {
					return nil, err
				}
				chunks = append(chunks, chunk)
			}
			chunk_candidate = make([]string, 0)
		} else {
			chunk_candidate = append(chunk_candidate, line)
		}
	}

	// Add the last chunk if EOF is detected
	if len(chunk_candidate) > 0 && IsChunkCandidateValid(chunk_candidate) {
		chunk, err := PostProcessChunk(chunk_candidate)
		if err != nil {
			return nil, err
		}
		chunks = append(chunks, chunk)
	}

	return chunks, nil
}

func IsChunkCandidateValid(chunk_candidate []string) bool {
	for _, line := range chunk_candidate {
		if strings.Contains(line, "-->") {
			return true
		}
	}
	return false
}

func PostProcessChunk(validated_chunk_candidate []string) (shared.SubtitleChunk, error) {
	var newChunk []string
	for i, line := range validated_chunk_candidate {
		if strings.Contains(line, "-->") {
			newChunk = validated_chunk_candidate[i:]
			break
		}
	}
	start_time, end_time, err := GetChunkTimes(newChunk[0])
	if err != nil {
		return shared.SubtitleChunk{}, err
	}

	return shared.SubtitleChunk{StartTimeMs: start_time, EndTimeMs: end_time, Lines: newChunk[1:]}, nil
}

func GetChunkTimes(timeLine string) (int, int, error) {
	if !strings.Contains(timeLine, "-->") {
		return 0, 0, errors.New("invalid time line")
	}
	timestamps := strings.Split(timeLine, "-->")
	start_time := TimestampToMilliseconds(strings.Trim(timestamps[0], " "))
	end_time := TimestampToMilliseconds(strings.Trim(timestamps[1], " "))

	return start_time, end_time, nil
}

func TimestampToMilliseconds(timestamp string) int {
	parts := strings.Split(timestamp, ":")
	hours, _ := strconv.Atoi(parts[0])
	minutes, _ := strconv.Atoi(parts[1])
	secondsAndMilliseconds := strings.Split(parts[2], ".")
	seconds, _ := strconv.Atoi(secondsAndMilliseconds[0])
	milliseconds, _ := strconv.Atoi(secondsAndMilliseconds[1])

	totalMilliseconds := (hours * 3600000) + (minutes * 60000) + (seconds * 1000) + milliseconds
	return totalMilliseconds
}
