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

	// Create a copy of the Applicants
	applicantCopy := make([]model.Applicant, len(applicants.Applicants))
	copy(applicantCopy, applicants.Applicants)

	admittedApplicants := make(map[string][]model.Applicant)

	for i := 0; i < 3; i++ {
		for _, department := range departments {
			slot := 5
			var remainingApplicants []model.Applicant
			for _, applicant := range applicantCopy {
				if slot > 0 && i < len(applicant.Department) && applicant.Department[i] == department {
					admittedApplicants[department] = append(admittedApplicants[department], applicant)
					slot--
				} else {
					remainingApplicants = append(remainingApplicants, applicant)
				}
			}
			applicantCopy = remainingApplicants
		}
	}

	for department, applicants := range admittedApplicants {
		fmt.Println("-------------------------Department:", department, "-------------------------")
		for _, applicant := range applicants {
			fmt.Printf("First Name: %s, Last Name: %s, GPA: %.2f\n", applicant.FirstName, applicant.LastName, applicant.GPA)
		}
	}
}
