package main

import (
	"fmt"
)

//    q.Eq("name", "moqada")
//    // where name = 'moqada'
//
//    q.And(Eq("name", "moqada"), Eq("status", "active"))
//    // where name = 'moqada' and status = 'active'
//
//    q.And(Eq("name", "moqada"), NotEq("age", 10)).And(Eq("status", "active")).
//    // where (name = 'moqada' and age = 10) and status = 'active'
//
//    q.Or(And(Eq("name", "moqada"), NotEq("age", 10)), Eq("name", "achiku")).
//    And(Eq("status", "active")).
//    // where ((name = 'moqada' and age = 10) or (name = 'achiku')) and status = 'active'

// Condition SQL conditions
type Condition interface {
	ToSQL() string
}

// Conjunction conjunction type
type Conjunction string

const (
	andConjunction = "and"
	orConjunction  = "or"
)

// Predicates predicates
type Predicates struct {
	conditions map[Conjunction][]Condition
	siblings   map[Conjunction][]Condition
}

// ToSQL predicates to SQL conditions
func (ps *Predicates) ToSQL() string {
	var s string
	for conj, preds := range ps.conditions {
		for i, p := range preds {
			if i == 0 {
				s = fmt.Sprintf("%s", p.ToSQL())
			} else {
				s = fmt.Sprintf("%s %s %s", s, conj, p.ToSQL())
			}
		}
	}
	if ps.siblings != nil {
		s = fmt.Sprintf("( %s )", s)
		for conj, preds := range ps.siblings {
			for _, p := range preds {
				s = fmt.Sprintf("%s %s %s", s, conj, p.ToSQL())
			}
		}
	}
	return s
}

func conj(ps *Predicates, con Conjunction, preds ...Condition) *Predicates {
	// too naive
	if ps.conditions == nil {
		ps.conditions = map[Conjunction][]Condition{con: []Condition{}}
		for _, p := range preds {
			ps.conditions[con] = append(ps.conditions[con], p)
		}
	} else {
		ps.siblings = map[Conjunction][]Condition{con: []Condition{}}
		for _, p := range preds {
			ps.siblings[con] = append(ps.siblings[con], p)
		}
	}
	return ps
}

// And conjunction and
func (ps *Predicates) And(preds ...Condition) *Predicates {
	return conj(ps, andConjunction, preds...)
}

// Or conjunction and
func (ps *Predicates) Or(preds ...Condition) *Predicates {
	return conj(ps, orConjunction, preds...)
}

// Predicate predicate
type Predicate struct {
	Subject string
	Verb    string
	Object  interface{}
}

// ToSQL predicate to sql
func (p *Predicate) ToSQL() string {
	var s string
	switch v := p.Object.(type) {
	case string:
		s = fmt.Sprintf("'%s'", v)
	case int, int32, int64:
		s = fmt.Sprintf("%d", v)
	}
	return fmt.Sprintf("%s %s %s", p.Subject, p.Verb, s)
}

// Eq return equeals predicates
func Eq(subj string, obj interface{}) *Predicate {
	return &Predicate{
		Subject: subj,
		Verb:    "=",
		Object:  obj,
	}
}

// NotEq return not equeals predicates
func NotEq(subj string, obj interface{}) *Predicate {
	return &Predicate{
		Subject: subj,
		Verb:    "!=",
		Object:  obj,
	}
}
