package events

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserAuthorizedEvent_JSON(t *testing.T) {
	t.Run("check GenderToStringMap", func(t *testing.T) {
		var genderString string
		var ok bool

		genderString, ok = GenderToStringMap[UnknownGender]
		assert.True(t, ok)
		assert.Empty(t, genderString)

		genderString, ok = GenderToStringMap[Male]
		assert.True(t, ok)
		assert.NotEmpty(t, genderString)

		genderString, ok = GenderToStringMap[Female]
		assert.True(t, ok)
		assert.NotEmpty(t, genderString)

	})

	for genderId, genderString := range GenderToStringMap {
		t.Run("gender_"+genderString, func(t *testing.T) {
			event := UserAuthorizedEvent{
				Client:       "123",
				ClientUserId: "9239wqeeqw",
				StudentId:    929292,
				LastName:     "Іванов",
				FirstName:    "Петро",
				MiddleName:   "Петрович",
				Gender:       genderId,
			}

			jsonData, err := json.Marshal(event)
			assert.NoError(t, err)
			assert.NotEmpty(t, jsonData)

			encodedEvent := UserAuthorizedEvent{}

			err = json.Unmarshal(jsonData, &encodedEvent)
			assert.NoError(t, err)

			assert.Equal(t, event, encodedEvent)
		})
	}

}
