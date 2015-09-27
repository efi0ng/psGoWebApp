package models

import (

)

type Member struct {
	email string
	id int
	password string
	firstName string
}

func (this *Member) Email() string {
	return this.email
}

func (this *Member) Id() int {
	return this.id
}

func (this *Member) Password() string {
	return this.password
}

func (this *Member) FirstName() string {
	return this.firstName
}

func (this *Member) SetEmail(value string) {
	this.email = value
}

func (this *Member) SetId(value int) {
	this.id = value
}

func (this *Member) SetPassword(value string) {
	this.password = value
}

func (this *Member) SetFirstName(value string) {
	this.firstName = value
}

func GetMember(email string, password string) (Member, error) {
	return Member{}, nil
}