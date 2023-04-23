package events

import (
	"encoding/json"
)

const UserAuthorizedEventName = "UserAuthorizedEvent"

type UserAuthorizedEvent struct {
	Client       string
	ClientUserId string
	StudentId    uint

	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	Gender     Gender `json:"gender"`
}

type Gender int8

const (
	UnknownGender Gender = iota
	Male
	Female
)

var GenderToStringMap = map[Gender]string{
	UnknownGender: "",
	Male:          "male",
	Female:        "female",
}

func (gender *Gender) String() string {
	return GenderToStringMap[*gender]
}

func GenderFromString(GenderInputString string) (gender Gender) {
	for GenderCandidate, GenderCandidateString := range GenderToStringMap {
		if GenderInputString == GenderCandidateString {
			gender = GenderCandidate
			break
		}
	}
	return
}

// MarshalJSON marshals the enum as a quoted json string
func (gender Gender) MarshalJSON() ([]byte, error) {
	return json.Marshal(gender.String())
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (gender *Gender) UnmarshalJSON(b []byte) error {
	var genderJsonString string
	err := json.Unmarshal(b, &genderJsonString)
	if err == nil {
		*gender = GenderFromString(genderJsonString)
	}

	return err
}
