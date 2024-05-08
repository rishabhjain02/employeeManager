package main

import (
	"fmt"

	"github.com/tokopedia/employeeManager/internal/delivery"
	"github.com/tokopedia/employeeManager/internal/model"
	"github.com/tokopedia/employeeManager/internal/repository"
)

func main() {
	store := &repository.EmployeeStore{
		Employees: make(map[int]model.Employee),
	}

	employee1 := model.Employee{
		ID:       1,
		Name:     "Employee1",
		Position: "Software Engineer",
		Salary:   80000,
	}

	employee2 := model.Employee{
		ID:       2,
		Name:     "Employee2",
		Position: "Product Manager",
		Salary:   90000,
	}

	employee3 := model.Employee{
		ID:       3,
		Name:     "Employee3",
		Position: "Data Scientist",
		Salary:   70000,
	}

	employee4 := model.Employee{
		ID:       4,
		Name:     "Employee4",
		Position: "Tester",
		Salary:   50000,
	}

	employee5 := model.Employee{
		ID:       5,
		Name:     "Employee5",
		Position: "HR",
		Salary:   75000,
	}

	// Creating some employees for testing
	store.CreateEmployee(employee1)
	store.CreateEmployee(employee2)
	store.CreateEmployee(employee3)
	store.CreateEmployee(employee4)
	store.CreateEmployee(employee5)

	// Get an employee by id
	employee, GetErr := store.GetEmployeeByID(2)
	if GetErr != nil {
		fmt.Println(GetErr)
	}

	fmt.Println("Details of the employee having id as 2")
	fmt.Printf("Name: %s | Position: %s | Salary: %0.2f", employee.Name, employee.Position, employee.Salary)

	// Update employee
	newEmployee := model.Employee{
		ID:       5,
		Name:     "Employee5",
		Position: "HR",
		Salary:   85000,
	}
	updateErr := store.UpdateEmployee(newEmployee)
	if updateErr != nil {
		fmt.Println(updateErr)
	}

	// Delete employee
	errDelete := store.DeleteEmployee(4)
	if errDelete != nil {
		fmt.Println(errDelete)
	}

	// Intializing and calling delivery layer
	delivery := delivery.NewHTTPDelivery(store)
	delivery.Serve()
}
