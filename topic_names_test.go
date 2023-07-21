package events

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTopics(t *testing.T) {
	topics := GetTopics()

	for index, topicName := range topics {
		if topicName == "" {
			t.Fatalf(`Wrong topic name GetTopics[%d] = %s`, index, topicName)
		}
	}
}

func TestGetPartitionsByTopicName(t *testing.T) {
	assert.Equal(t, 1, GetPartitionsByTopicName(MetaEventsTopic))
	assert.Equal(t, 1, GetPartitionsByTopicName(AuthorizedUsersTopic))
	assert.Equal(t, 2, GetPartitionsByTopicName(RawScoresTopic))
	assert.Equal(t, 6, GetPartitionsByTopicName(ScoresChangesFeedTopic))
}
