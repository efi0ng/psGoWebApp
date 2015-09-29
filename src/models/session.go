package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
)

type Session struct {
	id int
	memberId int
	sessionId string
}

func (this *Session) Id() int {
	return this.id
}

func (this *Session) MemberId() int {
	return this.memberId
}

func (this *Session) SessionId() string {
	return this.sessionId
}

func CreateSession(member Member) (Session, error) {
	result := Session{}
	result.memberId = member.Id()
	sessionId := sha256.Sum256([]byte(member.Email() + time.Now().Format("15:04:05")))
	result.sessionId = hex.EncodeToString(sessionId[:])
	
	db, err := getDBConnection()
	if err != nil {
		return result, errors.New("Unable to connect to database.\n" + err.Error())
	}
	
	defer db.Close()
	sqlResult, err := db.Exec(`INSERT INTO Session
		(member_id, session_id)
		VALUES ($1, $2)`, member.Id(), result.sessionId)
	
	if err != nil {
		return Session{}, errors.New("Unable to create session\n" + err.Error())
	}
	
	id, err := sqlResult.LastInsertId()
	if err != nil {
		return Session{}, errors.New("Unable to create session. Id not returned.\n" + err.Error())		
	}
	
	result.id = int(id)
	return result, nil
}