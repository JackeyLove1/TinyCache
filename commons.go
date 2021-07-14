package TinyCache

import (
	"errors"
	"math"
)

var (
	ErrInvalidMaxBytes = errors.New("MaxBytes set error")
	ErrInvalidCache    = errors.New("Cache is nil")
	ErrKeyNotFound     = errors.New("Key Not Found")
)

// Value is a interface for value type
type Value interface {
	Len() int
}

// Just for testing
type String string

func (s String) Len() int {
	return len(s)
}

// Const Value Just For Testing
const (
	MaxCacheSize = int64(math.MaxInt32)
	DefaultLRUK  = 2
)
