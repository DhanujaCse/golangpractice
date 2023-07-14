package main

import (
	"fmt"
	"golangprg/coding/coding1"
)

func main() {
	s := coding1.Student{}
	fmt.Println("hello")
	fmt.Println("enter name of student")

	fmt.Scanln(&s.Name)
	fmt.Println("enter age of student")
	fmt.Scanln(&s.Age)
	fmt.Println("enter the value of subject1")
	fmt.Scanln(&s.Grades1[0])
	fmt.Println("enter the value of subject2")
	fmt.Scanln(&s.Grades1[1])
	fmt.Println("enter the value of subject2")
	fmt.Scanln(&s.Grades1[2])
	student1 := s
	avg := coding1.Calculateavg(student1)
	fmt.Printf("average of student %v is %v\n", s.Name, avg)

	student2 := coding1.Student{"dhanuja", 24, [...]int{2, 4, 4}}
	fmt.Println(coding1.Calculateavg(student2))

	//program 2

	r := coding1.Rectangle{Width: 3, Length: 4}
	coding1.PrintAreaAndPerimeter(r)
	s1 := coding1.Square{A: 4}
	coding1.PrintAreaAndPerimeter(s1)

}
