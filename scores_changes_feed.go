package events

const ScoreChangedEventName = "ScoreChangedEvent"

type ScoreChangedEvent struct {
	ScoreEvent

	Previous struct {
		Value     float32
		IsAbsent  bool
		IsDeleted bool
	}
}
