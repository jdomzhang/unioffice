// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/yansuan/unioffice/schema/soo/ofc/sharedTypes"
)

type CT_Ind struct {
	// Start Indentation
	StartAttr *ST_SignedTwipsMeasure
	// Start Indentation in Character Units
	StartCharsAttr *int64
	// End Indentation
	EndAttr *ST_SignedTwipsMeasure
	// End Indentation in Character Units
	EndCharsAttr *int64
	// Start Indentation
	LeftAttr *ST_SignedTwipsMeasure
	// Start Indentation in Character Units
	LeftCharsAttr *int64
	// End Indentation
	RightAttr *ST_SignedTwipsMeasure
	// End Indentation in Character Units
	RightCharsAttr *int64
	// Indentation Removed from First Line
	HangingAttr *sharedTypes.ST_TwipsMeasure
	// Indentation Removed From First Line in Character Units
	HangingCharsAttr *int64
	// Additional First Line Indentation
	FirstLineAttr *sharedTypes.ST_TwipsMeasure
	// Additional First Line Indentation in Character Units
	FirstLineCharsAttr *int64
}

func NewCT_Ind() *CT_Ind {
	ret := &CT_Ind{}
	return ret
}

func (m *CT_Ind) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.StartAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:start"},
			Value: fmt.Sprintf("%v", *m.StartAttr)})
	}
	if m.StartCharsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:startChars"},
			Value: fmt.Sprintf("%v", *m.StartCharsAttr)})
	}
	if m.EndAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:end"},
			Value: fmt.Sprintf("%v", *m.EndAttr)})
	}
	if m.EndCharsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:endChars"},
			Value: fmt.Sprintf("%v", *m.EndCharsAttr)})
	}
	if m.LeftAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:left"},
			Value: fmt.Sprintf("%v", *m.LeftAttr)})
	}
	if m.LeftCharsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:leftChars"},
			Value: fmt.Sprintf("%v", *m.LeftCharsAttr)})
	}
	if m.RightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:right"},
			Value: fmt.Sprintf("%v", *m.RightAttr)})
	}
	if m.RightCharsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rightChars"},
			Value: fmt.Sprintf("%v", *m.RightCharsAttr)})
	}
	if m.HangingAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hanging"},
			Value: fmt.Sprintf("%v", *m.HangingAttr)})
	}
	if m.HangingCharsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hangingChars"},
			Value: fmt.Sprintf("%v", *m.HangingCharsAttr)})
	}
	if m.FirstLineAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:firstLine"},
			Value: fmt.Sprintf("%v", *m.FirstLineAttr)})
	}
	if m.FirstLineCharsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:firstLineChars"},
			Value: fmt.Sprintf("%v", *m.FirstLineCharsAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Ind) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "start" {
			parsed, err := ParseUnionST_SignedTwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.StartAttr = &parsed
			continue
		}
		if attr.Name.Local == "startChars" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.StartCharsAttr = &parsed
			continue
		}
		if attr.Name.Local == "end" {
			parsed, err := ParseUnionST_SignedTwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.EndAttr = &parsed
			continue
		}
		if attr.Name.Local == "endChars" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.EndCharsAttr = &parsed
			continue
		}
		if attr.Name.Local == "left" {
			parsed, err := ParseUnionST_SignedTwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.LeftAttr = &parsed
			continue
		}
		if attr.Name.Local == "leftChars" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.LeftCharsAttr = &parsed
			continue
		}
		if attr.Name.Local == "right" {
			parsed, err := ParseUnionST_SignedTwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.RightAttr = &parsed
			continue
		}
		if attr.Name.Local == "rightChars" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.RightCharsAttr = &parsed
			continue
		}
		if attr.Name.Local == "hanging" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.HangingAttr = &parsed
			continue
		}
		if attr.Name.Local == "hangingChars" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.HangingCharsAttr = &parsed
			continue
		}
		if attr.Name.Local == "firstLine" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.FirstLineAttr = &parsed
			continue
		}
		if attr.Name.Local == "firstLineChars" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.FirstLineCharsAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Ind: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Ind and its children
func (m *CT_Ind) Validate() error {
	return m.ValidateWithPath("CT_Ind")
}

// ValidateWithPath validates the CT_Ind and its children, prefixing error messages with path
func (m *CT_Ind) ValidateWithPath(path string) error {
	if m.StartAttr != nil {
		if err := m.StartAttr.ValidateWithPath(path + "/StartAttr"); err != nil {
			return err
		}
	}
	if m.EndAttr != nil {
		if err := m.EndAttr.ValidateWithPath(path + "/EndAttr"); err != nil {
			return err
		}
	}
	if m.LeftAttr != nil {
		if err := m.LeftAttr.ValidateWithPath(path + "/LeftAttr"); err != nil {
			return err
		}
	}
	if m.RightAttr != nil {
		if err := m.RightAttr.ValidateWithPath(path + "/RightAttr"); err != nil {
			return err
		}
	}
	if m.HangingAttr != nil {
		if err := m.HangingAttr.ValidateWithPath(path + "/HangingAttr"); err != nil {
			return err
		}
	}
	if m.FirstLineAttr != nil {
		if err := m.FirstLineAttr.ValidateWithPath(path + "/FirstLineAttr"); err != nil {
			return err
		}
	}
	return nil
}
