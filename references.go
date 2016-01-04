package gobible

import (
    "strconv"
)

// The maximum number of verses in a chapter in the Bible, plus one.
const MaxVerses = 177

// The maximum number of chapters in the Bible, plus one.
const MaxChapters = 151



func MakeReference(b, c, v int) Reference {
    return newReference(b, c, v)
}


// Specifies a single verse, chapter or book reference, e.g. John 3:16 or Genesis 1
type Reference int

// Specifies a group of verses, e.g. Mark 16:15-16. Can be same passage.
type Passage struct {
	Begin Reference
	End Reference
}


// Every book of the Bible in order, as an array.
var englishBooksList = []string{
    "Genesis",
    "Exodus",
    "Leviticus",
    "Numbers",
    "Deuteronomy",
    "Joshua",
    "Judges",
    "Ruth",
    "I Samuel",
    "II Samuel",
    "I Kings",
    "II Kings",
    "I Chronicles",
    "II Chronicles",
    "Ezra",
    "Nehemiah",
    "Esther",
    "Job",
    "Psalms",
    "Proverbs",
    "Ecclesiastes",
    "Song of Solomon",
    "Isaiah",
    "Jeremiah",
    "Lamentations",
    "Ezekiel",
    "Daniel",
    "Hosea",
    "Joel",
    "Amos",
    "Obadiah",
    "Jonah",
    "Micah",
    "Nahum",
    "Habakkuk",
    "Zephaniah",
    "Haggai",
    "Zechariah",
    "Malachi",
    "Matthew",
    "Mark",
    "Luke",
    "John",
    "Acts",
    "Romans",
    "I Corinthians",
    "II Corinthians",
    "Galatians",
    "Ephesians",
    "Philippians",
    "Colossians",
    "I Thessalonians",
    "II Thessalonians",
    "I Timothy",
    "II Timothy",
    "Titus",
    "Philemon",
    "Hebrews",
    "James",
    "I Peter",
    "II Peter",
    "I John",
    "II John",
    "III John",
    "Jude",
    "Revelation",
}



func SingleVersePassage(ref Reference) Passage {
    return Passage{ref, ref}
}

func (ref Reference) divide() (book, chapter, verse int) {
    code := int(ref)
    verse = code % MaxVerses
    code = code / MaxVerses
    chapter = code % MaxChapters
    book = code / MaxChapters
    return
}

func (ref Reference) book() string {
    n, _, _ := ref.divide()
    
    return englishBooksList[n - 1]
}

func (ref Reference) chapter() int {
    _, n, _ := ref.divide()
    return n
}

func (ref Reference) verse() int {
    _, _, n := ref.divide()
    return n
}



// Builds a Bible reference from individual numbers. You shouldn't need this.
func newReference(book, chapter, verse int) Reference {
    k := verse
    k = (chapter * MaxVerses) + k
    k = (book * MaxChapters * MaxVerses) + k
    return Reference(k)
}


// Get the full title of a reference, e.g. "John 3:16"
func (ref Reference) Title() string {
    return ref.book() + " " +
        strconv.Itoa(ref.chapter()) +
        ":" + strconv.Itoa(ref.verse())
}


func (r Passage) Title() string {
	if r.Begin  == r.End {
		return r.Begin.Title()
	}
	
	if r.Begin.chapter() == r.End.chapter() {
		return r.Begin.Title() + "-" + strconv.Itoa(r.End.verse())
	}
	
	if r.Begin.book()== r.End.book(){
		return r.Begin.Title() + "-" +  strconv.Itoa(r.End.chapter())+":" +strconv.Itoa(r.End.verse())
	}
	
	return r.Begin.Title() + " - " + r.End.Title()
}