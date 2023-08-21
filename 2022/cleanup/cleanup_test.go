package cleanup

import (
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
