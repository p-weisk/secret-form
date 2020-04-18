package answer

import (
	"github.com/google/uuid"
	"time"
)

type Answer struct {
	ID			string
	Data		string
	Timestamp	time.Time
	Form		uuid.UUID
}
