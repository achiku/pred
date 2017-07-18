package pred

import "fmt"

type or []Cond

var _ Cond = or{}

// Or or condition
func Or(conds ...Cond) Cond {
	var result = make(or, 0, len(conds))
	for _, cond := range conds {
		if cond == nil {
			continue
		}
		result = append(result, cond)
	}
	return result
}

// And pred or
func (o or) And(conds ...Cond) Cond {
	return And(o, And(conds...))
}

// Or pred or
func (o or) Or(conds ...Cond) Cond {
	return Or(o, Or(conds...))
}

// WriteTo write sql
func (o or) WriteTo(w Writer) error {
	for i, cond := range o {
		var needQuote bool
		switch cond.(type) {
		case and:
			needQuote = true
			// case Eq:
			// 	needQuote = (len(cond.(Eq)) > 1)
		}

		if needQuote {
			fmt.Fprint(w, "(")
		}

		err := cond.WriteTo(w)
		if err != nil {
			return err
		}

		if needQuote {
			fmt.Fprint(w, ")")
		}

		if i != len(o)-1 {
			fmt.Fprint(w, " OR ")
		}
	}
	return nil
}
