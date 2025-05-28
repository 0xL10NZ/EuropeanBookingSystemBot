package data

import (
	"errors"
	"sync"
)

type PaymentStatus int

const (
	StatusPendingPayment PaymentStatus = iota
	StatusPaid
)

type User struct {
	ID     int64
	Name   string
	Status PaymentStatus
	// можно добавить другие поля (страна, город, организация)
}

var (
	users   = make(map[int64]*User)
	usersMu sync.RWMutex
)

func GetUser(id int64) (*User, bool) {
	usersMu.RLock()
	defer usersMu.RUnlock()
	u, ok := users[id]
	return u, ok
}

func NewUser(id int64, name string) *User {
	return &User{
		ID:     id,
		Name:   name,
		Status: StatusPendingPayment,
	}
}

func SaveUser(u *User) error {
	if u == nil {
		return errors.New("nil user")
	}
	usersMu.Lock()
	defer usersMu.Unlock()
	users[u.ID] = u
	return nil
}

func GetAllUsers() []*User {
	usersMu.RLock()
	defer usersMu.RUnlock()
	all := make([]*User, 0, len(users))
	for _, u := range users {
		all = append(all, u)
	}
	return all
}
