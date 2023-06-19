package events

const TopicsCount = 6

const MetaEventsTopic = "meta-events"
const DisciplinesTopic = "disciplines"
const RawLessonsTopic = "raw-lessons"
const RawScores = "raw-scores"
const ScoresChangesFeedTopic = "scores-changes-feed"
const AuthorizedUsersTopic = "authorized-users"

func GetTopics() [TopicsCount]string {
	return [TopicsCount]string{
		MetaEventsTopic,
		DisciplinesTopic,
		RawLessonsTopic,
		RawScores,
		ScoresChangesFeedTopic,
		AuthorizedUsersTopic,
	}
}
