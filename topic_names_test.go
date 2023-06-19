package events

import "testing"

func TestGetTopics(t *testing.T) {
	topics := GetTopics()

	for index, topicName := range topics {
		if topicName == "" {
			t.Fatalf(`Wrong topic name GetTopics[%d] = %s`, index, topicName)
		}
	}
}
