package postgres

import (
	"gopkg.in/pg.v5/orm"
)

type Condition func(*orm.Query) (*orm.Query, error)

func NewConditions(scopes ...Condition) *Conditions {
	return &Conditions{scopes}
}

type Conditions struct {
	scopes []Condition
}

// Add store provided scopes in Conditions object
func (conditions *Conditions) Add(scopes ...Condition) *Conditions {
	conditions.scopes = append(conditions.scopes, scopes...)
	return conditions
}

func (conditions *Conditions) Apply(query *orm.Query) *orm.Query {
	if conditions == nil {
		return query
	}
	for _, f := range conditions.scopes {
		query = query.Apply(f)
	}
	return query
}

func (conditions *Conditions) Empty() bool {
	if conditions == nil {
		return true
	}

	return len(conditions.scopes) == 0
}

func (conditions *Conditions) Clone() *Conditions {
	newConditions := &Conditions{make([]Condition, len(conditions.scopes))}
	copy(newConditions.scopes, conditions.scopes)
	return newConditions
}
