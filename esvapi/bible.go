package esvapi


import (
    bible "../."
    "net/http"
    "net/url"
    "strings"
    "encoding/xml"
)

const baseUrl string = "http://www.esvapi.org/v2/rest/"
const passageQuery string = baseUrl + "passageQuery?"

const EnglishStandardVersion string = "crossway-esv"

func EsvApiProvider() esvApiProvider {
    return esvApiProvider(1)
}


type esvApiProvider int

func (esv esvApiProvider) Text(translation string, passage bible.Passage) (string, error) {
    fixedTitle := strings.Replace(passage.Title(), " ", "+", -1)
    
    encoded := url.Values{
        "key": {"TEST"},
        "passage": {fixedTitle},
        "output-format": {url.QueryEscape("crossway-xml-1.0")},
         "include-quote-entities": {url.QueryEscape("false")},
        }.Encode()
    
    resp, err := http.Get(passageQuery + encoded)
        
   if err != nil {return "", err}
  
   
   decoder := xml.NewDecoder(resp.Body)
   
   inVerseUnit := false
   stacked := 0
   
   text := ""
   
   
   token, err := decoder.Token()
   
   for token != nil && err == nil  {
       switch t := token.(type) {
       case xml.StartElement:
            if t.Name.Local == "verse-unit" {
                inVerseUnit = true
            } else if inVerseUnit && t.Name.Local != "woc" {
                stacked += 1
            }
            
            if t.Name.Local == "q" {
                text += "\""
            }
            
       case xml.EndElement:
            if t.Name.Local == "verse-unit" {
                inVerseUnit = false
                stacked = 0
                text += " "
            } else if inVerseUnit && t.Name.Local != "woc" {
                stacked -= 1
            }
            
       case xml.CharData:
            if inVerseUnit && stacked == 0 {
                text += string(t)
            }
       }
       
       token, err = decoder.Token()
   }
   
    return text, nil
}

func (esv esvApiProvider) TranslationsProvided() ([]string, error) {
    return []string{EnglishStandardVersion}, nil
}