package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/bbolt"
)

type SessionRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailEmail(email string) (model.Session, error)
	SessionAvailToken(token string) (model.Session, error)
	TokenExpired(session model.Session) bool
}

type sessionsRepo struct {
	filebasedDb *filebased.Data
}

func NewSessionsRepo(filebasedDb *filebased.Data) *sessionsRepo {
	return &sessionsRepo{filebasedDb}
}

func (u *sessionsRepo) AddSessions(session model.Session) error {
	sessionJSON, err := json.Marshal(session)
	if err != nil {
		return err
	}
	return u.filebasedDb.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Sessions"))
		return b.Put([]byte(session.Token), sessionJSON)
	})
}

func (u *sessionsRepo) DeleteSession(token string) error {
	return u.filebasedDb.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Sessions"))
		return b.Delete([]byte(token))
	})
}

func (u *sessionsRepo) UpdateSessions(session model.Session) error {
	sessionJSON, err := json.Marshal(session)
	if err != nil {
		return err
	}
	return u.filebasedDb.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Sessions"))
		if b == nil {
			return fmt.Errorf("sessions bucket not found")
		}
		return b.Put([]byte(session.Token), sessionJSON)
	})
}

func (u *sessionsRepo) SessionAvailEmail(email string) (model.Session, error) {
	var session model.Session
	found := false // Flag to check if at least one session matches the email

	err := u.filebasedDb.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Sessions"))
		if b == nil {
			return fmt.Errorf("sessions bucket not found")
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var s model.Session
			if err := json.Unmarshal(v, &s); err != nil {
				continue // Skip badly formatted session records
			}
			if s.Email == email {
				session = s
				found = true
				break // Stop the iteration as we found the session
			}
		}
		return nil
	})

	if err != nil {
		return model.Session{}, err // Return error encountered during the View transaction
	}

	if !found {
		return model.Session{}, fmt.Errorf("no session available for email: %s", email) // No session was found for the given email
	}

	return session, nil // Return the found session
}

func (u *sessionsRepo) SessionAvailToken(token string) (model.Session, error) {
	session, err := u.filebasedDb.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Sessions"))
		if b == nil {
			return fmt.Errorf("sessions bucket not found")
		}

		sessionData := b.Get([]byte(token))
		if sessionData == nil {
			return fmt.Errorf("session not found")
		}

		var s model.Session
		if err := json.Unmarshal(sessionData, &s); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return model.Session{}, err
	}

	return session, nil
}

func (u *sessionsRepo) TokenValidity(token string) (model.Session, error) {
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(session) {
		err := u.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, err
	}

	return session, nil
}

func (u *sessionsRepo) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}
