package set

import (
	"testing"
)

func Test(t *testing.T) {
	s := New()

	s.Insert(5)

	if s.Len() != 1 {
		t.Errorf("Length should be 1")
	}

	s.Remove(5)
	if s.Len() != 0 {
		t.Errorf("Length should be 0")
	}

	x := New()
	s.Insert(x)

	if s.Has(5) {
		t.Errorf("The set should be empty")
	}

	// Difference
	s1 := New(1, 2, 3, 4, 5, 6)
	s2 := New(4, 5, 6)
	s3 := s1.Difference(s2)
	if s3.Len() != 3 {
		t.Errorf("Length should be 3")
	}

	if !(s3.Has(1) && s3.Has(2) && s3.Has(3)) {
		t.Errorf("Set should only contain 1, 2, 3")
	}

	// Intersection
	s3 = s1.Intersection(s2)
	if s3.Len() != 3 {
		t.Errorf("Length should be 3 after intersection")
	}

	if !(s3.Has(4) && s3.Has(5) && s3.Has(6)) {
		t.Errorf("Set should contain 4, 5, 6")
	}

	// Union
	s4 := New(7, 8, 9)
	s3 = s2.Union(s4)

	if s3.Len() != 6 {
		t.Errorf("Length should be 6 after union")
	}

	if !(s3.Has(7)) {
		t.Errorf("Set should contain 4, 5, 6, 7, 8, 9")
	}

	// Subset
	if !s1.Has(s1) {
		t.Errorf("set should be a subset of itself")
	}

	// Hierarchical
	alice := New()
	bob := New()
	jane := New()

	owners := New(alice)
	writers := New(bob)
	readers := New(jane)

	writers.Insert(owners)
	readers.Insert(writers)
	writers.Insert("3")

	if !owners.Check(alice) {
		t.Errorf("alice should be an owner")
	}
	if owners.Check(bob) {
		t.Errorf("bob should not be an owner")
	}
	if owners.Check(jane) {
		t.Errorf("jane should not be an owner")
	}
	if owners.Check("3") {
		t.Errorf("3 should not be a owner")
	}

	if !writers.Check(alice) {
		t.Errorf("alice should be an writer")
	}
	if !writers.Check(bob) {
		t.Errorf("bob should be a writers")
	}
	if writers.Check(jane) {
		t.Errorf("jane should not be a writer")
	}
	if !writers.Check("3") {
		t.Errorf("3 should be a reader")
	}

	if !readers.Check(alice) {
		t.Errorf("alice should be an reader")
	}
	if !readers.Check(bob) {
		t.Errorf("bob should be a reader")
	}
	if !readers.Check(jane) {
		t.Errorf("jane should be a reader")
	}
	if !readers.Check("3") {
		t.Errorf("3 should be a reader")
	}
}
