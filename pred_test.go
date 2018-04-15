package main

import "testing"

func TestEq(t *testing.T) {
	pred := Eq("name", "moqada")
	expected := "name = 'moqada'"
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

func TestNestedAnd(t *testing.T) {
	q := Predicates{}
	pred := q.And(Eq("name", "moqada"), Eq("age", 32)).And(Eq("status", "active"))
	t.Logf("%s", pred.ToSQL())
}
