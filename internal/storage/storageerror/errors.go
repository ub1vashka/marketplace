package storageerror

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrEmptyUserStorage = errors.New("user storage is empty")

	ErrProductIDNotFound  = errors.New("product id not found")
	ErrProductAlredyExist = errors.New("product alredy exist")
	ErrEmptyStorage       = errors.New("storage is empty")
)
