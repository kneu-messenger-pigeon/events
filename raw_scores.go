package events

import "time"

const ScoreEventName = "ScoreEvent"

type ScoreEvent struct {
	// score id can be replaced with discipline_id + student_id + lesson_id OR student_id + lesson_id;
	Id uint

	StudentId    uint
	LessonId     uint
	LessonPart   uint8
	DisciplineId uint
	Year         int
	Semester     uint8

	ScoreValue
	UpdatedAt time.Time
	SyncedAt  time.Time
}

type ScoreValue struct {
	Value     float32
	IsAbsent  bool
	IsDeleted bool
}
