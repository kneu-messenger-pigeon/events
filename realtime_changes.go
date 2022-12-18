package events

import "time"

const LessonCreateEventName = "LessonCreateEvent"
const LessonEditEventName = "LessonEditEvent"
const LessonDeletedEventName = "LessonDeletedEvent"
const ScoreEditEventName = "ScoreEditEventName"

type LessonCreateEvent struct {
	// n=4 and action:insert
	ActionNumber int8      `form:"n"` // default 4
	SessionId    string    `form:"sesID"`
	DisciplineId uint      `form:"prt"`
	TypeId       int       `form:"tzn"`
	Semester     int8      `form:"hlf"`
	Date         time.Time `form:"date_z" time_format:"31.01.2006" time_utc:"0"`
	TeacherId    uint      `form:"teacher"`
}

type LessonEditEvent struct {
	// n=4 and 	action: edit
	LessonId     uint      `form:"prti"`
	ActionNumber int8      `form:"n"` // default 4
	SessionId    string    `form:"sesID"`
	DisciplineId uint      `form:"prt"`
	TypeId       int       `form:"tzn"`
	Semester     int8      `form:"hlf"`
	Date         time.Time `form:"date_z" time_format:"31.01.2006" time_utc:"0"`
	TeacherId    uint      `form:"teacher"`
}

type LessonDeletedEvent struct {
	// n = 3
	ActionNumber int8   `form:"n"` // default 3
	SessionId    string `form:"sesID"`
	LessonId     uint   `form:"prti"`
	DisciplineId uint   `form:"prt"`
	Semester     int8   `form:"hlf"`
}

type ScoreEditEvent struct {
	// n = 4 and have keys "^st[0-9]+_(1|2)_[0-9]+$" st119906_2-2616031
	ActionNumber int8      `form:"n"` // default 4
	SessionId    string    `form:"sesID"`
	LessonId     uint      `form:"prti"`
	DisciplineId uint      `form:"prt"`
	Semester     int8      `form:"hlf"`
	Date         time.Time `form:"d2" time_format:"31.01.2006" time_utc:"0"`
	Scores       map[int]string
	/* map from form:
	st119905_1-2616031:
	st119905_2-2616031: нб/нп
	st119906_1-2616031:
	st119909_1-2616031:	6
	st119909_2-2616031:
	st119910_1-2616031:
	st119910_2-2616031:
	*/

}
