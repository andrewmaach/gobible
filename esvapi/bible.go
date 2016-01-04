package esvapi


import (
    bible "../."
    "net/http"
    "net/url"
    "log"
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
    //fixedTitle := strings.Replace("John 3:1-20", " ", "+", -1)
    
    // crossway-bible passage content verse-unit
    
    type woc struct {
        
    }
    
    type cnt struct {
        VerseUnit    []string   `xml:"verse-unit"`
    }
    
    type psg struct {
        Content    cnt   `xml:"content"`
    }
    
    type cb struct {
        Passage    psg   `xml:"passage"`
    }
    

    

    
    
    encoded := url.Values{
        "key": {"TEST"},
        "passage": {fixedTitle},
        "output-format": {url.QueryEscape("crossway-xml-1.0")},
         "include-quote-entities": {url.QueryEscape("false")},
        }.Encode()
    
    
    
    resp, err := http.Get(passageQuery + encoded)
        
   if err != nil {return "", err}
   
   var v cb
   
   
   
   err = xml.NewDecoder(resp.Body).Decode(&v)
   if err != nil {return "", err}
   
   text := ""
   for _, unit := range v.Passage.Content.VerseUnit {
       text += unit + " "
   }
    
    log.Println(text)
    
    return text, nil
}

func (esv esvApiProvider) TranslationsProvided() ([]string, error) {
    return []string{EnglishStandardVersion}, nil
}