package main

import (
	domain "customerApp/collections"
	"fmt"
)

type CustomerController struct {
	// Embeding the CustomerStore interface
	store domain.CustomerStore // Customer Store value
}

// Used to create a new instance of Customer
func (c CustomerController) AddNewCustomer(newcustomer domain.Customer) {
	err := c.store.Create(newcustomer)
	if err != nil {
		fmt.Println("Error while trying to create a new Customer :", err)
		return
	}
}

// Used to update the details of the customer
func (c CustomerController) UpdateCustomer(oldcustomer domain.Customer, updatedcustomer domain.Customer) {
	err := c.store.Update(oldcustomer.ID, updatedcustomer)
	if err != nil {
		fmt.Println("Error while trying to update a Customer :", err)
		return
	}
}

// Used to reterive the details of the specifed customer's ID
func (c CustomerController) GetCustomerByID(ID string) {
	customer, err := c.store.GetByID(ID)
	if err != nil {
		fmt.Println("Error while trying to get the details of the Customer:", err)
		return
	}
	fmt.Println("Specific Customer details are:", customer)
}

// Used to delete the details of the customer with the specified ID
func (c CustomerController) DeleteCustomer(ID string) {
	if err := c.store.Delete(ID); err != nil {
		fmt.Println("Error while trying to delete the Customer:", err)
		return
	}
}

// Used to get all the data store under the Customer store
func (c CustomerController) GetAllCustomers() {
	data, err := c.store.GetAll()
	if err != nil {
		fmt.Println("Error while trying to get all the Customers:", err)
		return
	}
	for _, customer := range data {
		fmt.Println("Customer details are:", customer)
	}
}

func main() {
	controller := CustomerController{
		store: domain.NewMapStore(),
	}
	nc := domain.Customer{
		ID:    "1",
		Name:  "Hitesh",
		Email: "hitesh.balegar@betsol.com",
	}
	// Create a new customer
	controller.AddNewCustomer(nc)

	// Get all the customers
	controller.GetAllCustomers()

	// Update the customer
	nc.Name = "Hitesh Kumar"
	controller.UpdateCustomer(nc, nc)

	// Get the customer
	controller.GetCustomerByID("1")

	// Delete the customer
	controller.DeleteCustomer("1")

	// Adding a new customer
	nc2 := domain.Customer{
		ID:    "2",
		Name:  "Hitesh",
		Email: "hitesh.balegar@betsol.com",
	}
	// Create a new customer
	controller.AddNewCustomer(nc2)

	// Get the customer
	controller.GetCustomerByID("2")

	// Get all the customers
	controller.GetAllCustomers()

}
