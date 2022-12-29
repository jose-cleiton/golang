package main

import (

	"fmt"

	"github.com/jose-cleiton/data"
	


)


func main() {
	zoo:=data.Data

	isManager := data.IsManager("burlId", zoo)
	
	allLionsOlderThan12 := data.GetAnimalsOlderThan(zoo, "lions", 12)

	allTigersOlderThan5 := data.GetAnimalsOlderThan(zoo, "tigers", 5)

	speciesById := data.GetSpeciesByIds(zoo, "0938aa23-f153-4937-9f88-4858b24d6bce")

	employee := data.GetEmployeeByName(zoo.Employees, "Sharonda", "Spry")


	fmt.Printf("\033[2J")
	fmt.Println("-->Resultado da fucção: GetAnimalsOlderThan<--")
	
	fmt.Println("Leão Maior que 12",allLionsOlderThan12)
	fmt.Println("Tigre Maior que 5",allTigersOlderThan5)
	

	fmt.Println("-->Resultado da fucção: GetSpeciesByIds<--")

	
	fmt.Println("Nome: ", speciesById[0].Name)
	fmt.Println("Localização: ", speciesById[0].Location)
	fmt.Println("Disponibilidade: ", speciesById[0].Availability)
	fmt.Println("Residentes: ", speciesById[0].Residents)

	fmt.Println("-->Resultado da fucção: GetEmployeeByName<--")

  
  
	fmt.Printf("Nome: %s %s\n", employee.FirstName, employee.LastName)
	fmt.Printf("Responsável por: %s\n", employee.ResponsibleFor)
	fmt.Printf("Gerente: %s\n", employee.Managers)

  fmt.Println("-->Resultado da fucção: IsManager<--")

	fmt.Println("burlId e gerente:", isManager)

	fmt.Println("-->Resultado da fucção: GetRelatedEmployees<--")
	equip("stephanieId", zoo)

}


	func equip(manager string, zoo data.Zoo) {
		ger, err:= data.GetRelatedEmployees(manager, zoo)

		if err == nil {
			fmt.Println("Lista de gerenciados por", manager,":\n",  ger)
		}else {
	
			fmt.Println(err)
		}
	

	}

	


	


    