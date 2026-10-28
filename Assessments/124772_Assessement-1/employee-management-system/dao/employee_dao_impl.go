package dao

import "ems/model"

// Dependency Injection
type EmployeeDAOImpl struct {
	employees []model.Employee
}

// Create function to return address
func NewEmployeeDAO() EmployeeDAO {
	return &EmployeeDAOImpl{
		employees: make([]model.Employee, 0),
	}
}

// Implement Interface methods

func (d *EmployeeDAOImpl) Save(employee model.Employee) error {

	d.employees = append(d.employees, employee)

	return nil
}

func (d *EmployeeDAOImpl) FindAll() ([]model.Employee, error) {

	return d.employees, nil
}