// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml

import (
	"encoding/xml"
	"fmt"

	"github.com/yansuan/unioffice"
)

type CT_TLTimeAnimateValueList struct {
	// Time Animate Value
	Tav []*CT_TLTimeAnimateValue
}

func NewCT_TLTimeAnimateValueList() *CT_TLTimeAnimateValueList {
	ret := &CT_TLTimeAnimateValueList{}
	return ret
}

func (m *CT_TLTimeAnimateValueList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Tav != nil {
		setav := xml.StartElement{Name: xml.Name{Local: "p:tav"}}
		for _, c := range m.Tav {
			e.EncodeElement(c, setav)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLTimeAnimateValueList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TLTimeAnimateValueList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "tav"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "tav"}:
				tmp := NewCT_TLTimeAnimateValue()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tav = append(m.Tav, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_TLTimeAnimateValueList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLTimeAnimateValueList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLTimeAnimateValueList and its children
func (m *CT_TLTimeAnimateValueList) Validate() error {
	return m.ValidateWithPath("CT_TLTimeAnimateValueList")
}

// ValidateWithPath validates the CT_TLTimeAnimateValueList and its children, prefixing error messages with path
func (m *CT_TLTimeAnimateValueList) ValidateWithPath(path string) error {
	for i, v := range m.Tav {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tav[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
