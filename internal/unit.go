// Copyright 2018 ikawaha
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// 	You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"fmt"
)

const maxOffset = 1 << 29

type unit uint32

const unitSize = 4

func (u *unit) setHasLeaf(hasLeaf bool) {
	if hasLeaf {
		*u = unit(int32(*u) | 1<<8)
		return
	}
	*u = unit(uint32(*u) & ^(uint32(1) << 8))
}

func (u *unit) setValue(value uint32) {
	*u = unit(value | (1 << 31))
}

func (u *unit) setLabel(label byte) {
	*u = unit(uint32(*u) & ^uint32(0xFF) | uint32(label))
}

func (u *unit) setOffset(offset uint32) error {
	if offset >= maxOffset {
		return fmt.Errorf("failed to modify unit, too large offset")
	}
	*u = unit(uint32(*u) & ((1 << 31) | (1 << 8) | 0xFF))
	if offset < 1<<21 {
		*u = unit(uint32(*u) | (offset << 10))
		return nil
	}
	*u = unit(uint32(*u) | ((offset << 2) | (1 << 9)))
	return nil
}

func (u unit) label() byte {
	return byte(uint32(u) & ((1 << 31) | 0xFF))
}

func (u unit) offset() uint32 {
	return (uint32(u) >> 10) << ((uint32(u) & (1 << 9)) >> 6)
}

func (u unit) hasLeaf() bool {
	return ((uint32(u) >> 8) & 1) == 1
}

func (u unit) value() uint32 {
	return uint32(u) & ((1 << 31) - 1)
}
