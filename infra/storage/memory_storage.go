package storage

import "github.com/vicren/go-clean/domain/entity"

type MemoryStorage struct {
	Users []*entity.User
}

func (s *MemoryStorage) Find(out interface{}, conds ...interface{}) error {
	if users, ok := out.([]*entity.User); ok {
		for _, u := range s.Users {
			users = append(users, u)
		}
	}
	return nil
}
