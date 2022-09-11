package iter

type IteratorObj interface {
	HasNext() bool
	Next() int
}

type accumulateIterator struct {
	index   int
	length  int
	current int
	data    *[]int
	f       func(int, int) int
}

// 两位二元函数计算迭代
func (a *accumulateIterator) HasNext() bool { return a.index < a.length }
func (a *accumulateIterator) Next() int {
	a.current = a.f(a.current, (*a.data)[a.index])
	a.index++
	return a.current
}
func Accumulate(i *[]int, f ...func(a, b int) int) IteratorObj {
	a := &accumulateIterator{
		data:    i,
		length:  len(*i),
		index:   0,
		current: 0}
	if len(f) > 0 {
		a.f = f[0]
	} else {
		a.f = func(i1, i2 int) int { return i1 + i2 }
	}
	return a
}
func Chain() {

}
func Combinations()                {}
func CombinationsWithReplacement() {}
func Compress()                    {}

// 无限计数迭代
type countIterator struct {
	start int
	step  int
}

func (c *countIterator) HasNext() bool { return true }
func (c *countIterator) Next() int {
	result := c.start
	c.start += c.step
	return result
}
func Count(start int, step ...int) IteratorObj {
	c := &countIterator{
		start: 0,
		step:  1,
	}
	if len(step) > 0 {
		c.step = step[0]
	}
	return c
}

type cycleIterator struct {
	data      *[]int
	length    int
	index     int
	repeatNum int
	repeat    int
	hasRepeat bool
}

func (c *cycleIterator) HasNext() bool {
	return c.repeatNum < c.repeat
}
func (c *cycleIterator) Next() int {
	result := (*c.data)[c.index]
	c.index = (c.index + 1) % c.length
	if c.index == 0 && c.hasRepeat {
		c.repeatNum += 1
	}
	return result
}

func Cycle(i *[]int, repeat ...int) IteratorObj {

	c := &cycleIterator{
		data: i, hasRepeat: false, length: len(*i),
	}
	if len(repeat) > 0 {
		c.hasRepeat = true
		c.repeat = repeat[0]
	}
	return c
}
func Dropwhile()   {}
func FilterFalse() {}

type groupByIterator struct {
}

func GroupBy() {

}
func Islice()       {}
func Permutations() {}
func Product()      {}
func Repeat()       {}
func Starmap()      {}
func Takewhile()    {}
func Tee()          {}
func ZipLongest()   {}
