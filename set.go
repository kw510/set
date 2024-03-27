package set

import (
	"fmt"
)

type Set struct {
	h      map[any]Set
	parent map[string]*Set
}

func (s *Set) key() string {
	return fmt.Sprintf("%p", s)
}

// Create a new set
func New(values ...any) *Set {
	s := &Set{
		h:      map[any]Set{},
		parent: map[string]*Set{},
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
		if s.key() == v.key() {
			return true
		}
		for _, p := range v.parent {
			_, ok := p.parent[s.key()]
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
		for _, c := range s.h {
			if _, ok := c.h[v]; ok {
				return true
			}
			if c.Check(v) {
				return true
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
		if s.key() == v.key() {
			return true
		}
		// Check if its a direct descendent
		_, ok := v.parent[s.key()]
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

// Add an element/set to the set
func (s *Set) Insert(values ...any) {
	for _, i := range values {
		switch v := i.(type) {
		case *Set:
			v.parent[s.key()] = s
			s.h[v] = *v
		default:
			s.h[v] = *New()
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
