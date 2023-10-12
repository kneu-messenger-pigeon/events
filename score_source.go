package events

type ScoreSource int8

const (
	Unknown ScoreSource = iota
	Realtime
	Secondary
)
