package events

const UserAuthorizedEventName = "UserAuthorizedEvent"

type UserAuthorizedEvent struct {
	Client       string
	ClientUserId string
	StudentId    int
}
