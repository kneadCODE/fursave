package entity

import "time"

// Common represents common fields for entities.
type Common struct {
	ID        int64     // Unique identifier for the entity
	CreatedAt time.Time // Timestamp when the entity was created
	UpdatedAt time.Time // Timestamp when the entity was last updated
	DeletedAt time.Time // Timestamp when the entity was deleted (soft delete)
}
