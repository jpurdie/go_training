package main

import (
	"testing"
)

func TestRuleOnePass(t *testing.T) {
	father := person{Gender: "M", CrimeStatus: "witness"}
	mother := person{Gender: "F", CrimeStatus: "accomplice"}

	s := scenario{}
	s.Persons = append(s.Persons, father, mother)

	got := ruleOnePass(s)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestRuleOneFail(t *testing.T) {
	father := person{Gender: "F", CrimeStatus: "witness"}
	mother := person{Gender: "F", CrimeStatus: "accomplice"}
	s := scenario{}
	s.Persons = append(s.Persons, father, mother)

	got := ruleOnePass(s)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestRuleTwoPass(t *testing.T) {
	father := person{Age: 44, Gender: "M", CrimeStatus: "accomplice"}
	mother := person{Age: 43, Gender: "F", CrimeStatus: "murderer"}
	son := person{Age: 12, Gender: "M", CrimeStatus: "victim"}
	daughter := person{Age: 5, Gender: "F", CrimeStatus: "witness"}

	s := scenario{}
	s.Persons = append(s.Persons, father, mother, son, daughter)

	got := ruleTwoPass(s)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestRuleTwoFail(t *testing.T) {
	father := person{Age: 44, Gender: "M", CrimeStatus: "accomplice"}
	mother := person{Age: 43, Gender: "F", CrimeStatus: "murderer"}
	son := person{Age: 12, Gender: "M", CrimeStatus: "witness"}
	daughter := person{Age: 5, Gender: "F", CrimeStatus: "victim"}

	s := scenario{}
	s.Persons = append(s.Persons, father, mother, son, daughter)

	got := ruleTwoPass(s)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
