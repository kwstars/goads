package arraylist

var (
	IntsCmp = func(a, b int) int8 {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	}
)
