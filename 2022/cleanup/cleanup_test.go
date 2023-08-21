package cleanup

import (
	"strings"
	"testing"
)

func TestIsFullyOverlapping(t *testing.T) {
	// First section fully overlaps second section
	firstSection := Section{3, 6}
	secondSection := Section{0, 10}
	overlaps := isFullyOverlapping(firstSection, secondSection)
	if !overlaps {
		t.Fatalf(`First and second section must overlap`)
	}

	// Second section fully overlaps first section
	firstSection = Section{0, 10}
	secondSection = Section{4, 5}
	overlaps = isFullyOverlapping(firstSection, secondSection)
	if !overlaps {
		t.Fatalf(`First and second section must overlap`)
	}

	// First section partially overlaps second section
	firstSection = Section{7, 12}
	secondSection = Section{5, 10}
	overlaps = isFullyOverlapping(firstSection, secondSection)
	if overlaps {
		t.Fatalf(`First section partially overlaps second section`)
	}

	// Second section partially overlaps second section
	firstSection = Section{5, 10}
	secondSection = Section{7, 12}
	overlaps = isFullyOverlapping(firstSection, secondSection)
	if overlaps {
		t.Fatalf(`Second section partially overlaps first section`)
	}

	firstSection = Section{0, 4}
	secondSection = Section{5, 10}
	overlaps = isFullyOverlapping(firstSection, secondSection)
	if overlaps {
		t.Fatalf(`First section does not overlap second section`)
	}

	firstSection = Section{5, 10}
	secondSection = Section{0, 4}
	overlaps = isFullyOverlapping(firstSection, secondSection)
	if overlaps {
		t.Fatalf(`Second section does not overlap first section`)
	}
}

func TestIsPartiallyOverlapping(t *testing.T) {
	// First section partially overlaps second section
	firstSection := Section{5, 10}
	secondSection := Section{7, 12}
	partiallyOverlaps := isPartiallyOverlapping(firstSection, secondSection)
	if !partiallyOverlaps {
		t.Fatalf(`First section must partially overlap second section`)
	}

	// Second section partially overlaps second sectiono
	firstSection = Section{7, 12}
	secondSection = Section{5, 10}
	partiallyOverlaps = isPartiallyOverlapping(firstSection, secondSection)
	if !partiallyOverlaps {
		t.Fatalf(`Second section must partially overlap first section`)
	}

	// First and second section do not overlap
	firstSection = Section{0, 4}
	secondSection = Section{5, 10}
	partiallyOverlaps = isPartiallyOverlapping(firstSection, secondSection)
	if partiallyOverlaps {
		t.Fatalf(`First section does not partially overlap second section`)
	}

	// First and second section do not overlap
	firstSection = Section{5, 10}
	secondSection = Section{0, 4}
	partiallyOverlaps = isFullyOverlapping(firstSection, secondSection)
	if partiallyOverlaps {
		t.Fatalf(`Second section does not partially overlap first section`)
	}
}

func TestCountFullyOverlappingAssignments(t *testing.T) {
	s := `70-70,66-87
70-70,40-87
32-99,31-78
20-54,69-96`
	reader := strings.NewReader(s)
	count := CountFullOverlappingAssignments(reader)
	if count != 2 {
		t.Fatalf(`Number of overlaps is 2 but got %d`, count)
	}
}
func TestCountPartiallyOverlappingAssignments(t *testing.T) {
	s := `70-70,66-87
40-70,50-90
50-90,40-70
20-54,69-96
90-99,1-10`
	reader := strings.NewReader(s)
	count := CountPartialOverlappingAssignments(reader)
	if count != 3 {
		t.Fatalf(`Number of partial overlaps is 3 but got %d`, count)
	}
}
