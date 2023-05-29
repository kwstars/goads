package singlylinkedlist

var (
	IntCmp = func(a, b int) int8 {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	}
	StringCmp = func(a, b string) int8 {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	}
)
