package events

const ScoreChangedEventName = "ScoreChangedEvent"

type ScoreChangedEvent struct {
	Discipline
	ScoreEvent
	Previous ScoreValue
}
