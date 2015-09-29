package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

type Member struct {
	email     string
	id        int
	password  string
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
	result := Member{}

	db, err := getDBConnection()

	if err != nil {
		return result, errors.New("Unable to connect to database.\n" + err.Error())
	}

	defer db.Close()
	pwd := sha256.Sum256([]byte(password))
	row := db.QueryRow(`SELECT id, email, first_name
		FROM Member
		WHERE email = $1 AND password = $2`, email, hex.EncodeToString(pwd[:]))

	err = row.Scan(&result.id, &result.email, &result.firstName)
	if err != nil {
		return result, errors.New("Unable to find Member with email: " + email + "\n" + err.Error())
	}
	
	return result, nil
}

