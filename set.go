package set

import (
	"github.com/rs/xid"
)

type Set struct {
	key    string
	h      map[any]Set
	parent map[string]*Set
}

// Create a new set
func New(values ...any) *Set {
	s := &Set{
		key:    xid.New().String(),
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

// Test to see whether or not the element/set is in the set
func (s Set) Has(value any) bool {
	switch v := value.(type) {
	case *Set:
		// Check if its itself
		if s.key == v.key {
			return true
		}
		// Check if its a direct descendent
		p, ok := v.parent[s.key]
		if ok && p.key == s.key {
			return true
		}
		// Check chain - Depth First Search (Inverse)
		// Maybe make this optional, or add a new search func
		for _, p := range v.parent {
			if p.Has(value) {
				return true
			}
		}

	default:
		if _, ok := s.h[v]; ok {
			return true
		}
	}
	return false
}

// Add an element/set to the set
func (s Set) Insert(values ...any) {
	for _, i := range values {
		switch v := i.(type) {
		case *Set:
			v.parent[s.key] = &s
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
