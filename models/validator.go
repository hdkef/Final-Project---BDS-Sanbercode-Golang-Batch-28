package models

import "errors"

func MustNotEmpty(param ...string) error {
	for _, v := range param {
		if v == "" {
			return errors.New("some value cannot be empty")
		}
	}
	return nil
}
