package  dao

import "ems/model"

type EmployeeDAO interface{
	Save(employee model.Employee) error
	FindAll() ([]model.Employee, error)
}