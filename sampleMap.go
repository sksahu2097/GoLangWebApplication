package main

type Student struct {
	name   string
	rollno int
}

func (std *Student) printstudent() string {
	return std.name
}

// func main() {
// 	myMap := make(map[string]Student)
// 	myMap["std1"] = Student{
// 		name:   "santosh",
// 		rollno: 35,
// 	}
// 	log.Println("Value = ", myMap["std1"])
// 	myMap["std2"] = Student{
// 		name:   "Amrapali",
// 		rollno: 35,
// 	}
// 	log.Println("name = ", myMap)

// }
