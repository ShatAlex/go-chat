package handler

import (
	"errors"
)

func (h *Handler) validatePassword(pass1 string, pass2 string) error {
	if pass1 != pass2 {
		return errors.New("mismatched passwords")
	}

	if pass1 == "" || len(pass1) < 10 {
		return errors.New("bad password")
	}
	return nil
}
