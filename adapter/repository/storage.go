package repository

type Storage interface {
	Find(out interface{}, conds ...interface{}) error
}
