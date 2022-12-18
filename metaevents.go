package events

import "time"

const CurrentYearEventName = "CurrentYearEvent"
const SecondaryDbLoadedEventName = "SecondaryDbLoadedEvent"
const SecondaryDbLessonProcessedEventName = "SecondaryDbLessonProcessedEvent"
const SecondaryDbScoreBulkProcessedEventName = "SecondaryDbScoreBulkProcessedEvent"

type CurrentYearEvent struct {
	Year uint
}

type SecondaryDbLoadedEvent struct {
	Year                              uint
	CurrentSecondaryDatabaseDatetime  time.Time
	PreviousSecondaryDatabaseDatetime time.Time
}

type SecondaryDbLessonProcessedEvent struct {
	CurrentSecondaryDatabaseDatetime  time.Time
	PreviousSecondaryDatabaseDatetime time.Time
}

type SecondaryDbScoreBulkProcessedEvent struct {
	ScoreIdMinOrEqual                 uint
	ScoreIdMAx                        uint
	CurrentSecondaryDatabaseDatetime  time.Time
	PreviousSecondaryDatabaseDatetime time.Time
}
