package main

import (
	"errors"
	"strings"
)

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")

// LowercaseService models a service which convert a string to lowercase
type LowercaseService interface {
	Lowercase(string) (string, error)
}

type lowercaseService struct{}

// ServiceMiddleware is a convinient type to expose service middleware
type ServiceMiddleware func(LowercaseService) LowercaseService

func (lowercaseService) Lowercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToLower(s), nil
}
