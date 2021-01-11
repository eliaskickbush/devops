package cache

import "errors"

// Dao errors
var ErrSettingValue = errors.New("there was an error setting cache value")
var ErrGettingValue = errors.New("there was an error getting cache value")
var ErrDeletingValue = errors.New("there was an error deleting cache value")