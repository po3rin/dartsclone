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

package progressbar

import (
	progressbar "github.com/schollz/progressbar/v2"
)

type ProgressBar struct {
	*progressbar.ProgressBar
}

func New() *ProgressBar {
	return &ProgressBar{}
}

func (p *ProgressBar) SetMaximum(max int) {
	p.ProgressBar = progressbar.New(max)
}

func (p *ProgressBar) Increment() {
	if p.ProgressBar != nil {
		p.Add(1)
	}
}
