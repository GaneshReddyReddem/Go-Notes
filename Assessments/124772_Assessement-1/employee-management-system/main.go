package main

import(
	"ems/dao"
	"ems/service"
	"ems/controller"
	"ems/view"
)

func main(){

	employeeDao := dao.NewEmployeeDAO()
	employeeService := service.NewEmployeeService(employeeDao)
	employeeView := view.NewEmployeeView()
	employeeController := controller.NewEmployeeController(employeeView,employeeService)
	employeeController.Start()
}