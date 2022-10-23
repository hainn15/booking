package main

//complite funtion in package so we need import package.
import "fmt" //fmt stand for Format package

//pkg.go.dev/std to find funtion go with package

//main funtiion is the entrypoint of a Go program
//A program can only have 1 main funtion because you can only have 1 entrypoint
func main() {

	//
	var (
		name    string = "Tran A"
		address string = "Ha Noi"
		number  int    = 012
	)

	fmt.Println(name) //use funtion of package
	fmt.Println(address)
	fmt.Println(number)

	//type inference
	//khac kieu du lieu
	var name1, address1, age = "Tran Nam", "Hai Phong", 14
	fmt.Println(name1)
	fmt.Println(address1)
	fmt.Println(age)

	var conferenceName = "Conferance Name"
	fmt.Printf(conferenceName)

}

//go run, go build, go install
