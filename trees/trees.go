package trees

import "github.com/kwstars/goads/containers"

// Tree interface that all trees implement
type Tree[T any] interface {
	containers.Container[T]
}
