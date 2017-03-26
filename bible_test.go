package gobible

import (
	"testing"
)

func TestReferences(t *testing.T) {

	type ReferenceTest struct {
		reference   Reference
		expectation string
	}

	test := []ReferenceTest{
		{MakeReference(1, 1, 1), "Genesis 1:1"},
		{MakeReference(40, 1, 1), "Matthew 1:1"},
		{MakeReference(40, 20, 1), "Matthew 20:1"},
		{MakeReference(1, 1, 10), "Genesis 1:10"},
	}

	for _, m := range test {
		title := m.reference.Title()
		if title != m.expectation {
			t.Errorf("expected %s, got %s\n", m.expectation, title)
		}
		title, err := SingleVersePassage(m.reference).Title()
		if err != nil {
			t.Errorf(err.Error())
		}
		if title != m.expectation {
			t.Errorf("expected %s, got %s; single verse\n", m.expectation, title)
		}
	}
}

/*
func TestBasic(t *testing.T) {
    bible := Bible(nil)

    err := bible.RegisterProvider(esv.EsvApiProvider())

    if err != nil {
        t.Error(err)
        return
    }

    passage := MeultiVersePassage(MakeReference(43, 3, 15), 3)

    text, err := bible.Text(esv.EnglishStandardVersion, passage)
    if err != nil {
        t.Error(err)
        return
    }

    if !strings.Contains(text, "For God so loved the world") {
        t.Error("John 3:16 pull request failed.")
        return
    }


}
*/
