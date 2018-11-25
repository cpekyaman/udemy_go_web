package store

import d "github.com/cpekyaman/udemy_go_web/s09_session/domain"

type UserStore struct {
	store map[string]d.User
}

func (s *UserStore) Init() {
	s.store = make(map[string]d.User)
}

func (s *UserStore) Save(u d.User) {
	s.store[u.UserName] = u
}

func (s *UserStore) FindByUserName(userName string) (d.User, bool) {
	u, ok := s.store[userName]
	return u, ok
}

func (s *UserStore) Exists(userName string) bool {
	_, ok := s.FindByUserName(userName)
	return ok
}
