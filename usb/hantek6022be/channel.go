//  Copyright 2016 The goscope Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package hantek6022be

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/zagrodzki/goscope/scope"
)

type ch struct {
	id        scope.ChanID
	osc       *Scope
	voltRange scope.VoltRange
}

func (c ch) ID() scope.ChanID { return c.id }
func (ch) GetVoltRanges() []scope.VoltRange {
	return voltRanges
}
func (c ch) GetVoltRange() scope.VoltRange {
	return c.voltRange
}
func (c ch) SetVoltRange(v scope.VoltRange) error {
	var req uint8
	switch c.id {
	case "CH1":
		req = ch1VoltRangeReq
	case "CH2":
		req = ch2VoltRangeReq
	}
	val, ok := voltRangeToID[v]
	if !ok {
		return errors.New(fmt.Sprintf("Channel %s: SetVoltRange(%s): range must be one of %v", c, v, voltRanges))
	}
	if _, err := c.osc.dev.Control(0x40, req, 0, 0, val.data()); err != nil {
		return errors.Wrapf(err, "Control(voltage range %s(%x))", v, val)
	}
	c.voltRange = v
	return nil
}
