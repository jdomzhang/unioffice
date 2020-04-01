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

type CT_SingleXmlCell struct {
	// Table Id
	IdAttr uint32
	// Reference
	RAttr string
	// Connection ID
	ConnectionIdAttr uint32
	// Cell Properties
	XmlCellPr *CT_XmlCellPr
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_SingleXmlCell() *CT_SingleXmlCell {
	ret := &CT_SingleXmlCell{}
	ret.XmlCellPr = NewCT_XmlCellPr()
	return ret
}

func (m *CT_SingleXmlCell) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
		Value: fmt.Sprintf("%v", m.RAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "connectionId"},
		Value: fmt.Sprintf("%v", m.ConnectionIdAttr)})
	e.EncodeToken(start)
	sexmlCellPr := xml.StartElement{Name: xml.Name{Local: "ma:xmlCellPr"}}
	e.EncodeElement(m.XmlCellPr, sexmlCellPr)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SingleXmlCell) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.XmlCellPr = NewCT_XmlCellPr()
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "r" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RAttr = parsed
			continue
		}
		if attr.Name.Local == "connectionId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ConnectionIdAttr = uint32(parsed)
			continue
		}
	}
lCT_SingleXmlCell:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "xmlCellPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "xmlCellPr"}:
				if err := d.DecodeElement(m.XmlCellPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_SingleXmlCell %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SingleXmlCell
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SingleXmlCell and its children
func (m *CT_SingleXmlCell) Validate() error {
	return m.ValidateWithPath("CT_SingleXmlCell")
}

// ValidateWithPath validates the CT_SingleXmlCell and its children, prefixing error messages with path
func (m *CT_SingleXmlCell) ValidateWithPath(path string) error {
	if err := m.XmlCellPr.ValidateWithPath(path + "/XmlCellPr"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
