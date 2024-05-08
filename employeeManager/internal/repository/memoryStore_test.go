package repository

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tokopedia/employeeManager/internal/model"
)

func Test_memoryStore_CreateEmployee(t *testing.T) {
	store := &EmployeeStore{
		Employees: make(map[int]model.Employee),
	}

	employee := model.Employee{
		ID:       1,
		Name:     "Test Employee",
		Position: "Test Position",
		Salary:   1000,
	}

	type args struct {
		employee model.Employee
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Create employee executes properly",
			args: args{
				employee: employee,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store.CreateEmployee(tt.args.employee)
			assert.Equal(t, tt.args.employee, store.Employees[tt.args.employee.ID], "Employee should be created and stored")
		})
	}
}

func Test_memoryStore_GetEmployeeByID(t *testing.T) {
	store := &EmployeeStore{
		Employees: map[int]model.Employee{
			1: {ID: 1, Name: "Test Employee1"},
			2: {ID: 2, Name: "Test Employee2"},
		},
	}

	type args struct {
		id int
	}

	tests := []struct {
		name       string
		args       args
		wantResult model.Employee
		wantErr    bool
	}{
		{
			name: "Get employee by id executes properly",
			args: args{
				id: 1,
			},
			wantResult: model.Employee{ID: 1, Name: "Test Employee1"},
			wantErr:    false,
		},
		{
			name: "Get employee by id gives error",
			args: args{
				id: 3,
			},
			wantResult: model.Employee{},
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := store.GetEmployeeByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get employee by id got error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(data, tt.wantResult) {
				t.Errorf("Get employee by id got = %v, want %v", data, tt.wantResult)
			}
		})
	}
}

func Test_memoryStore_UpdateEmployee(t *testing.T) {
	store := &EmployeeStore{
		Employees: map[int]model.Employee{
			1: {ID: 1, Name: "Test Employee1"},
		},
	}

	type args struct {
		employee model.Employee
	}

	updatedEmployee := model.Employee{
		ID:   1,
		Name: "Test Employee100",
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Update employee executes properly",
			args: args{
				employee: updatedEmployee,
			},
			wantErr: false,
		},
		{
			name: "Update employee gives error",
			args: args{
				employee: model.Employee{ID: 3, Name: "Test Employee"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := store.UpdateEmployee(tt.args.employee)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update employee got error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_memoryStore_DeleteEmployee(t *testing.T) {
	store := &EmployeeStore{
		Employees: map[int]model.Employee{
			1: {ID: 1, Name: "Test Employee1"},
		},
	}

	type args struct {
		id int
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Delete employee executes properly",
			args: args{
				id: 1,
			},
			wantErr: false,
		},
		{
			name: "Delete employee gives error",
			args: args{
				id: 3,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := store.DeleteEmployee(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete employee got error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_memoryStore_ListEmployees(t *testing.T) {
	store := &EmployeeStore{
		Employees: map[int]model.Employee{
			1: {ID: 1, Name: "Test Employee1"},
			2: {ID: 2, Name: "Test Employee2"},
		},
	}

	type args struct {
		page     int
		pageSize int
	}

	tests := []struct {
		name       string
		args       args
		wantResult []model.Employee
		wantErr    bool
	}{
		{
			name: "List employee executes properly",
			args: args{
				page:     1,
				pageSize: 2,
			},
			wantResult: []model.Employee{
				{ID: 1, Name: "Test Employee1"},
				{ID: 2, Name: "Test Employee2"},
			},
			wantErr: false,
		},
		{
			name: "List employee gives error",
			args: args{
				page:     3,
				pageSize: 1,
			},
			wantResult: nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			employees, err := store.ListEmployees(tt.args.page, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete employee got error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(employees, tt.wantResult) {
				t.Errorf("Get employee by id got = %v, want %v", employees, tt.wantResult)
			}
		})
	}
}
