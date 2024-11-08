package model

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Applicant is a struct that contains information of an applicant
type Applicant struct {
	FirstName  string
	LastName   string
	GPA        float64
	Department []string
}

// Applicants is a struct that contains a list of applicants
type Applicants struct {
	Applicants []Applicant
}

const (
	MaxColumn = 6
)

// InsertFile is a method that reads a file and insert the data into the Applicants struct
func (a *Applicants) InsertFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer file.Close()

	var applicants []Applicant

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		applicant := strings.Fields(line)

		firstName := applicant[0]
		lastName := applicant[1]
		GPA, _ := strconv.ParseFloat(applicant[2], 64)
		department := applicant[3:]

		if len(applicant) >= MaxColumn {
			applicant := Applicant{
				FirstName:  firstName,
				LastName:   lastName,
				GPA:        GPA,
				Department: department,
			}

			applicants = append(applicants, applicant)
		}
	}

	a.Applicants = applicants
}

func (a *Applicants) RemoveApplicant(index int) {
	a.Applicants = append(a.Applicants[:index], a.Applicants[index+1:]...)
}

// DepartmentClassification is a method that classify applicant by department
func (a *Applicants) DepartmentClassification(departmentName string, top int) Applicants {
	var applicants []Applicant
	for _, applicant := range a.Applicants {
		if applicant.Department[top] == departmentName {
			applicants = append(applicants, applicant)
		}
	}

	return Applicants{Applicants: applicants}
}

// SortByGPA is a method that sort applicant by GPA
func (a *Applicants) SortByGPA() {
	sort.Slice(a.Applicants, func(i, j int) bool {
		if a.Applicants[i].GPA == a.Applicants[j].GPA {
			if a.Applicants[i].FirstName == a.Applicants[j].FirstName {
				return a.Applicants[i].LastName > a.Applicants[j].LastName
			}
			return a.Applicants[i].FirstName > a.Applicants[j].FirstName
		}
		return a.Applicants[i].GPA > a.Applicants[j].GPA
	})
}

// GetAllDepartment is a method that get all department in the list of applicants
func (a *Applicants) GetAllDepartment() []string {
	var departments []string
	departmentDOT := strings.Join(departments, ",")

	for _, applicant := range a.Applicants {
		for _, department := range applicant.Department {
			if !strings.Contains(departmentDOT, department) {
				departmentDOT += department + ","
				departments = append(departments, department)
			}
		}
	}

	return departments
}
