package set

// Set represents set structure
type Set struct {
	data []string
}

// New constructs new set
func New() *Set {
	return &Set{}
}

// Size resurns length of set
func (s *Set) Size() int {
	return len(s.data)
}

func (s *Set) hasElement(element string) (bool, int) {
	for index, el := range s.data {
		if el == element {
			return true, index
		}
	}

	return false, 0
}

// Values returns set values
func (s *Set) Values() []string {
	return s.data
}

// Add adds element to set, returns true if element added successfully
func (s *Set) Add(element string) bool {
	hasElement, _ := s.hasElement(element)
	if !hasElement {
		s.data = append(s.data, element)

		return true
	}

	return false
}

// Remove removes element from set
func (s *Set) Remove(element string) bool {
	hasElement, i := s.hasElement(element)
	if hasElement {
		copy(s.data[i:], s.data[i+1:])
		s.data[len(s.data)-1] = ""
		s.data = s.data[:len(s.data)-1]

		return true
	}

	return false
}

// Union returns unioned set as new set
func (s *Set) Union(set *Set) *Set {
	unionSet := New()

	firstSet := s.data
	secondSet := set.Values()

	for _, el := range firstSet {
		unionSet.Add(el)
	}
	for _, el := range secondSet {
		unionSet.Add(el)
	}

	return unionSet
}

// Intersect returns intersection set as new set
func (s *Set) Intersect(set *Set) *Set {
	interSectionSet := New()
	otherSet := set.Values()

	for _, el := range otherSet {
		hasElement, _ := s.hasElement(el)

		if hasElement {
			interSectionSet.Add(el)
		}
	}

	return interSectionSet
}

// Difference returns difference set as new set
func (s *Set) Difference(set *Set) *Set {
	differenceSet := New()
	otherSet := set.Values()

	for _, el := range otherSet {
		hasElement, _ := s.hasElement(el)

		if !hasElement {
			differenceSet.Add(el)
		}
	}

	return differenceSet
}

// Subset returns true if provided set contains all elements in current set
func (s *Set) Subset(set *Set) bool {
	otherSet := set.Values()

	for _, el := range otherSet {
		hasElement, _ := s.hasElement(el)
		if !hasElement {
			return false
		}
	}

	return true
}
