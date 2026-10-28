package service

import (
	"ems/dao"
	"ems/model"
	"ems/utility"
)

// Dependency Injection
type EmployeeServiceImpl struct {
	dao dao.EmployeeDAO
}

// Function to return address
func NewEmployeeService(dao dao.EmployeeDAO) EmployeeService {
	return &EmployeeServiceImpl{
		dao: dao,
	}
}

// CRUD

func (s *EmployeeServiceImpl) AddEmployee(employee model.Employee) error {

	// Employee validation
	err := utility.ValidateEmployee(employee)
	if err != nil {
		return err
	}

	// Email validation
	err = utility.ValidateEmail(employee.Email)
	if err != nil {
		return err
	}

	// Save employee
	err = s.dao.Save(employee)
	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeServiceImpl) GetEmployees() ([]model.Employee, error) {
	return s.dao.FindAll()
}