package models

import (
	m "github.com/andrewarrow/feedback/models"
)

type User struct {
	m.BaseModel
}

func NewUser(item map[string]any) *User {
	s := User{}
	s.Item = item
	return &s
}

func (u *User) NRG() string {
	return ""
}
