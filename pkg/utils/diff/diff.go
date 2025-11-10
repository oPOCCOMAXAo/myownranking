package diff

type Diff[E any] struct {
	Created []E
	Updated []E
	Deleted []E
}
