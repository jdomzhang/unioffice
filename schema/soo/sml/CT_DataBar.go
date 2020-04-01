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
	"strconv"

	"github.com/yansuan/unioffice"
)

type CT_DataBar struct {
	// Minimum Length
	MinLengthAttr *uint32
	// Maximum Length
	MaxLengthAttr *uint32
	// Show Values
	ShowValueAttr *bool
	// Conditional Format Value Object
	Cfvo []*CT_Cfvo
	// Data Bar Color
	Color *CT_Color
}

func NewCT_DataBar() *CT_DataBar {
	ret := &CT_DataBar{}
	ret.Color = NewCT_Color()
	return ret
}

func (m *CT_DataBar) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.MinLengthAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minLength"},
			Value: fmt.Sprintf("%v", *m.MinLengthAttr)})
	}
	if m.MaxLengthAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxLength"},
			Value: fmt.Sprintf("%v", *m.MaxLengthAttr)})
	}
	if m.ShowValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showValue"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowValueAttr))})
	}
	e.EncodeToken(start)
	secfvo := xml.StartElement{Name: xml.Name{Local: "ma:cfvo"}}
	for _, c := range m.Cfvo {
		e.EncodeElement(c, secfvo)
	}
	secolor := xml.StartElement{Name: xml.Name{Local: "ma:color"}}
	e.EncodeElement(m.Color, secolor)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DataBar) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Color = NewCT_Color()
	for _, attr := range start.Attr {
		if attr.Name.Local == "minLength" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.MinLengthAttr = &pt
			continue
		}
		if attr.Name.Local == "maxLength" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.MaxLengthAttr = &pt
			continue
		}
		if attr.Name.Local == "showValue" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowValueAttr = &parsed
			continue
		}
	}
lCT_DataBar:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "cfvo"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "cfvo"}:
				tmp := NewCT_Cfvo()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cfvo = append(m.Cfvo, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "color"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "color"}:
				if err := d.DecodeElement(m.Color, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_DataBar %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DataBar
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DataBar and its children
func (m *CT_DataBar) Validate() error {
	return m.ValidateWithPath("CT_DataBar")
}

// ValidateWithPath validates the CT_DataBar and its children, prefixing error messages with path
func (m *CT_DataBar) ValidateWithPath(path string) error {
	for i, v := range m.Cfvo {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Cfvo[%d]", path, i)); err != nil {
			return err
		}
	}
	if err := m.Color.ValidateWithPath(path + "/Color"); err != nil {
		return err
	}
	return nil
}
