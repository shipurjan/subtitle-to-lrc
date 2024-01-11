package shared

type UserArgs struct {
	Separator string
}

type SubtitleChunk struct {
	StartTimeMs int
	EndTimeMs   int
	Lines       []string
}
