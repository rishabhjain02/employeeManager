package repository

import (
	"fmt"
	"sync"

	"github.com/tokopedia/employeeManager/internal/model"
)

type EmployeeStore struct {
	sync.RWMutex
	Employees map[int]model.Employee
}

func (store *EmployeeStore) CreateEmployee(employee model.Employee) {
	store.Lock()
	defer store.Unlock()

	store.Employees[employee.ID] = employee
}

func (store *EmployeeStore) GetEmployeeByID(id int) (model.Employee, error) {
	store.RLock()
	defer store.RUnlock()

	employee, ok := store.Employees[id]

	if !ok {
		return model.Employee{}, fmt.Errorf("employee with id %d not found", id)
	}

	return employee, nil
}

func (store *EmployeeStore) UpdateEmployee(employee model.Employee) error {
	store.Lock()
	defer store.Unlock()

	if _, ok := store.Employees[employee.ID]; !ok {
		return fmt.Errorf("employee with id %d not found", employee.ID)
	}

	store.Employees[employee.ID] = employee
	return nil
}

func (store *EmployeeStore) DeleteEmployee(id int) error {
	store.Lock()
	defer store.Unlock()

	if _, ok := store.Employees[id]; !ok {
		return fmt.Errorf("employee with id %d not found", id)
	}

	delete(store.Employees, id)
	return nil
}

func (store *EmployeeStore) ListEmployees(page, pageSize int) ([]model.Employee, error) {
	store.RLock()
	defer store.RUnlock()

	start := (page - 1) * pageSize
	end := page * pageSize

	if start < 0 || start >= len(store.Employees) || end <= 0 || end > len(store.Employees) {
		return nil, fmt.Errorf("invalid pagination parameters")
	}

	employees := make([]model.Employee, 0, pageSize)
	for _, employee := range store.Employees {
		employees = append(employees, employee)
	}

	return employees[start:end], nil
}
