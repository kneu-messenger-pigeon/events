package events

type CurrentYearEvent struct {
	year int
}

type SecondaryDbLoadedEvent struct {
	currentSecondaryDatabaseDatetime  int
	previousSecondaryDatabaseDatetime int
}

type SecondaryDbScoreBulkProcessedEvent struct {
	lessonIdMin                       int
	lessonIdMAx                       int
	currentSecondaryDatabaseDatetime  int
	previousSecondaryDatabaseDatetime int
}
