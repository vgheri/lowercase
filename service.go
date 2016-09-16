package main

import (
	"errors"
	"strings"

	"golang.org/x/net/context"
)

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")

// LowercaseService models a service which convert a string to lowercase
type LowercaseService interface {
	Lowercase(context.Context, string) (string, error)
}

type lowercaseService struct{}

// ServiceMiddleware is a convinient type to expose service middleware
type ServiceMiddleware func(LowercaseService) LowercaseService

func (lowercaseService) Lowercase(ctx context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToLower(s), nil
}
