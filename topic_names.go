package events

const TopicsCount = 6

const MetaEventsTopic = "meta-events"
const DisciplinesTopic = "disciplines"
const RawLessonsTopic = "raw-lessons"
const RawScoresTopic = "raw-scores"
const ScoresChangesFeedTopic = "scores-changes-feed"
const AuthorizedUsersTopic = "authorized-users"

func GetTopics() [TopicsCount]string {
	return [TopicsCount]string{
		MetaEventsTopic,
		DisciplinesTopic,
		RawLessonsTopic,
		RawScoresTopic,
		ScoresChangesFeedTopic,
		AuthorizedUsersTopic,
	}
}

func GetPartitionsByTopicName(topicName string) int {
	switch topicName {
	case RawScoresTopic:
		return 2
	case ScoresChangesFeedTopic:
		return 6
	default:
		return 1
	}
}
