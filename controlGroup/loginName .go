package main

import "github.com/alibabacloud-go/tea/tea"

type LoginName struct {
	// {"en":"Account", "zh_CN":"账号"}
	LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty"`
}

func (s LoginName) String() string {
	return tea.Prettify(s)
}

func (s LoginName) GoString() string {
	return s.String()
}

func (s *LoginName) SetLoginName(v string) *LoginName {
	s.LoginName = &v
	return s
}
