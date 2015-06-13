package engine

import (
	"encoding/xml"
	"encoding/json"
	"net/http"
)

const (
	MIME_JSON string = "application/json"
	MIME_XML  string = "application/xml"
)

type (

	Renderer interface {
		Render(http.ResponseWriter, int, ...interface{}) error
	}

	rendererJSON struct{}
	rendererXML  struct{}
)

var (
	RenderDriverJSON = rendererJSON{}
	RenderDriverXML  = rendererXML{}
)

func (_ rendererJSON) Render(writer http.ResponseWriter, status int, data ...interface{}) error {

	// render header data
	writer.Header().Set("Content-Type", MIME_JSON + ";charset=utf-8")
	writer.WriteHeader(status)

	// encode response
	enc := json.NewEncoder(writer)
	err := enc.Encode(data[0])

	return err
}

func (_ rendererXML) Render(writer http.ResponseWriter, status int, data ...interface{}) error {

	// render header data
	writer.Header().Set("Content-Type", MIME_XML + ";charset=utf-8")
	writer.WriteHeader(status)

	// encode response
	enc := xml.NewEncoder(writer)
	err := enc.Encode(data[0])

	return err
}
