// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/tea"
)

type EditControlGroupRequest struct {
  // {"en":"Control Group name, which only the User Customized type Control Group can be modified, customer type Control Group and product type Control Group can not be modified. User Customized type Control Group keeps the original Control Group name if no value is passed", "zh_CN":"Control Group名称，只有自定义类型的Control Group可做修改，若是客户类型与合同类型Control Group则不做修改。自定义类型Control Group若不传值则保持原来的Control Group名称"}
  ControlGroupName *string `json:"controlGroupName,omitempty" xml:"controlGroupName,omitempty"`
  // {"en":"Account object array,Used to specify accounts with permission.  all types of Control Group can be modified, if no value is passed, the original accountList will be emptied", "zh_CN":"账号对象数组, 用来指定有权限访问的账号。客户类型，合同类型与自定义类型的Control Group都可以做修改，若不传值则将原accountList清空"}
  AccountList []*LoginName `json:"accountList,omitempty" xml:"accountList,omitempty" type:"Repeated"`
  // {"en":"Domain array, which only the User Customized type Control Group can be modified, customer type Control Group and product type Control Group can not be modified.User Customized type Control Group empties the original domainList if no value is passed", "zh_CN":"域名字符串数组，只有自定义类型的Control Group可做修改，若是客户类型与合同类型Control Group则不做修改。自定义类型Control Group若不传值则将原domainList清空"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" type:"Repeated"`
  // {"en":"Whether to add:
  // 1. Do not pass or pass false: rewrite method;
  // 2. Pass true: append method.", "zh_CN":"是否追加:
  // 1.不传或false：覆盖方式;
  // 2.传true：追加方式."}
  IsAdd *bool `json:"isAdd,omitempty" xml:"isAdd,omitempty"`
}

func (s EditControlGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s EditControlGroupRequest) GoString() string {
  return s.String()
}

func (s *EditControlGroupRequest) SetControlGroupName(v string) *EditControlGroupRequest {
  s.ControlGroupName = &v
  return s
}

func (s *EditControlGroupRequest) SetAccountList(v []*LoginName) *EditControlGroupRequest {
  s.AccountList = v
  return s
}

func (s *EditControlGroupRequest) SetDomainList(v []*string) *EditControlGroupRequest {
  s.DomainList = v
  return s
}

func (s *EditControlGroupRequest) SetIsAdd(v bool) *EditControlGroupRequest {
  s.IsAdd = &v
  return s
}

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

type EditControlGroupResponse struct {
  // {"en":"Status Code", "zh_CN":"错误具体状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message", "zh_CN":"消息提示"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Success Mark", "zh_CN":"成功标记"}
  Success *bool `json:"success,omitempty" xml:"success,omitempty" require:"true"`
  // {"en":"Control Group Name", "zh_CN":"Control Group名称"}
  ControlGroupName *string `json:"controlGroupName,omitempty" xml:"controlGroupName,omitempty" require:"true"`
}

func (s EditControlGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s EditControlGroupResponse) GoString() string {
  return s.String()
}

func (s *EditControlGroupResponse) SetCode(v int) *EditControlGroupResponse {
  s.Code = &v
  return s
}

func (s *EditControlGroupResponse) SetMsg(v string) *EditControlGroupResponse {
  s.Msg = &v
  return s
}

func (s *EditControlGroupResponse) SetSuccess(v bool) *EditControlGroupResponse {
  s.Success = &v
  return s
}

func (s *EditControlGroupResponse) SetControlGroupName(v string) *EditControlGroupResponse {
  s.ControlGroupName = &v
  return s
}

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

type Parameters struct {
}

func (s Parameters) String() string {
  return tea.Prettify(s)
}

func (s Parameters) GoString() string {
  return s.String()
}

type RequestHeader struct {
}

func (s RequestHeader) String() string {
  return tea.Prettify(s)
}

func (s RequestHeader) GoString() string {
  return s.String()
}

type ResponseHeader struct {
}

func (s ResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ResponseHeader) GoString() string {
  return s.String()
}


