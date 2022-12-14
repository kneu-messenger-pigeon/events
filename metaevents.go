package events

import "time"

const CurrentYearEventName = "CurrentYearEvent"
const SecondaryDbLoadedEventName = "SecondaryDbLoadedEvent"
const SecondaryDbScoreBulkProcessedEventName = "SecondaryDbScoreBulkProcessedEvent"

type CurrentYearEvent struct {
	Year int
}

type SecondaryDbLoadedEvent struct {
	Year                              int
	CurrentSecondaryDatabaseDatetime  time.Time
	PreviousSecondaryDatabaseDatetime time.Time
}

type SecondaryDbScoreBulkProcessedEvent struct {
	LessonIdMin                       int
	LessonIdMAx                       int
	CurrentSecondaryDatabaseDatetime  time.Time
	PreviousSecondaryDatabaseDatetime time.Time
}
