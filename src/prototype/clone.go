package prototype

type cloneable interface {
	Clone() cloneable
}
