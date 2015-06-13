package engine

import (
	"encoding/xml"
	"encoding/json"
	"net/http"
)

type (

	Parser interface {
		Parse(*http.Request, interface{}) error
	}

	parserJSON struct{}
	parserXML  struct{}
)

var (
	ParserDriverJSON = parserJSON{}
	ParserDriverXML  = parserXML{}
)

func (_ parserJSON) Parse(req *http.Request, data interface{}) error {

    // decode request body
    dec := json.NewDecoder(req.Body)
    err := dec.Decode(data)

	return err
}

func (_ parserXML) Parse(req *http.Request, data interface{}) error {

	// decode request body
	dec := xml.NewDecoder(req.Body)
	err := dec.Decode(data)

	return err
}
