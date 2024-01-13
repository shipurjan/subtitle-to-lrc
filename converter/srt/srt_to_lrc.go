package srt

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/shipurjan/subtitle-to-lrc/converter/shared"
)

func ConvertToChunks(srt_file []string) ([]shared.SubtitleChunk, error) {
	body_chunks := SplitBodyIntoChunks(srt_file)

	return body_chunks, nil
}

func SplitBodyIntoChunks(body []string) []shared.SubtitleChunk {
	chunks := make([]shared.SubtitleChunk, 0)
	chunk_candidate := make([]string, 0)

	for _, line := range body {
		if line == "" {
			// Get rid of comments etc.
			if IsChunkCandidateValid(chunk_candidate) {
				chunks = append(chunks, PostProcessChunk(chunk_candidate))
			}
			chunk_candidate = make([]string, 0)
		} else {
			chunk_candidate = append(chunk_candidate, line)
		}
	}

	// Add the last chunk if EOF is detected
	if len(chunk_candidate) > 0 && IsChunkCandidateValid(chunk_candidate) {
		chunks = append(chunks, PostProcessChunk(chunk_candidate))
	}

	return chunks
}

func IsChunkCandidateValid(chunk_candidate []string) bool {
	for _, line := range chunk_candidate {
		if strings.Contains(line, "-->") {
			return true
		}
	}
	return false
}

func PostProcessChunk(validated_chunk_candidate []string) shared.SubtitleChunk {
	var newChunk []string
	for i, line := range validated_chunk_candidate {
		if strings.Contains(line, "-->") {
			newChunk = validated_chunk_candidate[i:]
			break
		}
	}
	start_time, end_time, err := GetChunkTimes(newChunk[0])
	if err != nil {
		log.Fatal(err)
	}

	return shared.SubtitleChunk{StartTimeMs: start_time, EndTimeMs: end_time, Lines: newChunk[1:]}
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
	secondsAndMilliseconds := strings.Split(parts[2], ",")
	seconds, _ := strconv.Atoi(secondsAndMilliseconds[0])
	milliseconds, _ := strconv.Atoi(secondsAndMilliseconds[1])

	totalMilliseconds := (hours * 3600000) + (minutes * 60000) + (seconds * 1000) + milliseconds
	return totalMilliseconds
}
