// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml

import (
	"encoding/xml"

	"github.com/yansuan/unioffice"
)

type Handles struct {
	CT_Handles
}

func NewHandles() *Handles {
	ret := &Handles{}
	ret.CT_Handles = *NewCT_Handles()
	return ret
}

func (m *Handles) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_Handles.MarshalXML(e, start)
}

func (m *Handles) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Handles = *NewCT_Handles()
lHandles:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "h"}:
				tmp := NewCT_H()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.H = append(m.H, tmp)
			default:
				unioffice.Log("skipping unsupported element on Handles %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lHandles
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Handles and its children
func (m *Handles) Validate() error {
	return m.ValidateWithPath("Handles")
}

// ValidateWithPath validates the Handles and its children, prefixing error messages with path
func (m *Handles) ValidateWithPath(path string) error {
	if err := m.CT_Handles.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
