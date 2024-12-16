package syntax

import (
	"fmt"
)

type SequenceStmt struct {
	first  Statement
	second Statement
}

func Sequence(first Statement, second Statement) *SequenceStmt {
	s := new(SequenceStmt)
	s.first = first
	s.second = second
	return s
}

func (s SequenceStmt) String() string {
	return fmt.Sprintf("%s ; %s", s.first, s.second)
}

func (s SequenceStmt) IsReducable() bool {
	return true
}

func (s SequenceStmt) Reduce(env *Environment) (Statement, *Environment) {
	if !s.first.IsReducable() {
		return s.second, env
	} else {
		new_first, env := s.first.Reduce(env)
		return Sequence(new_first, s.second), env
	}
}

func (s SequenceStmt) Evaluate(env *Environment) *Environment {
	return s.second.Evaluate(s.first.Evaluate(env))
}
