package main

import "github.com/alibabacloud-go/tea/tea"

type Paths struct {
	// {"en":"Control Group Code", "zh_CN":"Control Group 编号"}
	ControlGroupCode *string `json:"ControlGroupCode,omitempty" xml:"ControlGroupCode,omitempty" require:"true"`
}

func (s Paths) String() string {
	return tea.Prettify(s)
}

func (s Paths) GoString() string {
	return s.String()
}

func (s *Paths) SetControlGroupCode(v string) *Paths {
	s.ControlGroupCode = &v
	return s
}
