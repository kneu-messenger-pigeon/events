package events

import "time"

const CurrentYearEventName = "CurrentYearEvent"
const SecondaryDbLoadedEventName = "SecondaryDbLoadedEvent"
const SecondaryDbLessonProcessedEventName = "SecondaryDbLessonProcessedEvent"
const SecondaryDbScoreBulkProcessedEventName = "SecondaryDbScoreBulkProcessedEvent"

type CurrentYearEvent struct {
	Year int
}

type SecondaryDbLoadedEvent struct {
	Year                              int
	CurrentSecondaryDatabaseDatetime  time.Time
	PreviousSecondaryDatabaseDatetime time.Time
}

type SecondaryDbLessonProcessedEvent struct {
	CurrentSecondaryDatabaseDatetime  time.Time
	PreviousSecondaryDatabaseDatetime time.Time
}

type SecondaryDbScoreBulkProcessedEvent struct {
	ScoreIdMin                        uint
	ScoreIdMax                        uint
	CurrentSecondaryDatabaseDatetime  time.Time
	PreviousSecondaryDatabaseDatetime time.Time
}
