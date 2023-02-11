package main

type User struct {
	FirstName string
}

// recevier sample we have attached this function to the struct
func (m *User) printFirstName() string {
	return m.FirstName
}

// func main() {
// 	var myVar User
// 	myVar.FirstName = "Santosh"

// 	log.Println("My first Name = ", myVar.printFirstName())
// }
