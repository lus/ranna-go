package snippets

import (
	"log"
	"time"
)

func parseTimestamp(raw string) time.Time {
	time, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		log.Println(err)
	}
	return time
}
