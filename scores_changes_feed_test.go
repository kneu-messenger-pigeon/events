package events

import (
	"bytes"
	"testing"
)

func TestScoreChangedEvent_GetMessageKey(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		event := ScoreChangedEvent{
			ScoreEvent: ScoreEvent{
				LessonId:  123,
				StudentId: 1,
			},
		}

		event2 := ScoreChangedEvent{
			ScoreEvent: ScoreEvent{
				LessonId:  222,
				StudentId: 12,
			},
		}

		key := event.GetMessageKey()
		if GetEventName(key) != ScoreChangedEventName {
			t.Fatalf(`Wrong event name %s is not of %s`, key, ScoreChangedEventName)
		}

		if bytes.Equal(key, event2.GetMessageKey()) {
			t.Fatalf(`Two message generates same key %s == %s`, key, event2.GetMessageKey())
		}
	})
}
