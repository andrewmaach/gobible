/* Copyright (c) 2015 BibleBlurb LLC
 * Authored By Andrew Maach
 */
package bible


import (
    "strconv"
    "encoding/json"
    "os"
)

// Every book of the Bible in order, as an array.
var booksList = []string{
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




var referenceIndex map[string]string

type chapterTextContent map[string]string


func init() {
    /*
    referenceIndex = make(map[string]string)
    
    // LOAD INDEX
    
    indexFile, err := os.Open("bible/kjv_index.txt")
    if err != nil {
        panic("COULD NOT OPEN INDEX FILE.")
    }
    defer indexFile.Close()
    
    jsonParser := json.NewDecoder(indexFile)
    if err = jsonParser.Decode(&referenceIndex); err != nil {
        panic(err)
    }*/
 
}


// Returns the number of Chapters in a given book, by number.
func maxChapters(book int) int {
    file := referenceIndex[booksList[book - 1]]
    

    
    type BookText map[string]chapterTextContent
    
    text := make(BookText)
    
    indexFile, err := os.Open("bible/text/"+file)
    if err != nil {
        panic("COULD NOT OPEN TEXT FILE.")
    }
    defer indexFile.Close()
    
    jsonParser := json.NewDecoder(indexFile)
    if err = jsonParser.Decode(&text); err != nil {
        panic(err)
    }
    
    highest := 0
    for schapter, _ := range text {
        c, _ := strconv.Atoi(schapter)
        if c > highest{
            highest = c
        }
    }
    
    return highest
}


func chapterText(book, chapter int) chapterTextContent {
    
    //file := referenceIndex[booksList[book - 1]]
    
    type BookText map[string]chapterTextContent
    
    text := make(BookText)
    /*
    indexFile, err := os.Open(file)
    if err != nil {
        panic("COULD NOT OPEN TEXT FILE.")
    }
    defer indexFile.Close()
    
    jsonParser := json.NewDecoder(indexFile)
    if err = jsonParser.Decode(&text); err != nil {
        panic(err)
    }
    */
    
    return text[strconv.Itoa(chapter)]
    
}