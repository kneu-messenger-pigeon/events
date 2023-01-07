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
	Year                              int
	CurrentSecondaryDatabaseDatetime  time.Time
	PreviousSecondaryDatabaseDatetime time.Time
}

type SecondaryDbScoreProcessedEvent struct {
	Year                              int
	CurrentSecondaryDatabaseDatetime  time.Time
	PreviousSecondaryDatabaseDatetime time.Time
}

type SecondaryDbDeletedScoreProcessedEvent struct {
	Year                              int
	CurrentSecondaryDatabaseDatetime  time.Time
	PreviousSecondaryDatabaseDatetime time.Time
}
