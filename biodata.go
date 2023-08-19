package main

import (
	"fmt"
	"os"
)

type Person struct {
	name string
	addres  string
	job string
	reason string
}

func main() {

	 biodata := []Person{
		{name : "Boy", addres: "Jakarta Timur", job: "Programmer", reason: "Karena tertarik dengan Golang"},
		{name : "Sebastian", addres: "Kelapa Gading", job: "Quality Tester", reason: "Karena tertarik dengan Golang"},
		{name : "Noventi", addres: "Papua", job: "Front End Programmer", reason: "Karena tertarik dengan Golang"},
		{name : "Novita", addres: "Jakarta Timur", job: "Programmer", reason: "Karena tertarik dengan Golang"},
		
	}

	var biodata1 = os.Args[1]

	for i, v := range biodata {
		if biodata1 == v.name {
		fmt.Println("id:", i)
			fmt.Println("name :", v.name)
			fmt.Println("addres :", v.addres)
			fmt.Println("job :", v.job)
			fmt.Println("reason :", v.reason)
		}else{
			fmt.Println("data tidak ada")
		}
	} 
}	

