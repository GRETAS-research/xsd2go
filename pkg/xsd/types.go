package xsd

import (
	"encoding/xml"

	"github.com/iancoleman/strcase"
)

type Type interface {
	GoName() string
}

type ComplexType struct {
	XMLName    xml.Name    `xml:"http://www.w3.org/2001/XMLSchema complexType"`
	Name       string      `xml:"name,attr"`
	Mixed      string      `xml:"mixed,attr"`
	Attributes []Attribute `xml:"attribute"`
	Sequence   *Sequence   `xml:"sequence"`
}

func (ct *ComplexType) Elements() []Element {
	if ct.Sequence != nil {
		return ct.Sequence.Elements
	}
	return []Element{}
}

func (ct *ComplexType) GoName() string {
	return strcase.ToCamel(ct.Name)
}

type Sequence struct {
	XMLName  xml.Name  `xml:"http://www.w3.org/2001/XMLSchema sequence"`
	Elements []Element `xml:"element"`
}

type SimpleType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema simpleType"`
	Name    string   `xml:"name,attr"`
}

func (st *SimpleType) GoName() string {
	return strcase.ToCamel(st.Name)
}

type staticType string

func (st staticType) GoName() string {
	return string(st)
}
