package pred

import "fmt"

type and []Cond

var _ Cond = and{}

// And and condition
func And(conds ...Cond) Cond {
	var result = make(and, 0, len(conds))
	for _, cond := range conds {
		if cond == nil {
			continue
		}
		result = append(result, cond)
	}
	return result
}

// And pred or
func (a and) And(conds ...Cond) Cond {
	return And(a, And(conds...))
}

// Or pred or
func (a and) Or(conds ...Cond) Cond {
	return Or(a, Or(conds...))
}

// WriteTo write sql
func (a and) WriteTo(w Writer) error {
	for i, cond := range a {
		_, isOr := cond.(or)
		if isOr {
			fmt.Fprint(w, "(")
		}

		err := cond.WriteTo(w)
		if err != nil {
			return err
		}

		if isOr {
			fmt.Fprint(w, ")")
		}

		if i != len(a)-1 {
			fmt.Fprint(w, " AND ")
		}
	}
	return nil
}
