package esvapi


import (
    bible "github.com/andrewmaach/gobible"
    "net/http"
    "net/url"
    "strings"
    "encoding/xml"
)

const baseUrl string = "http://www.esvapi.org/v2/rest/"
const passageQuery string = baseUrl + "passageQuery?"

const EnglishStandardVersion string = "crossway-esv"

func EsvApiProvider(accessCode string) esvApiProvider {
    return esvApiProvider(accessCode)
}


type esvApiProvider string

func (esv esvApiProvider) accessCode() string {
    return string(esv)
}

func (esv esvApiProvider) Text(translation string, passage bible.Passage) (string, error) {
    title, err := passage.Title()
    if err != nil {return "", err}
    
    fixedTitle := strings.Replace(title, " ", "+", -1)
    
    encoded := url.Values{
        "key": {esv.accessCode()},
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
            } else if inVerseUnit && t.Name.Local != "woc" && t.Name.Local != "span" {
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
            } else if inVerseUnit && t.Name.Local != "woc" && t.Name.Local != "span" {
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