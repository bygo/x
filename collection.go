package x

type Collection[V comparable] struct {
	items []V
}

func Collect[V comparable](items []V) *Collection[V] {
	return &Collection[V]{items: items}
}

func (c *Collection[V]) Merge(items []V) *Collection[V] {
	var cNew = &Collection[V]{items: make([]V, len(c.items))}
	for idx := range c.items {
		cNew.items[idx] = c.items[idx]
	}
	cNew.items = append(cNew.items, items...)
	return cNew
}

func (c *Collection[V]) Replace(old V, new V, n int) *Collection[V] {
	var cNew = &Collection[V]{items: make([]V, len(c.items))}
	copy(cNew.items, c.items)
	for idx := range cNew.items {
		if n == 0 {
			break
		}

		if cNew.items[idx] == old {
			cNew.items[idx] = new
			n--
		}
	}
	return cNew
}

func (c *Collection[V]) Diff(items []V) *Collection[V] {
	itemMp := map[V]struct{}{}
	for _, item := range items {
		itemMp[item] = struct{}{}
	}
	cNew := &Collection[V]{items: []V{}}

	for _, item := range c.items {
		_, ok := itemMp[item]
		if !ok {
			cNew.items = append(cNew.items, item)
		}
	}
	return cNew
}

func (c *Collection[V]) ForEach(iteratee func(value V, key int)) *Collection[V] {
	itemsL := len(c.items)
	for idx := 0; idx < itemsL; idx++ {
		iteratee(c.items[idx], idx)
	}

	return c
}

func (c *Collection[V]) ForEachRight(iteratee func(value V, key int)) *Collection[V] {
	for idx := len(c.items) - 1; 0 <= idx; idx-- {
		iteratee(c.items[idx], idx)
	}
	return c
}

func (c *Collection[V]) Map(iteratee func(value V, key int) V) *Collection[V] {
	var cNew = &Collection[V]{items: make([]V, len(c.items))}
	for idx := range c.items {
		cNew.items[idx] = iteratee(c.items[idx], idx)
	}
	return cNew
}

func (c *Collection[V]) Filter(predicate func(value V, k int) bool) *Collection[V] {
	var cNew = &Collection[V]{items: []V{}}
	for k, v := range c.items {
		if !predicate(v, k) {
			continue
		}
		cNew.items = append(cNew.items, v)
	}
	return cNew
}

// Output

func (c *Collection[V]) ToSlice() []V {
	items := make([]V, len(c.items))
	copy(items, c.items)
	return items
}

func (c *Collection[V]) Chunk(size int) [][]V {
	itemsL := len(c.items)
	chunksL := (itemsL + size - 1) / size
	var chunks = make([][]V, chunksL)
	for idx := 0; idx < chunksL; idx++ {
		hi := (idx + 1) * size
		if itemsL < hi {
			hi = itemsL
		}
		chunks[idx] = append([]V{}, c.items[idx*size:hi]...)
	}
	return chunks
}

// Statistic

func (c *Collection[V]) Sum(iteratee func(value V) int) int {
	itemsL := len(c.items)
	var total int
	for idx := 0; idx < itemsL; idx++ {
		total += iteratee(c.items[idx])
	}
	return total
}

func (c *Collection[V]) Avg(iteratee func(value V) int) int {
	itemsL := len(c.items)
	var total int
	for idx := 0; idx < itemsL; idx++ {
		total += iteratee(c.items[idx])
	}
	return total / itemsL
}
