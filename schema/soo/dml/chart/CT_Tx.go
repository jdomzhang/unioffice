// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"

	"github.com/yansuan/unioffice"
)

type CT_Tx struct {
	Choice *CT_TxChoice
}

func NewCT_Tx() *CT_Tx {
	ret := &CT_Tx{}
	ret.Choice = NewCT_TxChoice()
	return ret
}

func (m *CT_Tx) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	m.Choice.MarshalXML(e, xml.StartElement{})
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Tx) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Choice = NewCT_TxChoice()
lCT_Tx:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "strRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "strRef"}:
				m.Choice = NewCT_TxChoice()
				if err := d.DecodeElement(&m.Choice.StrRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "rich"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "rich"}:
				m.Choice = NewCT_TxChoice()
				if err := d.DecodeElement(&m.Choice.Rich, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Tx %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Tx
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Tx and its children
func (m *CT_Tx) Validate() error {
	return m.ValidateWithPath("CT_Tx")
}

// ValidateWithPath validates the CT_Tx and its children, prefixing error messages with path
func (m *CT_Tx) ValidateWithPath(path string) error {
	if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
		return err
	}
	return nil
}
