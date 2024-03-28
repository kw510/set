package set

type Set struct {
	h      map[any]any
	parent map[any]*Set
}

// Create a new set
func New(values ...any) *Set {
	s := &Set{
		h:      map[any]any{},
		parent: map[any]*Set{},
	}

	s.Insert(values...)

	return s
}

// Find the difference between two sets
func (s *Set) Difference(set *Set) *Set {
	o := New()
	for k := range s.h {
		if !set.Has(k) {
			o.Insert(k)
		}
	}
	return o
}

// Search for the element/set is in the set
func (s *Set) Check(value any) bool {
	switch v := value.(type) {
	case *Set:
		// Check if its itself
		if s == v {
			return true
		}
		for _, p := range v.parent {
			_, ok := p.parent[s]
			if ok {
				return ok
			}
			if s.Check(p) {
				return true
			}
		}
	default:
		if _, ok := s.h[v]; ok {
			return true
		}

		for _, a := range s.h {
			if c, ok := a.(*Set); ok {
				if _, ok := c.h[v]; ok {
					return true
				}
				if c.Check(v) {
					return true
				}
			}
		}
	}
	return false
}

// Test to see whether or not the element/set is in the set
func (s *Set) Has(value any) bool {
	switch v := value.(type) {
	case *Set:
		// Check if its itself
		if s == v {
			return true
		}
		// Check if its a direct descendent
		_, ok := v.parent[s]
		if ok {
			return true
		}

	default:
		if _, ok := s.h[v]; ok {
			return true
		}
	}

	return false
}

// Add an element/set to the set.
func (s *Set) Insert(values ...any) {
	for _, i := range values {
		switch v := i.(type) {
		case *Set:
			if s.Check(v) { // force acyclic graph / uniqueness
				continue
			}
			v.parent[s] = s
			s.h[v] = v
		default:
			s.h[v] = v
		}
	}
}

// Find the intersection of two sets
func (s *Set) Intersection(set *Set) *Set {
	o := New()
	for k := range s.h {
		if set.Has(k) {
			o.Insert(k)
		}
	}
	return o
}

// Return the number of items in the set
func (s *Set) Len() int {
	return len(s.h)
}

// Remove a member from the set
func (s *Set) Remove(value any) {
	delete(s.h, value)
}

// Find the union of two sets
func (s *Set) Union(set *Set) *Set {
	o := New()
	for k := range s.h {
		o.Insert(k)
	}
	for k := range set.h {
		o.Insert(k)
	}
	return o
}

// Flattens the set, including subsets. Full graph traversal.
func (s *Set) Flatten() []any {
	o := []any{}
	for k := range s.h {
		switch v := k.(type) {
		case *Set:
			o = append(o, v.Flatten()...)
		default:
			o = append(o, k)
		}
	}

	return o
}
