package model

import (
	"context"
	"time"
)

type Semaphore struct {
	key  string
	user string
	ttl  time.Duration
}

func (m *Model) LockIfNotExists(ctx context.Context, lockTarget, user string, ttl time.Duration) (bool, string, time.Time, error) {
	isLocked, err := m.redis.SetNX(ctx, lockTarget, user, ttl).Result()
	if err != nil {
		return false, "", time.Now(), err
	}

	remainingTTL, err := m.redis.TTL(ctx, lockTarget).Result()
	if err != nil {
		return false, "", time.Now(), err
	}

	expireDate := time.Now().Add(remainingTTL)

	if !isLocked {
		lockedUser, err := m.redis.Get(ctx, lockTarget).Result()
		if err != nil {
			return false, "", time.Now(), err
		}

		return user == lockedUser, lockedUser, expireDate, err
	}

	return true, user, expireDate, nil
}
