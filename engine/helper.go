package engine

import "encoding/xml"

type H map[string]interface{}

func (h H) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name = xml.Name{
		Space: "",
		Local: "map",
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for key, value := range h {

		elem := xml.StartElement{
			Name: xml.Name{
				Space: "",
				Local: key,
			},
			Attr: []xml.Attr{},
		}

		if err := e.EncodeElement(value, elem); err != nil {
			return err
		}
	}

	if err := e.EncodeToken(
		xml.EndElement{Name: start.Name}); err != nil {
		return err
	}

	return nil
}
