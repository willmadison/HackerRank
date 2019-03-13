package main

import "testing"

func TestGetSecondLowestScoringStudent(t *testing.T) {
	students := []Student{
		{
			Name:  "Will",
			Grade: 100.0,
		},
		{
			Name:  "Lauren",
			Grade: 85.0,
		},
		{
			Name:  "Olivia",
			Grade: 70.0,
		},
		{
			Name:  "Hannah",
			Grade: 55.0,
		},
		{
			Name:  "Emani",
			Grade: 55.0,
		},
	}

	secondLowestScorers := FindSecondLowestScorers(students)

	if len(secondLowestScorers) > 1 {
		t.Fatal("expected only 1 second lowest scorer, got:", len(secondLowestScorers))
	}

	if secondLowestScorers[0].Name != "Olivia" {
		t.Error("expected Olivia to be the second lowest scorer, got:", secondLowestScorers[0])
	}
}
