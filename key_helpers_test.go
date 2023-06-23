package events

import (
	"encoding/binary"
	"strings"
	"testing"
)

func TestGetEventName(t *testing.T) {
	t.Run("no_delimiter", func(t *testing.T) {
		key := []byte(CurrentYearEventName)

		if GetEventName(key) != CurrentYearEventName {
			t.Fatalf(`Wrong event name %s != %s`, GetEventName(key), CurrentYearEventName)
		}
	})

	t.Run("delimiter", func(t *testing.T) {
		key := append(
			[]byte(CurrentYearEventName),
			EventNameDelimiter,
			'1', '2',
		)

		if GetEventName(key) != CurrentYearEventName {
			t.Fatalf(`Wrong event name %s != %s`, GetEventName(key), CurrentYearEventName)
		}
	})
}

func TestBuildEventMessageKey(t *testing.T) {
	t.Run("usual", func(t *testing.T) {
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, 31415926)
		key := string(buildEventMessageKey(CurrentYearEventName, bs))

		if !strings.HasPrefix(key, CurrentYearEventName+string(EventNameDelimiter)) {
			t.Fatalf(`Wrong message key %s doesn't start with %s%d`, key, CurrentYearEventName, EventNameDelimiter)
		}
	})
}
