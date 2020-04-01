// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/yansuan/unioffice/schema/soo/wml"
)

func TestCT_LvlLegacyConstructor(t *testing.T) {
	v := wml.NewCT_LvlLegacy()
	if v == nil {
		t.Errorf("wml.NewCT_LvlLegacy must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_LvlLegacy should validate: %s", err)
	}
}

func TestCT_LvlLegacyMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_LvlLegacy()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_LvlLegacy()
	xml.Unmarshal(buf, v2)
}
