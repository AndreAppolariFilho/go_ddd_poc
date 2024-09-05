package services

import (
	"testing"

	"github.com/AndreAppolariFilho/ddd-go/aggregate"
	"github.com/google/uuid"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}
	tavern, err := NewTavern(WithOrderService(os))

	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("percy")
	if err != nil {
		t.Fatal(err)
	}

	if err = os.customers.Add(cust); err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}
	err = tavern.Order(cust.GetID(), order)

	if err != nil {
		t.Fatal(err)
	}

}
