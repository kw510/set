package set

type parents map[any]*Set
type Set struct {
	h       map[any]any
	parents parents
}

type empty struct{}

// Create a new set
func New(values ...any) *Set {
	s := &Set{
		h:       map[any]any{},
		parents: map[any]*Set{},
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

func (s *Set) IsMember(value any) bool {
	switch v := value.(type) {
	case *Set:
		if s == v {
			return true
		}
	}

	return false
}

// DFS up the hierarchy. Can only be sets up.
func (s *Set) SearchParents(target *Set) bool {
	// Base case.
	if s == target {
		return true
	}

	// N + 1
	for _, parent := range s.parents {
		if parent.SearchParents(target) {
			return true
		}
	}

	return false
}

// DFS down the hierarchy. Supports any value.
func (s *Set) SearchMembers(target any) bool {
	// Base case.
	if _, ok := s.h[target]; ok {
		return true
	}

	// N + 1
	for _, member := range s.h {
		if memberSet, ok := member.(*Set); ok {
			if memberSet.SearchMembers(target) {
				return true
			}
		}
	}

	return false
}

// Search for the element/set is in the set.
func (s *Set) Check(value any) bool {
	switch v := value.(type) {
	case *Set:
		// If a set, we can check if the parents contain the value.
		return v.SearchParents(s)
	default:
		// If unknown, we have to check members.
		return s.SearchMembers(v)
	}
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
		_, ok := v.parents[s]
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
			v.parents[s] = s
			s.h[v] = v
		default:
			s.h[v] = empty{}
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
