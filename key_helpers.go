package events

import "bytes"

const EventNameDelimiter = '-'

func GetEventName(key []byte) string {
	pos := bytes.IndexByte(key, EventNameDelimiter)

	if pos == -1 {
		return string(key)
	} else {
		return string(key[0:pos])
	}
}

func buildEventMessageKey(eventName string, extraData []byte) []byte {
	return append(
		append(
			[]byte(eventName),
			EventNameDelimiter,
		),
		extraData...,
	)
}
