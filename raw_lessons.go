package events

import "time"

const LessonEventName = "LessonEvent"

type LessonEvent struct {
	Id           uint
	DisciplineId uint
	TypeId       uint8
	Date         time.Time
	Year         int
	Semester     uint8
	IsDeleted    bool
}
