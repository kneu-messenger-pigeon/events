package events

const DisciplineEventName = "DisciplineEvent"

type DisciplineEvent struct {
	Id   uint
	Name string
	Year int
}
