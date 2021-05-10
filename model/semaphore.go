package model

import (
	"time"
)

type Semaphore struct {
	key  string
	user string
	ttl  time.Duration
}

func (s *Semaphore) SetIfNotExists(key, user string, ttl time.Duration) (bool, string, time.Time, error) {
	expireDate := time.Now().Add(ttl)
	return true, user, expireDate, nil
}
