package events

import "time"

const CurrentYearEventName = "CurrentYearEvent"
const SecondaryDbLoadedEventName = "SecondaryDbLoadedEvent"
const SecondaryDbLessonProcessedEventName = "SecondaryDbLessonProcessedEvent"
const SecondaryDbScoreProcessedEventName = "SecondaryDbScoreProcessedEvent"
const SecondaryDbDeletedScoreProcessedEventName = "SecondaryDbDeletedScoreProcessedEvent"

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

type SecondaryDbScoreProcessedEvent struct {
	CurrentSecondaryDatabaseDatetime  time.Time
	PreviousSecondaryDatabaseDatetime time.Time
}

type SecondaryDbDeletedScoreProcessedEvent struct {
	CurrentSecondaryDatabaseDatetime  time.Time
	PreviousSecondaryDatabaseDatetime time.Time
}
