// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/tea"
)

type DisableSingleDomainServiceRequest struct {
}

func (s DisableSingleDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s DisableSingleDomainServiceRequest) GoString() string {
  return s.String()
}

type DisableSingleDomainServiceResponse struct {
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s DisableSingleDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s DisableSingleDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *DisableSingleDomainServiceResponse) SetCode(v string) *DisableSingleDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *DisableSingleDomainServiceResponse) SetMessage(v string) *DisableSingleDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *DisableSingleDomainServiceResponse) SetHttpStatus(v int) *DisableSingleDomainServiceResponse {
  s.HttpStatus = &v
  return s
}

func (s *DisableSingleDomainServiceResponse) SetXCncRequestId(v string) *DisableSingleDomainServiceResponse {
  s.XCncRequestId = &v
  return s
}

type Paths struct {
  // {"en":"", "zh_CN":"加速域名在系统中对应的ID
  // 1. 参看请求示例中的url，123344对应的就是domain-id
  // 2. 可以通过【获取域名配置】和【获取域名列表】接口查询到domain-id"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
}

func (s Paths) String() string {
  return tea.Prettify(s)
}

func (s Paths) GoString() string {
  return s.String()
}

func (s *Paths) SetDomainId(v int) *Paths {
  s.DomainId = &v
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


