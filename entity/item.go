//Package entities holds all the entities that are shared accross sub-domains

package entity

import "github.com/google/uuid"

// Item is an entity that represents a Item in all domains
type Item struct {
	//Id an identifier of the entity
	ID          uuid.UUID
	Name        string
	Description string
}
