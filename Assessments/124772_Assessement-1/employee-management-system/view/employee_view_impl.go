package view

import (
	"fmt"

	"ems/model"
)

type EmployeeViewImpl struct {
}

func NewEmployeeView() EmployeeView {
	return &EmployeeViewImpl{}
}

func (v *EmployeeViewImpl) ShowMenu() int {

	fmt.Println("============== Employee Management System ===============")
	fmt.Println("1. Add Employee")
	fmt.Println("2. View All Employees")
	fmt.Println("3. Exit")

	var choice int

	fmt.Print("\nEnter Your choice: ")
	fmt.Scan(&choice)

	return choice
}

func (v *EmployeeViewImpl) ReadEmployee() model.Employee {

	var employee model.Employee

	fmt.Println("------------- Read Employee -----------")

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&employee.ID)

	fmt.Print("Enter Employee Name: ")
	fmt.Scan(&employee.Name)

	fmt.Print("Enter Employee Age: ")
	fmt.Scan(&employee.Age)

	fmt.Print("Enter Employee Email: ")
	fmt.Scan(&employee.Email)

	return employee
}