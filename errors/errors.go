package errors

import "fmt"


type TypeError struct {
	Details string
}

func (e *TypeError) Error() string {
	return fmt.Sprintf("TypeError: %s", e.Details)
}
type DimensionError struct {
	Details string
}

func (e *DimensionError) Error() string {
	return fmt.Sprintf("DimensionError: %s", e.Details)
}

type DivisionByZeroError struct {
	Details string
}

func (e *DivisionByZeroError) Error() string {
	return fmt.Sprintf("DivisionByZeroError: %s", e.Details)
}

type OutOfBoundsError struct {
	Details string
}

func (e *OutOfBoundsError) Error() string {
	return fmt.Sprintf("OutofBoundsError: %s", e.Details)
}