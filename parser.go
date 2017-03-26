package gobible

import (
	"strconv"
	"strings"
)

var shortBookNames = map[string]int{
	"gen": 1, "exo": 2, "lev": 3, "num": 4, "deu": 5, "jos": 6, "jsh": 6, "jdg": 7,
	"rut": 8, "rth": 8, "1sa": 9, "1sm": 9, "2sa": 10, "2sm": 10, "1ki": 11, "1kg": 11, "2ki": 12, "2kg": 12,
	"1ch": 13, "2ch": 14, "ezr": 15, "neh": 16,
	"est": 17, "job": 18, "psa": 19, "psm": 19, "pss": 19, "pro": 20, "prv": 20, "ecc": 21, "qoh": 21, "son": 22, "sos": 22,
	"can": 22, "isa": 23, "jer": 24,
	"lam": 25, "eze": 26, "ezk": 26, "dan": 27, "hos": 28, "joe": 29, "amo": 30, "oba": 31, "jon": 32, "jnh": 32,
	"mic": 33, "nah": 34, "hab": 35, "zep": 36, "hag": 37, "zec": 38, "mal": 39,

	"matt": 40,
	"mat":  40, "mar": 41, "mrk": 41, "luk": 42, "joh": 43, "act": 44, "rom": 45, "1co": 46, "2co": 47,
	"gal": 48, "eph": 49, "phi": 50, "phl": 50, "php": 50, "col": 51, "1th": 52, "2th": 53, "1ti": 54,
	"2ti": 55, "tit": 56, "phm": 57, "heb": 58, "jam": 59, "jas": 59, "1pe": 60, "1pt": 60, "2pe": 61,
	"2pt": 61, "1jh": 62, "1jo": 62, "1jn": 62, "2jh": 63, "2jo": 63, "2jn": 63, "3jh": 64, "3jo": 64, "3jn": 64, "jud": 65, "rev": 66,

	"ge": 1, "gn": 1, "ex": 2, "le": 3, "lv": 3, "nu": 4, "nm": 4, "nb": 4, "dt": 5,
	"jg": 7, "ru": 8, "ne": 16, "es": 17, "jb": 18, "ps": 19, "pr": 20, "ec": 21, "so": 22,
	"is": 23, "je": 24, "jr": 24, "la": 25, "da": 27, "dn": 27, "jl": 29, "ob": 31, "na": 34, "zp": 36,
	"hg": 37, "zc": 38, "ml": 39, "mt": 40, "mk": 41, "mr": 41, "lk": 42, "jo": 43, "jn": 43, "ax": 44,
	"ac": 44, "ro": 45, "rm": 45, "ga": 48, "jm": 59, "js": 59, "re": 66,
}

var twoPart = map[string]int{
	"sam": 9,
	"kin": 11,
	"kgs": 11,
	"chr": 13, "cor": 46, "the": 52, "tim": 54, "pet": 60, "joh": 62,
	"sa": 9, "sm": 9, "ki": 11, "ch": 13, "co": 13, "jo": 62, "jn": 62, "th": 52, "ti": 54, "pe": 60, "pt": 60,
	"samuel": 9, "kings": 11, "chronicles": 13, "corinthians": 46, "thessalonians": 52, "timothy": 54, "peter": 60, "john": 62,
}

var wholeBookExceptions = []string{
	"job", "son", "can", "dan", "mat", "act", "ex",
	"na", "mt", "jo", "sam", "pt", "matt", "matthew", "solomon", "ps",
	"john", "peter", "tim", "timothy", "romans"}

var twoPartBooks = []int{9, 11, 13, 52, 54, 60, 61, 62}
var singleChapter = []int{31, 63, 64, 65}

// ParseReferencesFromText returns a list of passages detected in a given string.
func ParseReferencesFromText(text string) []Passage {
	var previous string
	segments := strings.Split(strings.ToLower(text), " ")
	var foundBook Reference
	bookFound := false
	results := []Passage{}
	segments = append(segments, " ")
	for _, segment := range segments {
		part := strings.Trim(segment, "!,./\\;:'\"-?")
		if len(part) == 0 {
			continue
		}

		if bookFound {
			passage := parseReferenceIntoPassage(part, foundBook)
			// Whole book exceptions list - we don't need it yet.
			//if !(foundBook == passage.Begin && foundBook == passage.End && isWholeBookExcluded(previous)) {
			results = append(results, passage)
			//}
			bookFound = false
		} else {
			bookIndex := checkBookName(part, previous)
			if bookIndex > 0 {
				foundBook = newReference(bookIndex, 0, 0)
				bookFound = true
			}
		}
		previous = segment
	}
	return results
}

func isWholeBookExcluded(part string) bool {
	for _, book := range wholeBookExceptions {
		if book == part {
			return true
		}
	}
	return false
}

func parseReferenceIntoPassage(part string, book Reference) Passage {
	if !strings.ContainsAny(part, "1234567890") {
		return SingleVersePassage(book)
	}
	// i:i-i:i
	partitions := strings.Split(part, "-")
	if len(partitions) > 2 {
		return SingleVersePassage(book)
	}

	if len(partitions) == 1 {
		return SingleVersePassage(parseReferencePart(partitions[0], book.bookIndex()))
	}

	// 2 length

	return Passage{
		Begin: parseReferencePart(partitions[0], book.bookIndex()),
		End:   parseReferencePart(partitions[1], book.bookIndex()),
	}

}

func parseReferencePart(part string, bookIndex int) Reference {
	partitions := strings.Split(part, ":")
	if len(partitions) > 2 {
		return MakeReference(bookIndex, 0, 0)
	}
	chp, err := strconv.Atoi(partitions[0])
	if err != nil {
		return MakeReference(bookIndex, 0, 0)
	}
	ver := 0
	if len(partitions) == 2 {
		// Don't support partial verses yet.
		ver, err = strconv.Atoi(partitions[1])
		if err != nil {
			return MakeReference(bookIndex, chp, 0)
		}
	}
	return MakeReference(bookIndex, chp, ver)
}

func checkBookName(part, previous string) int {
	if len(part) <= 1 {
		return -1
	}

	result := checkTwoPartBook(part, previous)
	if result > 0 {
		return result
	}

	if len(part) <= 4 {
		res := checkBookShortName(part)
		if res > 0 || len(part) <= 3 {
			return res
		}
	}
	return checkBookFullName(part)
}

func checkTwoPartBook(part, previous string) int {
	for bookName, bookIndex := range twoPart {
		if part != bookName {
			continue
		}
		// Check prefix
		prefix := convertBookPrefix(previous)
		if prefix < 0 {
			if bookIndex == 61 { // The John Exception
				return 32
			}
			return -1
		}
		return bookIndex + prefix - 1 // -1 because the prefix starts at 1
	}
	return -1
}

func checkBookShortName(part string) int {
	for refCheck, book := range shortBookNames {
		if refCheck != part {
			continue
		}
		return book
	}
	return -1
}

func checkBookFullName(part string) int {
	for index, bookName := range englishBooksList {
		bookDiv := strings.Split(bookName, " ")
		lastBookRef := strings.ToLower(bookDiv[len(bookDiv)-1])
		if part == lastBookRef {
			return index + 1
		}
	}
	return -1
}

func convertBookPrefix(prefix string) int {
	switch prefix {
	case "i":
		return 1
	case "ii":
		return 2
	case "iii":
		return 3
	case "first":
		return 1
	case "second":
		return 2
	case "third":
		return 3
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	default:
		return -1
	}
}
