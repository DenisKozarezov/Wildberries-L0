package database

type IDatabaseRepository[T any] interface {
	SelectAll() ([]T, error)
	SelectByUID(uid string) (*T, error)
}
