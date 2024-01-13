package shared

type UserArgs struct {
	Separator     string
	NoLengthLimit bool
}

type SubtitleChunk struct {
	StartTimeMs int
	EndTimeMs   int
	Lines       []string
}
