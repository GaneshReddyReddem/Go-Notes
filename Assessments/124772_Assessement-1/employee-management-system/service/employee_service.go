package service
import "ems/model"

type EmployeeService interface{
	AddEmployee(employee model.Employee) error
	GetEmployees() ([]model.Employee, error)
}