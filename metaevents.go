package events

import "time"

type CurrentYearEvent struct {
	Year int
}

type SecondaryDbLoadedEvent struct {
	CurrentSecondaryDatabaseDatetime  time.Time
	PreviousSecondaryDatabaseDatetime time.Time
}

type SecondaryDbScoreBulkProcessedEvent struct {
	LessonIdMin                       int
	LessonIdMAx                       int
	CurrentSecondaryDatabaseDatetime  time.Time
	PreviousSecondaryDatabaseDatetime time.Time
}
