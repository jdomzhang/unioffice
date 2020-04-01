// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/yansuan/unioffice/schema/soo/ofc/math"
)

func TestEG_OMathMathElementsConstructor(t *testing.T) {
	v := math.NewEG_OMathMathElements()
	if v == nil {
		t.Errorf("math.NewEG_OMathMathElements must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.EG_OMathMathElements should validate: %s", err)
	}
}

func TestEG_OMathMathElementsMarshalUnmarshal(t *testing.T) {
	v := math.NewEG_OMathMathElements()
	buf, _ := xml.Marshal(v)
	v2 := math.NewEG_OMathMathElements()
	xml.Unmarshal(buf, v2)
}
