// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"

	"github.com/yansuan/unioffice"
)

type CT_GvmlShape struct {
	NvSpPr *CT_GvmlShapeNonVisual
	SpPr   *CT_ShapeProperties
	TxSp   *CT_GvmlTextShape
	Style  *CT_ShapeStyle
	ExtLst *CT_OfficeArtExtensionList
}

func NewCT_GvmlShape() *CT_GvmlShape {
	ret := &CT_GvmlShape{}
	ret.NvSpPr = NewCT_GvmlShapeNonVisual()
	ret.SpPr = NewCT_ShapeProperties()
	return ret
}

func (m *CT_GvmlShape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	senvSpPr := xml.StartElement{Name: xml.Name{Local: "a:nvSpPr"}}
	e.EncodeElement(m.NvSpPr, senvSpPr)
	sespPr := xml.StartElement{Name: xml.Name{Local: "a:spPr"}}
	e.EncodeElement(m.SpPr, sespPr)
	if m.TxSp != nil {
		setxSp := xml.StartElement{Name: xml.Name{Local: "a:txSp"}}
		e.EncodeElement(m.TxSp, setxSp)
	}
	if m.Style != nil {
		sestyle := xml.StartElement{Name: xml.Name{Local: "a:style"}}
		e.EncodeElement(m.Style, sestyle)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GvmlShape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NvSpPr = NewCT_GvmlShapeNonVisual()
	m.SpPr = NewCT_ShapeProperties()
lCT_GvmlShape:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "nvSpPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "nvSpPr"}:
				if err := d.DecodeElement(m.NvSpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "spPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "spPr"}:
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "txSp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "txSp"}:
				m.TxSp = NewCT_GvmlTextShape()
				if err := d.DecodeElement(m.TxSp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "style"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "style"}:
				m.Style = NewCT_ShapeStyle()
				if err := d.DecodeElement(m.Style, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_GvmlShape %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GvmlShape
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GvmlShape and its children
func (m *CT_GvmlShape) Validate() error {
	return m.ValidateWithPath("CT_GvmlShape")
}

// ValidateWithPath validates the CT_GvmlShape and its children, prefixing error messages with path
func (m *CT_GvmlShape) ValidateWithPath(path string) error {
	if err := m.NvSpPr.ValidateWithPath(path + "/NvSpPr"); err != nil {
		return err
	}
	if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
		return err
	}
	if m.TxSp != nil {
		if err := m.TxSp.ValidateWithPath(path + "/TxSp"); err != nil {
			return err
		}
	}
	if m.Style != nil {
		if err := m.Style.ValidateWithPath(path + "/Style"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
