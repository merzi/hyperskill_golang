// write your code here
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var m, n int
	var inputBuffer string
	uni := university{}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	m, err := strconv.Atoi(input)
	if err != nil {
		return
	}

	scanner.Scan()
	input = scanner.Text()
	n, err = strconv.Atoi(input)
	if err != nil {
		return
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		inputBuffer = scanner.Text()

		parts := strings.Fields(inputBuffer)
		fmt.Println(parts)
		if len(parts) < 3 {
			fmt.Printf("only %d parts found!", len(parts))
			return
		}

		val, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			return
		}
		fmt.Println(parts)
		uni.addApplicant(parts[0], parts[1], val)
	}

	fmt.Println("Successful applicants:")
	for _, a := range uni.getRankedApplicants(m) {
		fmt.Println(a.getFullName())
	}

}

type university struct {
	department []department
	applicants []applicant
	students   []student
}

type department struct {
	applicants []applicant
	students   []student
}

type student struct {
	name      string
	firstName string
}

type applicant struct {
	name      string
	firstname string
	gpa       float64
}

func (d department) addStudent(firstname string, name string) {
	d.students = append(d.students, student{name: name, firstName: firstname})
}

func (d department) addApplicant(firstname string, name string, gpa float64) {
	d.applicants = append(d.applicants, applicant{name: name, firstname: firstname, gpa: gpa})
}

func (u university) getRankedApplicants(count int) []applicant {
	var a []applicant
	var buffer []applicant
	position := -1

	for _, currentApplicant := range u.applicants {
		buffer = []applicant{}
		if len(a) == 0 {
			a = append(a, currentApplicant)
			continue
		}

		if a[0].gpa > currentApplicant.gpa {
			buffer = append(buffer, currentApplicant)
			a = append(buffer, a...)
			continue
		} else if a[len(a)-1].gpa < currentApplicant.gpa {
			a = append(a, currentApplicant)
			continue
		}

		for key, v := range a {
			if key == len(u.applicants)-1 {
				break
			}

			if v.gpa < currentApplicant.gpa && u.applicants[key+1].gpa > currentApplicant.gpa {
				position = key + 1
				break
			}
		}

		if position == -1 {
			return nil
		}

		buffer = append(append(buffer, a[:position-1]...), currentApplicant)
		a = append(buffer, a[position:]...)
	}

	return a[:count-1]
}

func (a applicant) getFullName() string {
	return fmt.Sprintf("%s %s", a.firstname, a.name)
}

func universityAdmission() {
	var number1, number2, number3 int

	_, err := fmt.Scanf("%d", &number1)
	if err != nil {
		return
	}
	_, err = fmt.Scanf("%d", &number2)
	if err != nil {
		return
	}
	_, err = fmt.Scanf("%d", &number3)
	if err != nil {
		return
	}

	meanScore := (float64(number1) + float64(number2) + float64(number3)) / 3
	fmt.Println(fmt.Sprintf("%.2f", meanScore))
	if meanScore >= 60.0 {
		fmt.Println("Congratulations, you are accepted!")
	} else {
		fmt.Println("We regret to inform you that we will not be able to offer you admission.")
	}
}

func readFile(filePath string) []string {
	fileOpener, err := os.Open(filePath)
	if err != nil {
		return nil
	}

	var output []string
	reader := bufio.NewReader(fileOpener)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		output = append(output, string(line))
	}

	return output
}
