package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"time"
)

type SessionRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailEmail(email string) (model.Session, error)
	SessionAvailToken(token string) (model.Session, error)
	TokenValidity(token string) (model.Session, error)
}

type sessionsRepo struct {
	filebasedDb *filebased.Data
}

func NewSessionsRepo(filebasedDb *filebased.Data) *sessionsRepo {
	return &sessionsRepo{filebasedDb}
}

func (s *sessionsRepo) AddSessions(session model.Session) error {
	// Implementasi untuk menyimpan data sesuai parameter ke tabel sessions
	return nil // TODO:
}

func (s *sessionsRepo) DeleteSession(token string) error {
	// Implementasi untuk menghapus data sesuai target token dari parameter
	return nil // TODO:
}

func (s *sessionsRepo) UpdateSessions(session model.Session) error {
	// Implementasi untuk mengubah data session sesuai parameter ke tabel sessions dengan kondisi sama antara email parameter dengan database
	return nil // TODO:
}

func (s *sessionsRepo) SessionAvailEmail(email string) (model.Session, error) {
	// Implementasi untuk memeriksa apakah token tersedia pada tabel sessions sesuai dengan kolom email sama dengan nilai dari parameter
	return model.Session{}, nil //
}

func (s *sessionsRepo) SessionAvailToken(token string) (model.Session, error) {
	// Implementasi untuk memeriksa apakah token tersedia pada tabel sessions sesuai dengan kolom token sama dengan nilai dari parameter
	return model.Session{}, nil // TODO:
}

func (s *sessionsRepo) TokenValidity(token string) (model.Session, error) {
	// Implementasi untuk memvalidasi token
	session, err := s.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if s.TokenExpired(session) {
		err := s.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, err
	}

	return session, nil
}

func (s *sessionsRepo) TokenExpired(session model.Session) bool {
	// Implementasi untuk mengecek apakah token sudah kadaluarsa
	return session.Expiry.Before(time.Now())
}
