package slices

import "testing"

func TestU32SliceString(t *testing.T) {
	ConfirmString := func(s U32Slice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(U32Slice{}, "()")
	ConfirmString(U32Slice{0}, "(0)")
	ConfirmString(U32Slice{0, 1}, "(0 1)")
}

func TestU32SliceLen(t *testing.T) {
	ConfirmLength := func(s U32Slice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(U32Slice{0}, 1)
	ConfirmLength(U32Slice{0, 1}, 2)
}

func TestU32SliceSwap(t *testing.T) {
	ConfirmSwap := func(s U32Slice, i, j int, r U32Slice) {
		if s.Swap(i, j); !r.Equal(s) {
			t.Fatalf("Swap(%v, %v) should be %v but is %v", i, j, r, s)
		}
	}
	ConfirmSwap(U32Slice{0, 1, 2}, 0, 1, U32Slice{1, 0, 2})
	ConfirmSwap(U32Slice{0, 1, 2}, 0, 2, U32Slice{2, 1, 0})
}

func TestU32SliceCompare(t *testing.T) {
	ConfirmCompare := func(s U32Slice, i, j, r int) {
		if x := s.Compare(i, j); x != r {
			t.Fatalf("Compare(%v, %v) should be %v but is %v", i, j, r, x)
		}
	}

	ConfirmCompare(U32Slice{0, 1}, 0, 0, IS_SAME_AS)
	ConfirmCompare(U32Slice{0, 1}, 0, 1, IS_LESS_THAN)
	ConfirmCompare(U32Slice{0, 1}, 1, 0, IS_GREATER_THAN)
}

func TestU32SliceZeroCompare(t *testing.T) {
	ConfirmCompare := func(s U32Slice, i, r int) {
		if x := s.ZeroCompare(i); x != r {
			t.Fatalf("ZeroCompare(%v) should be %v but is %v", i, r, x)
		}
	}

	ConfirmCompare(U32Slice{0, 1, 2}, 0, IS_SAME_AS)
	ConfirmCompare(U32Slice{0, 1, 2}, 1, IS_LESS_THAN)
	ConfirmCompare(U32Slice{0, 1, 2}, 2, IS_LESS_THAN)
}

func TestU32SliceCut(t *testing.T) {
	ConfirmCut := func(s U32Slice, start, end int, r U32Slice) {
		if s.Cut(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmCut(U32Slice{0, 1, 2, 3, 4, 5}, 0, 1, U32Slice{1, 2, 3, 4, 5})
	ConfirmCut(U32Slice{0, 1, 2, 3, 4, 5}, 1, 2, U32Slice{0, 2, 3, 4, 5})
	ConfirmCut(U32Slice{0, 1, 2, 3, 4, 5}, 2, 3, U32Slice{0, 1, 3, 4, 5})
	ConfirmCut(U32Slice{0, 1, 2, 3, 4, 5}, 3, 4, U32Slice{0, 1, 2, 4, 5})
	ConfirmCut(U32Slice{0, 1, 2, 3, 4, 5}, 4, 5, U32Slice{0, 1, 2, 3, 5})
	ConfirmCut(U32Slice{0, 1, 2, 3, 4, 5}, 5, 6, U32Slice{0, 1, 2, 3, 4})

	ConfirmCut(U32Slice{0, 1, 2, 3, 4, 5}, -1, 1, U32Slice{1, 2, 3, 4, 5})
	ConfirmCut(U32Slice{0, 1, 2, 3, 4, 5}, 0, 2, U32Slice{2, 3, 4, 5})
	ConfirmCut(U32Slice{0, 1, 2, 3, 4, 5}, 1, 3, U32Slice{0, 3, 4, 5})
	ConfirmCut(U32Slice{0, 1, 2, 3, 4, 5}, 2, 4, U32Slice{0, 1, 4, 5})
	ConfirmCut(U32Slice{0, 1, 2, 3, 4, 5}, 3, 5, U32Slice{0, 1, 2, 5})
	ConfirmCut(U32Slice{0, 1, 2, 3, 4, 5}, 4, 6, U32Slice{0, 1, 2, 3})
	ConfirmCut(U32Slice{0, 1, 2, 3, 4, 5}, 5, 7, U32Slice{0, 1, 2, 3, 4})
}

func TestU32SliceTrim(t *testing.T) {
	ConfirmTrim := func(s U32Slice, start, end int, r U32Slice) {
		if s.Trim(start, end); !r.Equal(s) {
			t.Fatalf("Trim(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmTrim(U32Slice{0, 1, 2, 3, 4, 5}, 0, 1, U32Slice{0})
	ConfirmTrim(U32Slice{0, 1, 2, 3, 4, 5}, 1, 2, U32Slice{1})
	ConfirmTrim(U32Slice{0, 1, 2, 3, 4, 5}, 2, 3, U32Slice{2})
	ConfirmTrim(U32Slice{0, 1, 2, 3, 4, 5}, 3, 4, U32Slice{3})
	ConfirmTrim(U32Slice{0, 1, 2, 3, 4, 5}, 4, 5, U32Slice{4})
	ConfirmTrim(U32Slice{0, 1, 2, 3, 4, 5}, 5, 6, U32Slice{5})

	ConfirmTrim(U32Slice{0, 1, 2, 3, 4, 5}, -1, 1, U32Slice{0})
	ConfirmTrim(U32Slice{0, 1, 2, 3, 4, 5}, 0, 2, U32Slice{0, 1})
	ConfirmTrim(U32Slice{0, 1, 2, 3, 4, 5}, 1, 3, U32Slice{1, 2})
	ConfirmTrim(U32Slice{0, 1, 2, 3, 4, 5}, 2, 4, U32Slice{2, 3})
	ConfirmTrim(U32Slice{0, 1, 2, 3, 4, 5}, 3, 5, U32Slice{3, 4})
	ConfirmTrim(U32Slice{0, 1, 2, 3, 4, 5}, 4, 6, U32Slice{4, 5})
	ConfirmTrim(U32Slice{0, 1, 2, 3, 4, 5}, 5, 7, U32Slice{5})
}

func TestU32SliceDelete(t *testing.T) {
	ConfirmDelete := func(s U32Slice, index int, r U32Slice) {
		if s.Delete(index); !r.Equal(s) {
			t.Fatalf("Delete(%v) should be %v but is %v", index, r, s)
		}
	}

	ConfirmDelete(U32Slice{0, 1, 2, 3, 4, 5}, -1, U32Slice{0, 1, 2, 3, 4, 5})
	ConfirmDelete(U32Slice{0, 1, 2, 3, 4, 5}, 0, U32Slice{1, 2, 3, 4, 5})
	ConfirmDelete(U32Slice{0, 1, 2, 3, 4, 5}, 1, U32Slice{0, 2, 3, 4, 5})
	ConfirmDelete(U32Slice{0, 1, 2, 3, 4, 5}, 2, U32Slice{0, 1, 3, 4, 5})
	ConfirmDelete(U32Slice{0, 1, 2, 3, 4, 5}, 3, U32Slice{0, 1, 2, 4, 5})
	ConfirmDelete(U32Slice{0, 1, 2, 3, 4, 5}, 4, U32Slice{0, 1, 2, 3, 5})
	ConfirmDelete(U32Slice{0, 1, 2, 3, 4, 5}, 5, U32Slice{0, 1, 2, 3, 4})
	ConfirmDelete(U32Slice{0, 1, 2, 3, 4, 5}, 6, U32Slice{0, 1, 2, 3, 4, 5})
}

func TestU32SliceDeleteIf(t *testing.T) {
	ConfirmDeleteIf := func(s U32Slice, f interface{}, r U32Slice) {
		if s.DeleteIf(f); !r.Equal(s) {
			t.Fatalf("DeleteIf(%v) should be %v but is %v", f, r, s)
		}
	}

	ConfirmDeleteIf(U32Slice{0, 1, 0, 3, 0, 5}, uint32(0), U32Slice{1, 3, 5})
	ConfirmDeleteIf(U32Slice{0, 1, 0, 3, 0, 5}, uint32(1), U32Slice{0, 0, 3, 0, 5})
	ConfirmDeleteIf(U32Slice{0, 1, 0, 3, 0, 5}, uint32(6), U32Slice{0, 1, 0, 3, 0, 5})

	ConfirmDeleteIf(U32Slice{0, 1, 0, 3, 0, 5}, func(x interface{}) bool { return x == uint32(0) }, U32Slice{1, 3, 5})
	ConfirmDeleteIf(U32Slice{0, 1, 0, 3, 0, 5}, func(x interface{}) bool { return x == uint32(1) }, U32Slice{0, 0, 3, 0, 5})
	ConfirmDeleteIf(U32Slice{0, 1, 0, 3, 0, 5}, func(x interface{}) bool { return x == uint32(6) }, U32Slice{0, 1, 0, 3, 0, 5})

	ConfirmDeleteIf(U32Slice{0, 1, 0, 3, 0, 5}, func(x uint32) bool { return x == uint32(0) }, U32Slice{1, 3, 5})
	ConfirmDeleteIf(U32Slice{0, 1, 0, 3, 0, 5}, func(x uint32) bool { return x == uint32(1) }, U32Slice{0, 0, 3, 0, 5})
	ConfirmDeleteIf(U32Slice{0, 1, 0, 3, 0, 5}, func(x uint32) bool { return x == uint32(6) }, U32Slice{0, 1, 0, 3, 0, 5})
}

func TestU32SliceEach(t *testing.T) {
	var count	uint32
	U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Each(func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Each(func(index int, i interface{}) {
		if i != uint32(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})

	U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Each(func(key, i interface{}) {
		if i != uint32(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})

	count = 0
	U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Each(func(i uint32) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Each(func(index int, i uint32) {
		if i != uint32(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})

	U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.Each(func(key interface{}, i uint32) {
		if i != uint32(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestU32SliceWhile(t *testing.T) {
	ConfirmLimit := func(s U32Slice, l int, f interface{}) {
		if count := s.While(f); count != l {
			t.Fatalf("%v.While() should have iterated %v times not %v times", s, l, count)
		}
	}

	s := U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	count := 0
	limit := 5
	ConfirmLimit(s, limit, func(i interface{}) bool {
		if count == limit {
			return false
		}
		count++
		return true
	})

	ConfirmLimit(s, limit, func(index int, i interface{}) bool {
		return index != limit
	})

	ConfirmLimit(s, limit, func(key, i interface{}) bool {
		return key.(int) != limit
	})

	count = 0
	ConfirmLimit(s, limit, func(i uint32) bool {
		if count == limit {
			return false
		}
		count++
		return true
	})

	ConfirmLimit(s, limit, func(index int, i uint32) bool {
		return index != limit
	})

	ConfirmLimit(s, limit, func(key interface{}, i uint32) bool {
		return key.(int) != limit
	})
}

func TestU32SliceUntil(t *testing.T) {
	ConfirmLimit := func(s U32Slice, l int, f interface{}) {
		if count := s.Until(f); count != l {
			t.Fatalf("%v.Until() should have iterated %v times not %v times", s, l, count)
		}
	}

	s := U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	count := 0
	limit := 5
	ConfirmLimit(s, limit, func(i interface{}) bool {
		if count == limit {
			return true
		}
		count++
		return false
	})

	ConfirmLimit(s, limit, func(index int, i interface{}) bool {
		return index == limit
	})

	ConfirmLimit(s, limit, func(key, i interface{}) bool {
		return key.(int) == limit
	})

	count = 0
	ConfirmLimit(s, limit, func(i uint32) bool {
		if count == limit {
			return true
		}
		count++
		return false
	})

	ConfirmLimit(s, limit, func(index int, i uint32) bool {
		return index == limit
	})

	ConfirmLimit(s, limit, func(key interface{}, i uint32) bool {
		return key.(int) == limit
	})
}

func TestU32SliceBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(s U32Slice, destination, source, count int, r U32Slice) {
		s.BlockCopy(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("BlockCopy(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlockCopy(U32Slice{}, 0, 0, 1, U32Slice{})
	ConfirmBlockCopy(U32Slice{}, 1, 0, 1, U32Slice{})
	ConfirmBlockCopy(U32Slice{}, 0, 1, 1, U32Slice{})

	ConfirmBlockCopy(U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 0, 4, U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	ConfirmBlockCopy(U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, 9, 4, U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	ConfirmBlockCopy(U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, 0, 4, U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 0})
	ConfirmBlockCopy(U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, 0, 4, U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	ConfirmBlockCopy(U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, 10, 4, U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	ConfirmBlockCopy(U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 2, 4, U32Slice{0, 1, 2, 3, 4, 2, 3, 4, 5, 9})
	ConfirmBlockCopy(U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 2, 5, 4, U32Slice{0, 1, 5, 6, 7, 8, 6, 7, 8, 9})
}

func TestU32SliceBlockClear(t *testing.T) {
	ConfirmBlockClear := func(s U32Slice, start, count int, r U32Slice) {
		s.BlockClear(start, count)
		if !r.Equal(s) {
			t.Fatalf("BlockClear(%v, %v) should be %v but is %v", start, count, r, s)
		}
	}

	ConfirmBlockClear(U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 4, U32Slice{0, 0, 0, 0, 4, 5, 6, 7, 8, 9})
	ConfirmBlockClear(U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, 4, U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	ConfirmBlockClear(U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 4, U32Slice{0, 1, 2, 3, 4, 0, 0, 0, 0, 9})
}

func TestU32SliceOverwrite(t *testing.T) {
	ConfirmOverwrite := func(s U32Slice, offset int, v, r U32Slice) {
		s.Overwrite(offset, v)
		if !r.Equal(s) {
			t.Fatalf("Overwrite(%v, %v) should be %v but is %v", offset, v, r, s)
		}
	}

	ConfirmOverwrite(U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, U32Slice{10, 9, 8, 7}, U32Slice{10, 9, 8, 7, 4, 5, 6, 7, 8, 9})
	ConfirmOverwrite(U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, U32Slice{10, 9, 8, 7}, U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	ConfirmOverwrite(U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, U32Slice{11, 12, 13, 14}, U32Slice{0, 1, 2, 3, 4, 11, 12, 13, 14, 9})
}

func TestU32SliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s U32Slice, l, c int, r U32Slice) {
		o := s.String()
		el := l
		if el > c {
			el = c
		}
		switch s.Reallocate(l, c); {
		case s == nil:				t.Fatalf("%v.Reallocate(%v, %v) created a nil value for Slice", o, l, c)
		case s.Cap() != c:			t.Fatalf("%v.Reallocate(%v, %v) capacity should be %v but is %v", o, l, c, c, s.Cap())
		case s.Len() != el:			t.Fatalf("%v.Reallocate(%v, %v) length should be %v but is %v", o, l, c, el, s.Len())
		case !r.Equal(s):			t.Fatalf("%v.Reallocate(%v, %v) should be %v but is %v", o, l, c, r, s)
		}
	}

	ConfirmReallocate(U32Slice{}, 0, 10, make(U32Slice, 0, 10))
	ConfirmReallocate(U32Slice{0, 1, 2, 3, 4}, 3, 10, U32Slice{0, 1, 2})
	ConfirmReallocate(U32Slice{0, 1, 2, 3, 4}, 5, 10, U32Slice{0, 1, 2, 3, 4})
	ConfirmReallocate(U32Slice{0, 1, 2, 3, 4}, 10, 10, U32Slice{0, 1, 2, 3, 4, 0, 0, 0, 0, 0})
	ConfirmReallocate(U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 1, 5, U32Slice{0})
	ConfirmReallocate(U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 5, U32Slice{0, 1, 2, 3, 4})
	ConfirmReallocate(U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, 5, U32Slice{0, 1, 2, 3, 4})
}

func TestU32SliceExtend(t *testing.T) {
	ConfirmExtend := func(s U32Slice, n int, r U32Slice) {
		c := s.Cap()
		s.Extend(n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Extend(%v) len should be %v but is %v", n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Extend(%v) cap should be %v but is %v", n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Extend(%v) should be %v but is %v", n, r, s)
		}
	}

	ConfirmExtend(U32Slice{}, 1, U32Slice{0})
	ConfirmExtend(U32Slice{}, 2, U32Slice{0, 0})
}

func TestU32SliceExpand(t *testing.T) {
	ConfirmExpand := func(s U32Slice, i, n int, r U32Slice) {
		c := s.Cap()
		s.Expand(i, n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Expand(%v, %v) len should be %v but is %v", i, n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Expand(%v, %v) cap should be %v but is %v", i, n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Expand(%v, %v) should be %v but is %v", i, n, r, s)
		}
	}

	ConfirmExpand(U32Slice{}, -1, 1, U32Slice{0})
	ConfirmExpand(U32Slice{}, 0, 1, U32Slice{0})
	ConfirmExpand(U32Slice{}, 1, 1, U32Slice{0})
	ConfirmExpand(U32Slice{}, 0, 2, U32Slice{0, 0})

	ConfirmExpand(U32Slice{0, 1, 2}, -1, 2, U32Slice{0, 0, 0, 1, 2})
	ConfirmExpand(U32Slice{0, 1, 2}, 0, 2, U32Slice{0, 0, 0, 1, 2})
	ConfirmExpand(U32Slice{0, 1, 2}, 1, 2, U32Slice{0, 0, 0, 1, 2})
	ConfirmExpand(U32Slice{0, 1, 2}, 2, 2, U32Slice{0, 1, 0, 0, 2})
	ConfirmExpand(U32Slice{0, 1, 2}, 3, 2, U32Slice{0, 1, 2, 0, 0})
	ConfirmExpand(U32Slice{0, 1, 2}, 4, 2, U32Slice{0, 1, 2, 0, 0})
}

func TestU32SliceDepth(t *testing.T) {
	ConfirmDepth := func(s U32Slice, i int) {
		if x := s.Depth(); x != i {
			t.Fatalf("%v.Depth() should be %v but is %v", s, i, x)
		}
	}
	ConfirmDepth(U32Slice{0, 1}, 0)
}

func TestU32SliceReverse(t *testing.T) {
	sxp := U32Slice{1, 2, 3, 4, 5}
	rxp := U32Slice{5, 4, 3, 2, 1}
	sxp.Reverse()
	if !rxp.Equal(sxp) {
		t.Fatalf("Reversal failed: %v", sxp)
	}
}

func TestU32SliceAppend(t *testing.T) {
	ConfirmAppend := func(s U32Slice, v interface{}, r U32Slice) {
		s.Append(v)
		if !r.Equal(s) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppend(U32Slice{}, uint32(0), U32Slice{0})

	ConfirmAppend(U32Slice{}, U32Slice{0}, U32Slice{0})
	ConfirmAppend(U32Slice{}, U32Slice{0, 1}, U32Slice{0, 1})
	ConfirmAppend(U32Slice{0, 1, 2}, U32Slice{3, 4}, U32Slice{0, 1, 2, 3, 4})
}

func TestU32SlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s U32Slice, v interface{}, r U32Slice) {
		if s.Prepend(v); !r.Equal(s) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrepend(U32Slice{}, uint32(0), U32Slice{0})
	ConfirmPrepend(U32Slice{0}, uint32(1), U32Slice{1, 0})

	ConfirmPrepend(U32Slice{}, U32Slice{0}, U32Slice{0})
	ConfirmPrepend(U32Slice{}, U32Slice{0, 1}, U32Slice{0, 1})
	ConfirmPrepend(U32Slice{0, 1, 2}, U32Slice{3, 4}, U32Slice{3, 4, 0, 1, 2})
}

func TestU32SliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s U32Slice, count int, r U32Slice) {
		if x := s.Repeat(count); !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(U32Slice{}, 5, U32Slice{})
	ConfirmRepeat(U32Slice{0}, 1, U32Slice{0})
	ConfirmRepeat(U32Slice{0}, 2, U32Slice{0, 0})
	ConfirmRepeat(U32Slice{0}, 3, U32Slice{0, 0, 0})
	ConfirmRepeat(U32Slice{0}, 4, U32Slice{0, 0, 0, 0})
	ConfirmRepeat(U32Slice{0}, 5, U32Slice{0, 0, 0, 0, 0})
}

func TestU32SliceCar(t *testing.T) {
	ConfirmCar := func(s U32Slice, r uint32) {
		n := s.Car().(uint32)
		if ok := n == r; !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCar(U32Slice{1, 2, 3}, 1)
}

func TestU32SliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r U32Slice) {
		if n := s.Cdr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCdr(U32Slice{1, 2, 3}, U32Slice{2, 3})
}

func TestU32SliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s U32Slice, v interface{}, r U32Slice) {
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplaca(U32Slice{1, 2, 3, 4, 5}, uint32(0), U32Slice{0, 2, 3, 4, 5})
}

func TestU32SliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s U32Slice, v interface{}, r U32Slice) {
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplacd(U32Slice{1, 2, 3, 4, 5}, nil, U32Slice{1})
	ConfirmRplacd(U32Slice{1, 2, 3, 4, 5}, uint32(10), U32Slice{1, 10})
	ConfirmRplacd(U32Slice{1, 2, 3, 4, 5}, U32Slice{5, 4, 3, 2}, U32Slice{1, 5, 4, 3, 2})
	ConfirmRplacd(U32Slice{1, 2, 3, 4, 5, 6}, U32Slice{2, 4, 8, 16}, U32Slice{1, 2, 4, 8, 16})
}

func TestU32SliceSetIntersection(t *testing.T) {
	ConfirmSetIntersection := func(s, o, r U32Slice) {
		x := s.SetIntersection(o)
		Sort(x)
		if !r.Equal(x) {
			t.Fatalf("%v.SetIntersection(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetIntersection(U32Slice{1, 2, 3}, U32Slice{}, U32Slice{})
	ConfirmSetIntersection(U32Slice{1, 2, 3}, U32Slice{1}, U32Slice{1})
	ConfirmSetIntersection(U32Slice{1, 2, 3}, U32Slice{1, 1}, U32Slice{1})
	ConfirmSetIntersection(U32Slice{1, 2, 3}, U32Slice{1, 2, 1}, U32Slice{1, 2})
}

func TestU32SliceSetUnion(t *testing.T) {
	ConfirmSetUnion := func(s, o, r U32Slice) {
		x := s.SetUnion(o)
		Sort(x)
		if !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(U32Slice{1, 2, 3}, U32Slice{}, U32Slice{1, 2, 3})
	ConfirmSetUnion(U32Slice{1, 2, 3}, U32Slice{1}, U32Slice{1, 2, 3})
	ConfirmSetUnion(U32Slice{1, 2, 3}, U32Slice{1, 1}, U32Slice{1, 2, 3})
	ConfirmSetUnion(U32Slice{1, 2, 3}, U32Slice{1, 2, 1}, U32Slice{1, 2, 3})
}

func TestU32SliceSetDifference(t *testing.T) {
	ConfirmSetUnion := func(s, o, r U32Slice) {
		x := s.SetDifference(o)
		Sort(x)
		if !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(U32Slice{1, 2, 3}, U32Slice{}, U32Slice{1, 2, 3})
	ConfirmSetUnion(U32Slice{1, 2, 3}, U32Slice{1}, U32Slice{2, 3})
	ConfirmSetUnion(U32Slice{1, 2, 3}, U32Slice{1, 1}, U32Slice{2, 3})
	ConfirmSetUnion(U32Slice{1, 2, 3}, U32Slice{1, 2, 1}, U32Slice{3})
}

func TestU32SliceFind(t *testing.T) {
	ConfirmFind := func(s U32Slice, v uint32, i int) {
		if x, ok := s.Find(v); !ok || x != i {
			t.Fatalf("%v.Find(%v) should be %v but is %v", s, v, i, x)
		}
	}

	ConfirmFind(U32Slice{0, 1, 2, 3, 4}, 0, 0)
	ConfirmFind(U32Slice{0, 1, 2, 3, 4}, 1, 1)
	ConfirmFind(U32Slice{0, 1, 2, 4, 3}, 2, 2)
	ConfirmFind(U32Slice{0, 1, 2, 4, 3}, 3, 4)
	ConfirmFind(U32Slice{0, 1, 2, 4, 3}, 4, 3)
}

func TestU32SliceFindN(t *testing.T) {
	ConfirmFindN := func(s U32Slice, v uint32, n int, i ISlice) {
		if x := s.FindN(v, n); !x.Equal(i) {
			t.Fatalf("%v.Find(%v, %v) should be %v but is %v", s, v, n, i, x)
		}
	}

	ConfirmFindN(U32Slice{1, 0, 1, 0, 1}, 2, 3, ISlice{})
	ConfirmFindN(U32Slice{1, 0, 1, 0, 1}, 1, 0, ISlice{0, 2, 4})
	ConfirmFindN(U32Slice{1, 0, 1, 0, 1}, 1, 1, ISlice{0})
	ConfirmFindN(U32Slice{1, 0, 1, 0, 1}, 1, 2, ISlice{0, 2})
	ConfirmFindN(U32Slice{1, 0, 1, 0, 1}, 1, 3, ISlice{0, 2, 4})
	ConfirmFindN(U32Slice{1, 0, 1, 0, 1}, 1, 4, ISlice{0, 2, 4})
}

func TestU32SliceKeepIf(t *testing.T) {
	ConfirmKeepIf := func(s U32Slice, f interface{}, r U32Slice) {
		if s.KeepIf(f); !r.Equal(s) {
			t.Fatalf("KeepIf(%v) should be %v but is %v", f, r, s)
		}
	}

	ConfirmKeepIf(U32Slice{0, 1, 0, 3, 0, 5}, uint32(0), U32Slice{0, 0, 0})
	ConfirmKeepIf(U32Slice{0, 1, 0, 3, 0, 5}, uint32(1), U32Slice{1})
	ConfirmKeepIf(U32Slice{0, 1, 0, 3, 0, 5}, uint32(6), U32Slice{})

	ConfirmKeepIf(U32Slice{0, 1, 0, 3, 0, 5}, func(x interface{}) bool { return x == uint32(0) }, U32Slice{0, 0, 0})
	ConfirmKeepIf(U32Slice{0, 1, 0, 3, 0, 5}, func(x interface{}) bool { return x == uint32(1) }, U32Slice{1})
	ConfirmKeepIf(U32Slice{0, 1, 0, 3, 0, 5}, func(x interface{}) bool { return x == uint32(6) }, U32Slice{})

	ConfirmKeepIf(U32Slice{0, 1, 0, 3, 0, 5}, func(x uint32) bool { return x == uint32(0) }, U32Slice{0, 0, 0})
	ConfirmKeepIf(U32Slice{0, 1, 0, 3, 0, 5}, func(x uint32) bool { return x == uint32(1) }, U32Slice{1})
	ConfirmKeepIf(U32Slice{0, 1, 0, 3, 0, 5}, func(x uint32) bool { return x == uint32(6) }, U32Slice{})
}

func TestU32SliceReverseEach(t *testing.T) {
	var count	uint32
	count = 9
	U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.ReverseEach(func(i interface{}) {
		if i != count {
			t.Fatalf("0: element %v erroneously reported as %v", count, i)
		}
		count--
	})

	U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.ReverseEach(func(index int, i interface{}) {
		if index != int(i.(uint32)) {
			t.Fatalf("1: element %v erroneously reported as %v", index, i)
		}
	})

	U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.ReverseEach(func(key, i interface{}) {
		if uint32(key.(int)) != i {
			t.Fatalf("2: element %v erroneously reported as %v", key, i)
		}
	})

	count = 9
	U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.ReverseEach(func(i uint32) {
		if i != count {
			t.Fatalf("3: element %v erroneously reported as %v", count, i)
		}
		count--
	})

	U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.ReverseEach(func(index int, i uint32) {
		if int(i) != index {
			t.Fatalf("4: element %v erroneously reported as %v", index, i)
		}
	})

	U32Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.ReverseEach(func(key interface{}, i uint32) {
		if key.(int) != int(i) {
			t.Fatalf("5: element %v erroneously reported as %v", key, i)
		}
	})
}

func TestU32SliceReplaceIf(t *testing.T) {
	ConfirmReplaceIf := func(s U32Slice, f, v interface{}, r U32Slice) {
		if s.ReplaceIf(f, v); !r.Equal(s) {
			t.Fatalf("ReplaceIf(%v, %v) should be %v but is %v", f, v, r, s)
		}
	}

	ConfirmReplaceIf(U32Slice{0, 1, 0, 3, 0, 5}, uint32(0), uint32(1), U32Slice{1, 1, 1, 3, 1, 5})
	ConfirmReplaceIf(U32Slice{0, 1, 0, 3, 0, 5}, uint32(1), uint32(0), U32Slice{0, 0, 0, 3, 0, 5})
	ConfirmReplaceIf(U32Slice{0, 1, 0, 3, 0, 5}, uint32(6), uint32(0), U32Slice{0, 1, 0, 3, 0, 5})

	ConfirmReplaceIf(U32Slice{0, 1, 0, 3, 0, 5}, func(x interface{}) bool { return x == uint32(0) }, uint32(1), U32Slice{1, 1, 1, 3, 1, 5})
	ConfirmReplaceIf(U32Slice{0, 1, 0, 3, 0, 5}, func(x interface{}) bool { return x == uint32(1) }, uint32(0), U32Slice{0, 0, 0, 3, 0, 5})
	ConfirmReplaceIf(U32Slice{0, 1, 0, 3, 0, 5}, func(x interface{}) bool { return x == uint32(6) }, uint32(0), U32Slice{0, 1, 0, 3, 0, 5})

	ConfirmReplaceIf(U32Slice{0, 1, 0, 3, 0, 5}, func(x uint32) bool { return x == uint32(0) }, uint32(1), U32Slice{1, 1, 1, 3, 1, 5})
	ConfirmReplaceIf(U32Slice{0, 1, 0, 3, 0, 5}, func(x uint32) bool { return x == uint32(1) }, uint32(0), U32Slice{0, 0, 0, 3, 0, 5})
	ConfirmReplaceIf(U32Slice{0, 1, 0, 3, 0, 5}, func(x uint32) bool { return x == uint32(6) }, uint32(0), U32Slice{0, 1, 0, 3, 0, 5})
}

func TestU32SliceReplace(t *testing.T) {
	ConfirmReplace := func(s U32Slice, v interface{}) {
		if s.Replace(v); !s.Equal(v) {
			t.Fatalf("Replace() should be %v but is %v", s, v)
		}
	}

	ConfirmReplace(U32Slice{0, 1, 2, 3, 4, 5}, U32Slice{9, 8, 7, 6, 5})
	ConfirmReplace(U32Slice{0, 1, 2, 3, 4, 5}, []uint32{9, 8, 7, 6, 5})
}

func TestU32SliceSelect(t *testing.T) {
	ConfirmSelect := func(s U32Slice, f interface{}, r U32Slice) {
		if x := s.Select(f); !r.Equal(x) {
			t.Fatalf("Select(%v) should be %v but is %v", f, r, s)
		}
	}

	ConfirmSelect(U32Slice{0, 1, 0, 3, 0, 5}, uint32(0), U32Slice{0, 0, 0})
	ConfirmSelect(U32Slice{0, 1, 0, 3, 0, 5}, uint32(1), U32Slice{1})
	ConfirmSelect(U32Slice{0, 1, 0, 3, 0, 5}, uint32(6), U32Slice{})

	ConfirmSelect(U32Slice{0, 1, 0, 3, 0, 5}, func(x interface{}) bool { return x == uint32(0) }, U32Slice{0, 0, 0})
	ConfirmSelect(U32Slice{0, 1, 0, 3, 0, 5}, func(x interface{}) bool { return x == uint32(1) }, U32Slice{1})
	ConfirmSelect(U32Slice{0, 1, 0, 3, 0, 5}, func(x interface{}) bool { return x == uint32(6) }, U32Slice{})

	ConfirmSelect(U32Slice{0, 1, 0, 3, 0, 5}, func(x uint32) bool { return x == uint32(0) }, U32Slice{0, 0, 0})
	ConfirmSelect(U32Slice{0, 1, 0, 3, 0, 5}, func(x uint32) bool { return x == uint32(1) }, U32Slice{1})
	ConfirmSelect(U32Slice{0, 1, 0, 3, 0, 5}, func(x uint32) bool { return x == uint32(6) }, U32Slice{})
}

func TestU32SliceUniq(t *testing.T) {
	ConfirmUniq := func(s, r U32Slice) {
		if s.Uniq(); !r.Equal(s) {
			t.Fatalf("Uniq() should be %v but is %v", r, s)
		}
	}

	ConfirmUniq(U32Slice{0, 0, 0, 0, 0, 0}, U32Slice{0})
	ConfirmUniq(U32Slice{0, 1, 0, 3, 0, 5}, U32Slice{0, 1, 3, 5})
}

func TestU32SliceValuesAt(t *testing.T) {
	ConfirmValuesAt := func(s U32Slice, i []int, r U32Slice) {
		if x := s.ValuesAt(i...); !r.Equal(x) {
			t.Fatalf("%v.ValuesAt(%v) should be %v but is %v", s, i, r, x)
		}
	}

	ConfirmValuesAt(U32Slice{0, 1, 2, 3, 4, 5}, []int{}, U32Slice{})
	ConfirmValuesAt(U32Slice{0, 1, 2, 3, 4, 5}, []int{ 0, 1 }, U32Slice{0, 1})
	ConfirmValuesAt(U32Slice{0, 1, 2, 3, 4, 5}, []int{ 0, 3 }, U32Slice{0, 3})
	ConfirmValuesAt(U32Slice{0, 1, 2, 3, 4, 5}, []int{ 0, 3, 4, 3 }, U32Slice{0, 3, 4, 3})
}

func TestU32SliceInsert(t *testing.T) {
	ConfirmInsert := func(s U32Slice, n int, v interface{}, r U32Slice) {
		if s.Insert(n, v); !r.Equal(s) {
			t.Fatalf("Insert(%v, %v) should be %v but is %v", n, v, r, s)
		}
	}

	ConfirmInsert(U32Slice{}, 0, uint32(0), U32Slice{0})
	ConfirmInsert(U32Slice{}, 0, U32Slice{0}, U32Slice{0})
	ConfirmInsert(U32Slice{}, 0, U32Slice{0, 1}, U32Slice{0, 1})

	ConfirmInsert(U32Slice{0}, 0, uint32(1), U32Slice{1, 0})
	ConfirmInsert(U32Slice{0}, 0, U32Slice{1}, U32Slice{1, 0})
	ConfirmInsert(U32Slice{0}, 1, uint32(1), U32Slice{0, 1})
	ConfirmInsert(U32Slice{0}, 1, U32Slice{1}, U32Slice{0, 1})

	ConfirmInsert(U32Slice{0, 1, 2}, 0, uint32(3), U32Slice{3, 0, 1, 2})
	ConfirmInsert(U32Slice{0, 1, 2}, 1, uint32(3), U32Slice{0, 3, 1, 2})
	ConfirmInsert(U32Slice{0, 1, 2}, 2, uint32(3), U32Slice{0, 1, 3, 2})
	ConfirmInsert(U32Slice{0, 1, 2}, 3, uint32(3), U32Slice{0, 1, 2, 3})

	ConfirmInsert(U32Slice{0, 1, 2}, 0, U32Slice{3, 4}, U32Slice{3, 4, 0, 1, 2})
	ConfirmInsert(U32Slice{0, 1, 2}, 1, U32Slice{3, 4}, U32Slice{0, 3, 4, 1, 2})
	ConfirmInsert(U32Slice{0, 1, 2}, 2, U32Slice{3, 4}, U32Slice{0, 1, 3, 4, 2})
	ConfirmInsert(U32Slice{0, 1, 2}, 3, U32Slice{3, 4}, U32Slice{0, 1, 2, 3, 4})
}