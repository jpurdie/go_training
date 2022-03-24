package main

import (
	"fmt"
	"math/rand"
	"time"
)

type scenario struct {
	Num     int
	Persons []person
}

type person struct {
	CrimeStatus  string
	Age          int
	Gender       string
	FamilyStatus string
}

/*
1. The witness and the accomplice of the murdered are not the same sex.
2. The oldest person of the family and the witness were a man and a woman
3. The youngest person and the victim are of the same sex
4. The murderer's accomplice was older than the poor victim
5. The father was the oldest member of the family
6. The murderer was not the youngest in the family
*/

func ruleSix(sc scenario) bool {
	youngestAge := 0
	murdererAge := 100

	for _, person := range sc.Persons {
		if person.Age < youngestAge {
			youngestAge = person.Age
		}
		if person.CrimeStatus == "murderer" {
			murdererAge = person.Age
		}
	}
	if murdererAge == youngestAge {
		return false
	}
	return true
}

func ruleFive(sc scenario) bool {
	oldestAge := 0
	fatherAge := 0

	for _, person := range sc.Persons {
		if person.Age > oldestAge {
			oldestAge = person.Age
		}
		if person.FamilyStatus == "father" {
			fatherAge = person.Age
		}
	}
	if fatherAge == oldestAge {
		return true
	}
	return false
}

func ruleFour(sc scenario) bool {
	accompliceAge := 0
	victimAge := 0

	for _, person := range sc.Persons {
		if person.CrimeStatus == "accomplice" {
			accompliceAge = person.Age
		}
		if person.CrimeStatus == "victim" {
			victimAge = person.Age
		}
	}
	if accompliceAge > victimAge {
		return true
	}
	return false
}

func ruleThree(sc scenario) bool {
	youngestGender := ""
	youngestAge := 1000
	victimGender := ""

	for _, person := range sc.Persons {
		if person.Age < youngestAge {
			youngestAge = person.Age
			youngestGender = person.Gender
		}
		if person.CrimeStatus == "victim" {
			victimGender = person.Gender
		}
	}
	if youngestGender != victimGender {
		return false
	}
	return true
}

func ruleTwoPass(sc scenario) bool {
	oldestGender := ""
	oldestAge := 0
	witnessGender := ""

	for _, person := range sc.Persons {
		if person.Age > oldestAge {
			oldestAge = person.Age
			oldestGender = person.Gender
		}
		if person.CrimeStatus == "witness" {
			witnessGender = person.Gender
		}
	}
	if oldestGender == witnessGender {
		return false
	}
	return true
}

func ruleOnePass(sc scenario) bool {
	witnessGender := ""
	accompliceGender := ""
	for _, person := range sc.Persons {
		if person.CrimeStatus == "witness" {
			witnessGender = person.Gender
		}
		if person.CrimeStatus == "accomplice" {
			accompliceGender = person.Gender
		}
	}
	if witnessGender != accompliceGender {
		return true
	}
	return false
}

func buildScenarios(num int) []scenario {

	scenarios := make([]scenario, 0)

	for i := 0; i < num; i++ {
		s := scenario{}
		s.Num = i
		crimeStatuses := []string{"accomplice", "witness", "murderer", "victim"}

		minAgeChild := 1
		maxAgeChild := 20

		time.Sleep(3)
		rand.Seed(time.Now().UnixNano())

		statusInt := rand.Intn(len(crimeStatuses))
		father := person{FamilyStatus: "father", Age: 51, Gender: "M", CrimeStatus: crimeStatuses[statusInt]}
		crimeStatuses = remove(crimeStatuses, statusInt)
		s.Persons = append(s.Persons, father)

		time.Sleep(3)
		rand.Seed(time.Now().UnixNano())

		statusInt = rand.Intn(len(crimeStatuses))
		mother := person{FamilyStatus: "mother", Age: 50, Gender: "F", CrimeStatus: crimeStatuses[statusInt]}
		crimeStatuses = remove(crimeStatuses, statusInt)
		s.Persons = append(s.Persons, mother)

		time.Sleep(3)
		rand.Seed(time.Now().UnixNano())

		statusInt = rand.Intn(len(crimeStatuses))
		son := person{FamilyStatus: "son", Age: rand.Intn(maxAgeChild-minAgeChild) + minAgeChild, Gender: "M", CrimeStatus: crimeStatuses[statusInt]}
		crimeStatuses = remove(crimeStatuses, statusInt)
		s.Persons = append(s.Persons, son)

		time.Sleep(3)
		rand.Seed(time.Now().UnixNano())

		statusInt = rand.Intn(len(crimeStatuses))
		daughter := person{FamilyStatus: "daughter", Age: rand.Intn(maxAgeChild-minAgeChild) + minAgeChild, Gender: "F", CrimeStatus: crimeStatuses[statusInt]}
		crimeStatuses = remove(crimeStatuses, statusInt)
		s.Persons = append(s.Persons, daughter)
		scenarios = append(scenarios, s)
	}
	return scenarios
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func removeScenario(slice []scenario, s int) []scenario {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	rulePass := false
	newScenarios := make([]scenario, 0)
	scenarios := buildScenarios(10)

	fmt.Println("Rule 1) The witness and the accomplice of the murdered are not the same sex.")

	for _, sc := range scenarios {
		rulePass = ruleOnePass(sc)
		if rulePass {
			newScenarios = append(newScenarios, sc)
		}
	}

	fmt.Println("Rule 2) The oldest person of the family and the witness were a man and a woman.")

	scenarios = newScenarios
	newScenarios = make([]scenario, 0)
	for _, sc := range scenarios {
		rulePass = ruleTwoPass(sc)
		if rulePass {
			newScenarios = append(newScenarios, sc)
		}
	}

	fmt.Println("Rule 3) The youngest person and the victim are of the same sex.")

	scenarios = newScenarios
	newScenarios = make([]scenario, 0)
	for _, sc := range scenarios {
		rulePass = ruleThree(sc)
		if rulePass {
			newScenarios = append(newScenarios, sc)
		}
	}

	fmt.Println("Rule 4) The murderer's accomplice was older than the poor victim.")

	scenarios = newScenarios
	newScenarios = make([]scenario, 0)
	for _, sc := range scenarios {
		rulePass = ruleFour(sc)
		if rulePass {
			newScenarios = append(newScenarios, sc)
		}
	}

	fmt.Println("Rule 5) The father was the oldest member of the family.")

	scenarios = newScenarios
	newScenarios = make([]scenario, 0)
	for _, sc := range scenarios {
		rulePass = ruleFive(sc)
		if rulePass {
			newScenarios = append(newScenarios, sc)
		}
	}

	fmt.Println("Rule 6) The murderer was not the youngest in the family")

	scenarios = newScenarios
	newScenarios = make([]scenario, 0)
	for _, sc := range scenarios {
		rulePass = ruleSix(sc)
		if rulePass {
			newScenarios = append(newScenarios, sc)
		}
	}

	fmt.Println("-- Passed --")

	for _, sc := range scenarios {
		fmt.Println(sc)
	}

}
