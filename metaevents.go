package events

type CurrentYearMessage struct {
	year int
}

type SecondaryDbLoadedMessage struct {
	currentSecondaryDatabaseDatetime  int
	previousSecondaryDatabaseDatetime int
}

type SecondaryDbScoreBulkProcessed struct {
	lessonIdMin                       int
	lessonIdMAx                       int
	currentSecondaryDatabaseDatetime  int
	previousSecondaryDatabaseDatetime int
}
