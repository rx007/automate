package storage

import (
	"errors"
)

// Error responses common to all storage adapters, be it v1, v2, memstore, postgres, etc.
var (
	// ErrNotFound is returned when a requested policy wasn't found.
	ErrNotFound = errors.New("not found")

	// ErrCannotDelete is thrown by our custom pg error from migration 02
	// if a user tries to delete a policy that is marked as non-deletable.
	ErrCannotDelete = errors.New("policy not deletable")

	// ErrConflict indicates that the object being created already exists.
	ErrConflict = errors.New("conflict")

	// ErrDatabase results from unexpected database errors.
	ErrDatabase = errors.New("database internal")

	// ErrGenerateUUID occurs when a UUID could not be generated for a new object.
	ErrGenerateUUID = errors.New("could not generate UUID")

	// ErrForeignKey occurs, among other times, when attempting to insert a foreign key whose referent does not exist
	ErrForeignKey = errors.New("foreign key violation")

	// ErrMaxProjectsExceeded indicates that a new project cannot be created
	// since the max allowed are already created.
	ErrMaxProjectsExceeded = errors.New("max projects allowed")

	// ErrChangeProjectForRule indicates that an update operation attempted to change
	// the project for a rule, which is not allowed.
	ErrChangeProjectForRule = errors.New("cannot change rule")
)

// ErrTxCommit occurs when the database attempts to commit a transaction and
// fails.
type ErrTxCommit struct {
	underlying error
}

func NewErrTxCommit(e error) error {
	return &ErrTxCommit{underlying: e}
}

func (e *ErrTxCommit) Error() string {
	return "commit db transaction: " + e.Error()
}

// ErrMissingField occurs when a required field was not passed.
type ErrMissingField struct {
	field string
}

func NewMissingFieldError(f string) error {
	return &ErrMissingField{field: f}
}

func (e *ErrMissingField) Error() string {
	return "must supply policy " + e.field
}
