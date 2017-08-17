package iterator

type Iterator interface {
	Count() int
	Next() string
	Remove()
}
