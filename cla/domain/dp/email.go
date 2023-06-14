package dp

import (
	"errors"
	"regexp"
)

var reEmailAddr = regexp.MustCompile(`^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*(\.[a-zA-Z]{2,6})$`)

// EmailAddr
type EmailAddr interface {
	EmailAddr() string
}

func NewEmailAddr(v string) (EmailAddr, error) {
	if v == "" || !reEmailAddr.MatchString(v) {
		return nil, errors.New("invalid email address")
	}

	return emailAddr(v), nil
}

type emailAddr string

func (r emailAddr) EmailAddr() string {
	return string(r)
}
