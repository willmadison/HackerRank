package main

import (
	"fmt"
	"sort"
)

func main() {
	var numStudents int

	fmt.Scanf("%v", &numStudents)

	var students []Student

	for i := 0; i < numStudents; i++ {
		var name string
		var grade float32

		fmt.Scanf("%v", &name)
		fmt.Scanf("%v", &grade)

		student := Student{
			Name:  name,
			Grade: grade,
		}

		students = append(students, student)
	}

	secondLowestScorers := FindSecondLowestScorers(students)

	for _, s := range secondLowestScorers {
		fmt.Println(s.Name)
	}
}

type Student struct {
	Name  string
	Grade float32
}

func FindSecondLowestScorers(students []Student) []Student {
	studentsByGrade := map[float32][]Student{}
	uniqueGrades := map[float32]struct{}{}

	for _, student := range students {
		uniqueGrades[student.Grade] = struct{}{}
		studentsByGrade[student.Grade] = append(studentsByGrade[student.Grade], student)
	}

	var grades []float32

	for grade := range uniqueGrades {
		grades = append(grades, grade)
	}

	sort.Slice(grades, func(i, j int) bool {
		return grades[i] < grades[j]
	})

	secondLowestScore := grades[1]

	s := studentsByGrade[secondLowestScore]
	sort.Slice(s, func(i, j int) bool {
		return s[i].Name < s[j].Name
	})

	return s
}
