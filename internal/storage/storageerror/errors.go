package storageerror

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrEmptyUserStorage = errors.New("user storage is empty")

	ErrProductIDNotFound   = errors.New("product id not found")
	ErrProductAlreadyExist = errors.New("product already exist")
	ErrEmptyStorage        = errors.New("storage is empty")

	ErrNoPurchasesFound = errors.New("no purchases found")
)
