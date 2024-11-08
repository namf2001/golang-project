package main

import (
	"fmt"

	"golang-project/project/UniversityAdmissionProcedure/internal/model"
)

func main() {
	var applicants model.Applicants
	applicants.InsertFile("project/UniversityAdmissionProcedure/file_test/applicant_list.txt")
	applicants.SortByGPA()

	departments := applicants.GetAllDepartment()

	admittedApplicants := make(map[string][]model.Applicant)

	for _, department := range departments {
		fmt.Println("-----------------DEPARTMENT: ", department, "-----------------")
		slot := 2
		for i, applicant := range applicants.Applicants {
			if applicant.Department[0] == department && slot > 0 {
				fmt.Println(i+1, applicant.FirstName, applicant.LastName, applicant.GPA)
				slot--
			} else {
				continue
			}
		}
	}
}
