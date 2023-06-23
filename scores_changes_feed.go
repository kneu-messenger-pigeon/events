package events

import "encoding/binary"

const ScoreChangedEventName = "ScoreChangedEvent"

type ScoreChangedEvent struct {
	ScoreEvent
	Previous ScoreValue
}

func (event *ScoreChangedEvent) GetMessageKey() []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(event.LessonId))
	bs = binary.LittleEndian.AppendUint32(bs, uint32(event.StudentId))

	return buildEventMessageKey(ScoreChangedEventName, bs)
}
