package services

type Cache[T any] struct {
	cache map[string]T
}

func (self *Cache[T]) New() {
	self.cache = make(map[string]T)
}

func (self *Cache[T]) Add(uid string, obj T) {
	_, found := self.cache[uid]

	if !found {
		self.cache[uid] = obj
	}
}

func (self *Cache[T]) Get(uid string) (*T, bool) {
	obj, found := self.cache[uid]
	return &obj, found
}
