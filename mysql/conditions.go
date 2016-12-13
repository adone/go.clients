package mysql

import (
	"github.com/jinzhu/gorm"
)

type Condition func(*gorm.DB) *gorm.DB

func NewConditions(scopes ...func(*gorm.DB) *gorm.DB) *Conditions {
	return &Conditions{scopes}
}

type Conditions struct {
	scopes []func(*gorm.DB) *gorm.DB
}

func (conditions *Conditions) Add(scope Condition) *Conditions {
	return &Conditions{append(conditions.scopes, scope)}
}

func (conditions *Conditions) Apply(storage *gorm.DB) *gorm.DB {
	if conditions == nil {
		return storage
	}

	return storage.Scopes(conditions.scopes...)
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
