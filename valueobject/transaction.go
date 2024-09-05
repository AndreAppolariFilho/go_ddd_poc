package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is a valueobject because it has no identifier and it's unmutable
type Transaction struct {
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
