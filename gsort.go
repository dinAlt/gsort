//Package sort provides generic Sort method
package gsort

import "sort"

type sortable[S ~[]T, T any] struct {
  slc S
  cmpFunc func(T, T) bool
}

func (a sortable[S, T]) Len() int           { return len(a.slc) }
func (a sortable[S, T]) Swap(i, j int)      { a.slc[i], a.slc[j] = a.slc[j], a.slc[i] }
func (a sortable[S, T]) Less(i, j int) bool { return a.cmpFunc(a.slc[i], a.slc[j])}

// Sort slc with cmp function
func Sort[S ~[]T, T any](slc S, cmp func(T, T) bool) {
  s := sortable[[]T, T]{slc: slc, cmpFunc: cmp} 
  sort.Sort(s)
}
