package pred

import "testing"

func TestExp(t *testing.T) {
	cases := []struct {
		Col string
		Ope OperatorType
		Val interface{}
	}{
		{Col: "name", Ope: "=", Val: "moqada"},
	}

	for _, c := range cases {
		e := Eq(c.Col, c.Val).And(Eq(c.Col, c.Val)).And(Eq("test", "value"))
		t.Logf("%s", e)
	}
}

func TestAnd(t *testing.T) {
	p := And(Eq("a", 1), NotEq("b", 2)).Or(IsNull("c"))
	t.Logf("%s", p)
}
