package gobible

import (
	"testing"
)

func TestParsing(t *testing.T) {

	type ReferenceTest struct {
		input       string
		expectation string
	}

	test := []ReferenceTest{
		{"gen 1:1", "Genesis 1:1"},
		{"Mat. 1:1", "Matthew 1:1"},
		{"Matt. 1:10", "Matthew 1:10"},
		{"Matthew 20:1!", "Matthew 20:1"},
		{"Matthew 20:1-22:3", "Matthew 20:1-22:3"},
		{"Matthew 20:1-22:3", "Matthew 20:1-22:3"},
		{"Matthew 20:1-", "Matthew 20:1"},
		{"Hebrews", "Hebrews"},
		{"I read Hebrews", "Hebrews"},
		{"I read Hebrews again", "Hebrews"},
		{"Matthew", "Matthew"},
		{"Matthew 1:2-20:1", "Matthew 1:2-20:1"},

		{"1 Corinthians 1:10", "I Corinthians 1:10"},
		//{"I got a job reading 1 Peter 2 to the romans", "I Peter 2"},
		//{"I got a job 2 reading to the romans", "Job 2"},
		{"I read jude 1", "Jude 1"},
		{"I like 1 Corinthians 1:10 it's cool", "I Corinthians 1:10"},
	}

	for _, m := range test {
		results := ParseReferencesFromText(m.input)
		if len(results) != 1 {
			t.Errorf("expected 1 output, got %d for '%s'\n", len(results), m.input)
			continue
		}
		reference := results[0]

		title, err := reference.Title()
		if err != nil {
			t.Errorf(err.Error())
		}
		if title != m.expectation {
			t.Errorf("expected %s, got %s\n", m.expectation, title)
		}
	}
}
