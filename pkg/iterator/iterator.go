package iterator

type Iterable interface {
	Exists() bool
	Next()
	Get() string
}

type Store struct {
	Fruits []string
	index int
}

func (r Store) Exists() bool {
	return r.index < len(r.Fruits)
}

func (r Store) Get() string {
	return r.Fruits[r.index]
}

func (r *Store) Next() {
	r.index++
}

