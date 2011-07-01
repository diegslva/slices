package slices

import "fmt"

func F64List(n... float64) *F64Slice {
	return (*F64Slice)(&n)
}

type F64Slice	[]float64

func (s F64Slice) Len() int							{ return len(s) }
func (s F64Slice) Cap() int							{ return cap(s) }

func (s F64Slice) At(i int) interface{}				{ return s[i] }
func (s F64Slice) AtF32(i int) float32				{ return float32(s[i]) }
func (s F64Slice) AtF64(i int) float64				{ return s[i] }
func (s F64Slice) Set(i int, v interface{})			{ s[i] = v.(float64) }
func (s F64Slice) SetF32(i int, v float32)			{ s[i] = float64(v) }
func (s F64Slice) SetF64(i int, v float64)			{ s[i] = v }
func (s F64Slice) Clear(i int)						{ s[i] = 0 }
func (s F64Slice) Swap(i, j int)					{ s[i], s[j] = s[j], s[i] }

func (s F64Slice) Negate(i int)						{ s[i] = -s[i] }
func (s F64Slice) Increment(i int)					{ s[i] += 1 }
func (s F64Slice) Decrement(i int)					{ s[i] -= 1 }

func (s F64Slice) Add(i, j int)						{ s[i] += s[j] }
func (s F64Slice) Subtract(i, j int)				{ s[i] -= s[j] }
func (s F64Slice) Multiply(i, j int)				{ s[i] *= s[j] }
func (s F64Slice) Divide(i, j int)					{ s[i] /= s[j] }

func (s F64Slice) Less(i, j int) bool				{ return s[i] < s[j] }
func (s F64Slice) AtLeast(i, j int) bool			{ return s[i] <= s[j] }
func (s F64Slice) Same(i, j int) bool				{ return s[i] == s[j] }
func (s F64Slice) AtMost(i, j int) bool				{ return s[i] >= s[j] }
func (s F64Slice) More(i, j int) bool				{ return s[i] > s[j] }

func (s F64Slice) ZeroLess(i int) bool				{ return 0 < s[i] }
func (s F64Slice) ZeroAtLeast(i, j int) bool		{ return 0 <= s[j] }
func (s F64Slice) ZeroSame(i int) bool				{ return 0 == s[i] }
func (s F64Slice) ZeroAtMost(i, j int) bool			{ return 0 >= s[j] }
func (s F64Slice) ZeroMore(i int) bool				{ return 0 > s[i] }

func (s F64Slice) Compare(i, j int) (r int) {
	switch {
	case s[i] < s[j]:		r = IS_LESS_THAN
	case s[i] > s[j]:		r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s F64Slice) ZeroCompare(i int) (r int) {
	switch {
	case 0 < s[i]:			r = IS_LESS_THAN
	case 0 > s[i]:			r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s *F64Slice) Cut(i, j int) {
	a := *s
	l := len(a)
	if i < 0 {
		i = 0
	}
	if j > l {
		j = l
	}
	if j > i {
		if m := l - (j - i); m > 0 && l > m {
			copy(a[i:m], a[j:l])
			*s = a[0:m]
		}
	}
}

func (s *F64Slice) Delete(i int) {
	a := *s
	n := len(a)
	if i > -1 && i < n {
		copy(a[i:n - 1], a[i + 1:n])
		*s = a[0 : n - 1]
	}
}

func (s F64Slice) Each(f func(interface{})) {
	for _, v := range s {
		f(v)
	}
}

func (s F64Slice) EachWithIndex(f func(int, interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s F64Slice) EachWithKey(f func(key, value interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s F64Slice) EachF64(f func(float64)) {
	for _, v := range s {
		f(v)
	}
}

func (s F64Slice) EachF64WithIndex(f func(int, float64)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s F64Slice) EachF64WithKey(f func(interface{}, float64)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s F64Slice) String() (t string) {
	for _, v := range s {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%v)", t)
}

func (s F64Slice) BlockCopy(destination, source, count int) {
	end := source + count
	if end > len(s) {
		end = len(s)
	}
	copy(s[destination:], s[source:end])
}

func (s F64Slice) BlockClear(start, count int) {
	copy(s[start:], make(F64Slice, count, count))
}

func (s F64Slice) Overwrite(offset int, container interface{}) {
	switch container := container.(type) {
	case F64Slice:			copy(s[offset:], container)
	case []float64:			copy(s[offset:], container)
	}
}

func (s *F64Slice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)
	case capacity != cap(*s):	x := make(F64Slice, length, capacity)
								copy(x, *s)
								*s = x
	default:					*s = (*s)[:length]
	}
}

func (s *F64Slice) Extend(n int) {
	c := cap(*s)
	l := len(*s) + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *F64Slice) Expand(i, n int) {
	if i < 0 {
		i = 0
	}

	l := s.Len()
	if l < i {
		i = l
	}

	l += n
	c := s.Cap()
	if c < l {
		c = l
	}

	if c != s.Cap() {
		x := make(F64Slice, l, c)
		copy(x, (*s)[:i])
		copy(x[i + n:], (*s)[i:])
		*s = x
	} else {
		a := (*s)[:l]
		for j := l - 1; j >= i; j-- {
			a[j] = a[j - n]
		}
		*s = a
	}
}

func (s F64Slice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s F64Slice) Depth() int {
	return 0
}

func (s *F64Slice) Append(v interface{}) {
	s.AppendF64(v.(float64))
}

func (s *F64Slice) AppendF64(v float64) {
	*s = append(*s, v)
}

func (s *F64Slice) AppendSlice(o F64Slice) {
	*s = append(*s, o...)
}

func (s *F64Slice) Prepend(v interface{}) {
	s.PrependF64(v.(float64))
}

func (s *F64Slice) PrependF64(v float64) {
	l := s.Len() + 1
	n := make(F64Slice, l, l)
	n[0] = v
	copy(n[1:], *s)
	*s = n
}

func (s *F64Slice) PrependSlice(o F64Slice) {
	l := s.Len() + o.Len()
	n := make(F64Slice, l, l)
	copy(n, o)
	copy(n[o.Len():], *s)
	*s = n
}

func (s F64Slice) Subslice(start, end int) interface{} {
	return s[start:end]
}

func (s F64Slice) Repeat(count int) F64Slice {
	length := len(s) * count
	capacity := cap(s)
	if capacity < length {
		capacity = length
	}
	destination := make(F64Slice, length, capacity)
	for start, end := 0, len(s); count > 0; count-- {
		copy(destination[start:end], s)
		start = end
		end += len(s)
	}
	return destination
}

func (s *F64Slice) Flatten() {
	//	Flatten is a non-op for the F64Slice as they cannot contain nested elements
}

func (s F64Slice) equal(o F64Slice) (r bool) {
	switch {
	case s == nil:				r = o == nil
	case s.Len() == o.Len():	r = true
								for i, v := range s {
									if r = v == o[i]; !r {
										return
									}
								}
	}
	return
}

func (s F64Slice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *F64Slice:			r = o != nil && s.equal(*o)
	case F64Slice:			r = s.equal(o)
	case *[]float64:		r = o != nil && s.equal(*o)
	case []float64:			r = s.equal(o)
	}
	return
}

func (s F64Slice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s[0]
	}
	return
}

func (s F64Slice) Cdr() (t F64Slice) {
	if s.Len() > 1 {
		t = s[1:]
	}
	return
}

func (s *F64Slice) Rplaca(v interface{}) {
	switch {
	case s == nil:			*s = *F64List(v.(float64))
	case s.Len() == 0:		*s = append(*s, v.(float64))
	default:				(*s)[0] = v.(float64)
	}
}

func (s *F64Slice) Rplacd(v interface{}) {
	if s == nil {
		*s = *F64List(v.(float64))
	} else {
		ReplaceSlice := func(v F64Slice) {
			if l := len(v); l < cap(*s) {
				copy((*s)[1:], v)
			} else {
				l++
				n := make(F64Slice, l, l)
				copy(n, (*s)[:1])
				copy(n[1:], v)
				*s = n
			}
		}

		switch v := v.(type) {
		case *F64Slice:		ReplaceSlice(*v)
		case F64Slice:		ReplaceSlice(v)
		case *[]float64:	ReplaceSlice(F64Slice(*v))
		case []float64:		ReplaceSlice(F64Slice(v))
		case nil:			*s = (*s)[:1]
		default:			(*s)[1] = v.(float64)
							*s = (*s)[:2]
		}
	}
}