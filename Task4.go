package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	rollnumber int
	name       string
	address    string
	Subjects   []string
}

func NewStudent(rollno int, name string, address string, subjects []string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	s.Subjects = subjects
	return s
}

func getstring(rollno int, name string, address string, subjects []string) string {
	strrollno := strconv.Itoa(rollno)
	strsubjects := strings.Join(subjects, " ")
	block := name + strrollno + address + strsubjects
	fmt.Println(block)
	return block

}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string, subjects []string) *Student {
	st := NewStudent(rollno, name, address, subjects)
	ls.list = append(ls.list, st)
	return st
}

func (ls *StudentList) Print() {
	for i := range ls.list {
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Printf("Student Rollno: %d\n", ls.list[i].rollnumber)
		fmt.Printf("Student Name: %s\n", ls.list[i].name)
		fmt.Printf("Student Address: %s\n\n", ls.list[i].address)
		fmt.Printf("Student Subjects: %s\n\n", ls.list[i].Subjects)

	}
}

func main() {
	student := new(StudentList)
	subjects1 := []string{"Maths", "English", "Physics"}
	student.CreateStudent(1746, "Muhammad Ahmad", "Sahiwal", subjects1)
	block1 := getstring(1746, "Muhammad Ahmad", "Sahiwal", subjects1)
	sum := sha256.Sum256([]byte(block1))
	fmt.Printf("%x\n", sum) //hexadecimal

	subjects2 := []string{"Biology", "Islamiyat", "Computer Science"}
	student.CreateStudent(2192, "Zakkaulah", "Abotabad", subjects2)

	fmt.Println("\n")
	block2 := getstring(2192, "Zakkaulah", "Abotabad", subjects2)
	sum1 := sha256.Sum256([]byte(block2))
	fmt.Printf("%x\n", sum1) //hexadecimal
	student.Print()

}