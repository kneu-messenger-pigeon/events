package events

const DisciplineEventName = "DisciplineEvent"

type DisciplineEvent struct {
	Discipline
	Year int
}

type Discipline struct {
	Id   uint
	Name string
}
