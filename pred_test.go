package main

import "testing"

func TestEq(t *testing.T) {
	pred := Eq("name", "moqada")
	expected := "name = 'moqada'"
	if result := pred.ToSQL(); result != expected {
		t.Errorf("want %s got %s", expected, result)
	}
}

func TestNotEq(t *testing.T) {
	pred := NotEq("name", "moqada")
	expected := "name != 'moqada'"
	if result := pred.ToSQL(); result != expected {
		t.Errorf("want %s got %s", expected, result)
	}
}

func TestAnd(t *testing.T) {
	q := Predicates{}
	pred := q.And(Eq("name", "moqada"), Eq("status", "active"), Eq("age", 32))
	expected := "name = 'moqada' and status = 'active' and age = 32"
	if result := pred.ToSQL(); result != expected {
		t.Errorf("want %s got %s", expected, result)
	}
}

func TestNestedConjunction(t *testing.T) {
	q := Predicates{}
	pred := q.And(Eq("name", "moqada"), Eq("age", 32)).Or(Eq("status", "active"))
	expected := "( name = 'moqada' and age = 32 ) or status = 'active'"
	if result := pred.ToSQL(); result != expected {
		t.Errorf("want %s got %s", expected, result)
	}
}
