package main

import "fmt"

type student struct {
	firstname string
	lastname  string
	score     float64
}

func main() {
	studentdata := getdata()
	displaydata(studentdata)
}

func getdata() (studentdata student) {
	var studentfirstname, studentlastname string
	var studentscore float64
	fmt.Print("Input the student's first name : ")
	fmt.Scan(&studentfirstname)
	fmt.Print("Input the student's last name : ")
	fmt.Scan(&studentlastname)
	fmt.Print("Input the student's score : ")
	fmt.Scan(&studentscore)

	studentdata = student{firstname: studentfirstname,
		lastname: studentlastname,
		score:    studentscore}
	return

}

func displaydata(studentdata student) {
	fmt.Print("\nThe student's first name is ", studentdata.firstname)
	fmt.Print("\nThe student's last name is ", studentdata.lastname)
	fmt.Print("\nThe student's score is ", studentdata.score)
}
