package syntax

// Boolean
func (b BooleanValue) Evaluate(env *Environment) Expression {
	return b
}

// Number
func (n NumberValue) Evaluate(env *Environment) Expression {
	return n
}

//
// Compounds
//

// Add
func (a AddExpr) Evaluate(env *Environment) Expression {
	return Number(a.left.Evaluate(env).Value().(int) + a.right.Evaluate(env).Value().(int))
}

// And
func (a AndExpr) Evaluate(env *Environment) Expression {
	return Boolean(a.left.Evaluate(env).Value().(bool) && a.right.Evaluate(env).Value().(bool))
}

// Assign
func (a AssignStmt) Evaluate(env *Environment) *Environment {
	return env.Merge(a.name, a.exp.Evaluate(env))
}
