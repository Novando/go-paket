package common

import "github.com/google/uuid"

func UUIDv7() uuid.UUID {
	id, _ := uuid.NewV7()
	return id
}
