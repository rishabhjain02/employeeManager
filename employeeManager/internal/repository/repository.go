package repository

import "github.com/tokopedia/employeeManager/internal/model"

type Repository interface {
	CreateEmployee(employee model.Employee)
	GetEmployeeByID(id int) (model.Employee, error)
	UpdateEmployee(employee model.Employee) error
	DeleteEmployee(id int) error
	ListEmployees(page, pageSize int) ([]model.Employee, error)
}
