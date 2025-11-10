package diff

import (
	"github.com/samber/lo"
)

// OrderedSlice computes the diff between two ordered slices of comparable elements.
// It marks elements as Created, Updated or Deleted based on their order and equality.
// Elements are Updated if its position matches and equals(old, new) == false.
// Elements are Deleted if they exceed the length of the new slice.
// Elements are Created if they exist in newSlice but not in oldSlice.
//
// Best suited for inplace updates where order matters.
func OrderedSlice[E comparable, S ~[]E](
	oldSlice S,
	newSlice S,
	prepareForComparison func(E, E),
	equals func(E, E) bool,
) Diff[E] {
	res := Diff[E]{
		Updated: make([]E, 0, len(oldSlice)),
	}

	if equals == nil {
		equals = func(a, b E) bool { return a == b }
	}

	minLen := lo.Min([]int{len(oldSlice), len(newSlice)})

	for i := range minLen {
		oldElem := oldSlice[i]
		newElem := newSlice[i]

		prepareForComparison(newElem, oldElem)

		if !equals(newElem, oldElem) {
			res.Updated = append(res.Updated, newElem)
		}
	}

	if len(newSlice) > minLen {
		res.Created = append(res.Created, newSlice[minLen:]...)
	}

	if len(oldSlice) > minLen {
		res.Deleted = append(res.Deleted, oldSlice[minLen:]...)
	}

	return res
}
