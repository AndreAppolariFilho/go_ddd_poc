package product

import (
	"errors"

	"github.com/AndreAppolariFilho/ddd-go/aggregate"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound      = errors.New("No such products")
	ErrProductAlreadyExists = errors.New("there is already such an product")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
