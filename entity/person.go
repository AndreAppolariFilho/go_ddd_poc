//Package entities holds all the entities that are shared accross sub-domains

package entity

import "github.com/google/uuid"

// Person is an entity that represents a Person in all domains
type Person struct {
	//Id an identifier of the entity
	ID   uuid.UUID
	Name string
	Age  int
}
