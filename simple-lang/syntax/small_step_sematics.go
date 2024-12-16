package syntax

//
// Primatives
//

// Boolean
func (b BooleanValue) IsReducable() bool {
	return false
}

func (b BooleanValue) Reduce(env *Environment) Expression {
	panic("Attempt to reduce a Boolean")
}

// Number

func (n NumberValue) IsReducable() bool {
	return false
}

func (n NumberValue) Reduce(env *Environment) Expression {
	panic("Attempt to reduce a Number")
}

//
// Compounds
//

// Add

func (a AddExpr) IsReducable() bool {
	return true
}

func (a AddExpr) Reduce(env *Environment) Expression {
	if a.left.IsReducable() {
		return Add(a.left.Reduce(env), a.right)
	} else if a.right.IsReducable() {
		return Add(a.left, a.right.Reduce(env))
	} else {
		left, ok := a.left.Value().(int)
		if !ok {
			panic("Left value is not an int")
		}
		right, ok := a.right.Value().(int)
		if !ok {
			panic("Right value is not an int")
		}
		res := Number(left + right)
		return res
	}
}

// And
func (a AndExpr) IsReducable() bool {
	return true
}

func (a AndExpr) Reduce(env *Environment) Expression {
	if a.left.IsReducable() {
		return And(a.left.Reduce(env), a.right)
	} else if a.right.IsReducable() {
		return And(a.left, a.right.Reduce(env))
	} else {
		left, ok := a.left.Value().(bool)
		if !ok {
			panic("Left value is not an bool")
		}
		right, ok := a.right.Value().(bool)
		if !ok {
			panic("Right value is not an bool")
		}
		res := Boolean(left && right)
		return res
	}
}

// Assign
func (a AssignStmt) IsReducable() bool {
	return true
}

func (a AssignStmt) Reduce(env *Environment) (Statement, *Environment) {
	if a.exp.IsReducable() {
		return Assign(a.name, a.exp.Reduce(env)), env
	} else {
		return DoNothing(), env.Merge(a.name, a.exp)
	}
}
