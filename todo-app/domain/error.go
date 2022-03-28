package domain

import "fmt"

type NotFoundError struct {
	ID int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Todo %d is not found", e.ID)
}

func NewNotFoundError(id int) *NotFoundError {
	return &NotFoundError{
		ID: id,
	}
}

type InfraError struct {
	err error
}

func (e *InfraError) Error() string {
	return e.err.Error()
}

func NewInfraError(err error) *InfraError {
	return &InfraError{
		err: err,
	}
}

type InvalidRequestError struct {
	message string
}

func (e *InvalidRequestError) Error() string {
	return e.message
}

func NewInvalidRequestError(message string) *InvalidRequestError {
	return &InvalidRequestError{
		message: message,
	}
}
