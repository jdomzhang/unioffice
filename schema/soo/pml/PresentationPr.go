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

	"github.com/yansuan/unioffice"
	"github.com/yansuan/unioffice/schema/soo/dml"
)

type PresentationPr struct {
	CT_PresentationProperties
}

func NewPresentationPr() *PresentationPr {
	ret := &PresentationPr{}
	ret.CT_PresentationProperties = *NewCT_PresentationProperties()
	return ret
}

func (m *PresentationPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:p"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "p:presentationPr"
	return m.CT_PresentationProperties.MarshalXML(e, start)
}

func (m *PresentationPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_PresentationProperties = *NewCT_PresentationProperties()
lPresentationPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "htmlPubPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "htmlPubPr"}:
				m.HtmlPubPr = NewCT_HtmlPublishProperties()
				if err := d.DecodeElement(m.HtmlPubPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "webPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "webPr"}:
				m.WebPr = NewCT_WebProperties()
				if err := d.DecodeElement(m.WebPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "prnPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "prnPr"}:
				m.PrnPr = NewCT_PrintProperties()
				if err := d.DecodeElement(m.PrnPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "showPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "showPr"}:
				m.ShowPr = NewCT_ShowProperties()
				if err := d.DecodeElement(m.ShowPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "clrMru"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "clrMru"}:
				m.ClrMru = dml.NewCT_ColorMRU()
				if err := d.DecodeElement(m.ClrMru, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on PresentationPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lPresentationPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the PresentationPr and its children
func (m *PresentationPr) Validate() error {
	return m.ValidateWithPath("PresentationPr")
}

// ValidateWithPath validates the PresentationPr and its children, prefixing error messages with path
func (m *PresentationPr) ValidateWithPath(path string) error {
	if err := m.CT_PresentationProperties.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
