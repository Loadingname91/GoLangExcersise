package domain

// DO not use fmt in lib pkgs
import (
	"errors"
	"fmt"
	"log"
)

type MapStore struct {
	// An in mermory store with a map
	// Use Customer.ID as the key of the map
	store map[string]Customer
}

// Factory Method gives a new Instance of MapStore
func NewMapStore() *MapStore {
	return &MapStore{
		store: make(map[string]Customer),
	}
}

// Perform create operation for Customer
func (m *MapStore) Create(c Customer) error {
	if _, ok := m.store[c.ID]; ok {
		return errors.New("key already exists")
	}
	m.store[c.ID] = c
	log.Printf("Data with ID %s inserted successfully \n", c.ID)
	//fmt.Printf("Data with ID %s inserted successfully \n", c.ID)
	return nil
}

// Perform update operation for Customer
func (m *MapStore) Update(ID string, c Customer) error {
	if _, ok := m.store[c.ID]; !ok {
		return errors.New("key does not exist")
	}
	m.store[c.ID] = c

	fmt.Printf("Data with ID %s updated successfully for the Name %s \n", c.ID, c.Name)
	return nil
}

// Perform delete operation for Customer
func (m *MapStore) Delete(ID string) error {
	if _, ok := m.store[ID]; !ok {
		return errors.New("key does not exist")
	}
	delete(m.store, ID)
	fmt.Printf("Data with ID %s deleted successfully \n", ID)
	return nil
}

// Perform get operation for Customer
func (m *MapStore) GetByID(ID string) (c Customer, err error) {
	if _, ok := m.store[ID]; !ok {
		return c, errors.New("key does not exist")
	}
	return m.store[ID], nil
}

// Perform get all operation for Customer
func (m *MapStore) GetAll() ([]Customer, error) {
	if len(m.store) == 0 {
		return nil, errors.New("Customer store map is empty")
	}
	var c []Customer
	for _, v := range m.store {
		c = append(c, v)
	}
	return c, nil
}
