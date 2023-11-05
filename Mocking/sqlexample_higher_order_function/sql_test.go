package sqlexample_test

import (
	"database/sql"
	"errors"
	sqlexample "livrogolang/Mocking/sqlexample_higher_order_function"
	"testing"
)

func TestOpenDB(t *testing.T) {
	mockError := errors.New("uh oh")
	subtests := []struct {
		name        string
		u, p, a, db string
		sqlOpener   func(string, string) (*sql.DB, error)
		expectedErr error
	}{
		{
			name: "happy-path",
			u:    "user",
			p:    "password",
			a:    "addr",
			db:   "db",
			sqlOpener: func(driver string, source string) (db *sql.DB, err error) {
				if source != "user:password@addr/db" {
					return nil, errors.New("wrong connection string")
				}
				return nil, nil
			},
		},
		{
			name: "error-from-sqlOpener",
			sqlOpener: func(driver string, source string) (db *sql.DB, err error) {
				return nil, mockError
			},
			expectedErr: mockError,
		},
	}
	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			_, err := sqlexample.OpenDB(subtest.u, subtest.p, subtest.a, subtest.db, subtest.sqlOpener)
			if !errors.Is(err, subtest.expectedErr) {
				t.Errorf("expected error [%v], got error (%v)", subtest.expectedErr, err)
			}
		})
	}
}
