package main

import "fmt"

type Person struct {
	name string
	addres  string
	job string
	reason string
}

func main() {

	 biodata := []Person{
		{name : "Boy Valentyas", addres: "Jakarta Timur", job: "Programmer", reason: "Karena tertarik dengan Golang"},
		
	}

	for i, v := range biodata {
		fmt.Println("id:", i)
			fmt.Println("name :", v.name)
			fmt.Println("addres :", v.addres)
			fmt.Println("job :", v.job)
			fmt.Println("reason :", v.reason)
	}

}	

