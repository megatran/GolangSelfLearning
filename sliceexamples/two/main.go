package main

import "fmt"

func main() {
	records := make([][]string, 0)
	//student 1
	student1 := make([]string, 4)
	student1[0] = "Nhan"
	student1[1] = "Tran"
	student1[2] = "100.00"
	student1[3] = "89.00"
	//store the record
	records = append(records, student1)

	//student2
	student2 := make([]string, 4)
	student2[0] = "Jacob"
	student2[1] = "Swift"
	student2[2] = "92.00"
	student2[3] = "96.00"
	//store the record
	records = append(records, student2)

	//test creating a nil slice (uninitialized)
	var student3 []string
	records = append(records, student3)
	//print
	fmt.Println(records)
	fmt.Println("is student 3 nil? ", student3 == nil)
}
