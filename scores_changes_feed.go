package events

const ScoreChangedEventName = "ScoreChangedEvent"

type ScoreChangedEvent struct {
	ScoreEvent
	Previous ScoreValue
}
