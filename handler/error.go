package handler

import "net/http"

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Err  error
	Code int
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

func NewNotFound(err error) Error {
	return StatusError{
		Err:  err,
		Code: http.StatusNotFound,
	}
}
