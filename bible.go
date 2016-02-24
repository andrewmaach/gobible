package gobible

import (
    "errors"
)




var NoNewTranslationsError error = errors.New("No new translations were available in given provider.")

func Bible(cache BibleCacheProvider) BibleHandle {
    return BibleHandle{
        providers: make(map[string]BibleTextProvider),
        cache: cache,
    }
}

type BibleHandle struct {
    providers map[string]BibleTextProvider // key=translation
    cache BibleCacheProvider
    defaultTranslation string
}

func (bible *BibleHandle) RegisterProvider(handler BibleTextProvider) (error) {
    translations, err := handler.TranslationsProvided()
    if err != nil {return err}
    
    useful := false
    for _, translation := range translations {
        _, ok := bible.providers[translation]
        if !ok {
            useful = true
            bible.providers[translation] = handler
            bible.defaultTranslation = translation
        }
    }
    
    if !useful {
        return NoNewTranslationsError
    }
    
    return nil
    
    
}

func (bible *BibleHandle) Text(translation string, passage Passage) (string, error) {
    return bible.providers[translation].Text(translation, passage)
}

func (bible *BibleHandle) DefaultText(passage Passage) (string, error) {
    psg, err := shortPassage(passage)
    if err != nil {
        return "", err
    }
    
    res, err := bible.providers[bible.defaultTranslation].Text(bible.defaultTranslation, psg)
    
    if psg != passage {
        return res + "...", err
    } else {
        return res, err
    }
}


type BibleTextProvider interface {
    Text(translation string, passage Passage) (string, error)
    TranslationsProvided() ([]string, error)
}


// BibleCacheProvider a standard interface for 
type BibleCacheProvider interface {
    Save(text, translation string, passage Passage) error
    Recall(translation string, passage Passage) (string, error)
}
