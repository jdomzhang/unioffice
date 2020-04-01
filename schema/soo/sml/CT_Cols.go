// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/yansuan/unioffice"
)

type CT_Cols struct {
	// Column Width & Formatting
	Col []*CT_Col
}

func NewCT_Cols() *CT_Cols {
	ret := &CT_Cols{}
	return ret
}

func (m *CT_Cols) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secol := xml.StartElement{Name: xml.Name{Local: "ma:col"}}
	for _, c := range m.Col {
		e.EncodeElement(c, secol)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Cols) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Cols:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "col"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "col"}:
				tmp := NewCT_Col()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Col = append(m.Col, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_Cols %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Cols
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Cols and its children
func (m *CT_Cols) Validate() error {
	return m.ValidateWithPath("CT_Cols")
}

// ValidateWithPath validates the CT_Cols and its children, prefixing error messages with path
func (m *CT_Cols) ValidateWithPath(path string) error {
	for i, v := range m.Col {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Col[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
