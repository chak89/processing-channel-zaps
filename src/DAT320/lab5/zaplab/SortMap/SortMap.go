// SortMap

/**************************************************************************/
/****************** Algorithm for sorting MAP values **********************/
/************************************************************************/
package SortMap

import (
	"sort"
)

type sortI struct {
	l    int
	less func(int, int) bool
	swap func(int, int)
}

func (s *sortI) Len() int {
	return s.l
}

func (s *sortI) Less(i, j int) bool {
	return s.less(i, j)
}

func (s *sortI) Swap(i, j int) {
	s.swap(i, j)
}

// SortM sorts the data defined by the length, Less and Swap functions.
func SortM(Len int, Less func(int, int) bool, Swap func(int, int)) {
	sort.Sort(&sortI{l: Len, less: Less, swap: Swap})
}

// Mapvalues sorts values in descending order.
func Mapvalues(mvalues []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(mvalues)))
	return mvalues
}
