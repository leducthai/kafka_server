package unionfind

type uniStorage[T comparable] struct {
	Parent  map[T]T
	Size    map[T]int
	NumEdge map[T]int
}

func (u *uniStorage[T]) Union(child T, parent T) bool {
	cp := u.Find(child)
	pp := u.Find(parent)

	if cp == pp {
		u.NumEdge[pp] += 1
	} else {
		u.Parent[cp] = pp
		u.Size[pp] += u.Size[cp]
		u.NumEdge[pp] += u.NumEdge[cp] + 1
	}

	return u.NumEdge[pp] <= u.Size[pp]
}

func (u *uniStorage[T]) Find(ele T) T {
	parent, ok := u.Parent[ele]
	if !ok {
		u.Parent[ele] = ele
		return u.Parent[ele]
	}

	if parent != ele {
		u.Parent[ele] = u.Find(parent)
	}
	return u.Parent[ele]
}

func NewUnionStorage[T comparable](elements []T) uniStorage[T] {
	parent := map[T]T{}
	size := map[T]int{}
	numEdge := map[T]int{}

	for _, v := range elements {
		size[v] = 1
		numEdge[v] = 0
	}

	return uniStorage[T]{
		Parent:  parent,
		Size:    size,
		NumEdge: numEdge,
	}
}
