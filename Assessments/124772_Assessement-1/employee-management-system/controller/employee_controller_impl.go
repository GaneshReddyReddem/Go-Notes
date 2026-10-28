package controller

import (
	"fmt"

	"ems/service"
	"ems/view"
)

type EmployeeControllerImpl struct {
	view    view.EmployeeView
	service service.EmployeeService
}

func NewEmployeeController(view view.EmployeeView, service service.EmployeeService) EmployeeController {
	return &EmployeeControllerImpl{
		view:    view,
		service: service,
	}
}

func (c *EmployeeControllerImpl) Start() {

	for {
		ch := c.view.ShowMenu()

		switch ch {
		case 1:
			c.Process(1)

		case 2:
			c.Process(2)

		case 3:
			fmt.Println("Thank you..")
			return

		default:
			fmt.Println("Invalid choice..")
		}
	}
}

func (c *EmployeeControllerImpl) Process(choice int) {

	if choice == 1 {

		employee := c.view.ReadEmployee()

		err := c.service.AddEmployee(employee)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Employee added successfully..")
	}

	if choice == 2 {

		emps, err := c.service.GetEmployees()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if len(emps) == 0 {
			fmt.Println("No employees found.")
			return
		}

		fmt.Println("============== Employees ==============")

		for _, emp := range emps {
			fmt.Println(emp)
		}
	}
}