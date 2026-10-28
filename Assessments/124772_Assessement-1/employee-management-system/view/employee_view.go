package view
import "ems/model"

type EmployeeView interface{
	ShowMenu() int
	ReadEmployee() model.Employee
}