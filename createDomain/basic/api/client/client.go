// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/tea"
)

type AddCdnDomainRequest struct {
  // {"en":"The current version is 1.0.0", "zh_CN":"版本号，当前版本号1.0.0"}
  Version *string `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"CDN accelerated domain name. Support generic domain name, witch starts with a dot, such as: .example.com, If the domain example.com have China ICP, then the domain name xx.example.com have ICP too.", "zh_CN":"需要接入CDN的域名。支持泛域名，以符号“.”开头，如：.example.com，
  //   泛域名也包含多级 a.b.example.com。 如果example.com已备案，那么域名xx.example.com则不需要备案。"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"The value should be an effective domain. The added domain will copy the config of this domain.", 
  //   "zh_CN":"指定参考域名。当使用参考域名来新增域名是，新域名将复用参考域名的配置"}
  ReferencedDomainName *string `json:"referenced-domain-name,omitempty" xml:"referenced-domain-name,omitempty"`
  // {"en":"Configure single template. If you want to add a new accelerated domain name with a specified configuration, you can use a specified configured template. For details, please consult the corresponding customer support staff.", "zh_CN":"配置单模板，特定的使用场景下，如果希望新增的加速域名参照某些指定配置时，可以指定配置单模板，具体使用请咨询对应的客户负责人。"}
  ConfigFormId *string `json:"config-form-id,omitempty" xml:"config-form-id,omitempty"`
  // {"en":"The optional value is: true or false.", "zh_CN":"是否纯海外加速，入参范围：true、false"}
  AccelerateNoChina *string `json:"accelerate-no-china,omitempty" xml:"accelerate-no-china,omitempty"`
  // {"en":"The id of contract", "zh_CN":"合同号"}
  ContractId *string `json:"contract-id,omitempty" xml:"contract-id,omitempty" require:"true"`
  // {"en":"The id of item", "zh_CN":"产品号"}
  ItemId *string `json:"item-id,omitempty" xml:"item-id,omitempty" require:"true"`
  // {"en":"An alias for public service cname. When you have multiple domains that need to share a cname, then you may specify a cname-label here. If domains use the same cname-label, they will share the same cname and has the same dns coverage. Note: 1. Restrictions on domains with the same cname-label: they should have the same acceleration type,  certificate (if use ssl), service areas. 2. Multiple http domains can share the same cname; multiple sni https domains can share the same cname as wll. 3. When a single domain uses a cname-label, it can be cancelled acceleration; while multiple domains are not allowed to cancel acceleration for part of them.", "zh_CN":"共用一级别名，客户存在较多一级域名共用的需求，因此在接口中引入cname-label标识，即拥有相同cname-label的一组域名，共用一级cname。当加速域名的加速区域、加速类型、资源链是一致的时候，建议使用共用一级别名，相同的业务使用相同一级别名cname。 注意： 1、拥有相同cname-label的域名共用一级cname，且有完全一致的dns覆盖 2、共用一级的约束：配置一致、加速类型一致、证书id一致（如果有证书）、加速区域一致、平台套餐一致 3、多个http域名可共用一级，多个sni https域名可共用一级，多个共享证书域名可共用一级 4、单个域名使用cname-label时，域名可取消加速；多个域名共用一级时，不允许取消加速这些域名 5、支持通过修改cname-label达到修改cname的目的"}
  CnameLabel *string `json:"cname-label,omitempty" xml:"cname-label,omitempty"`
  // {"en":"Remarks, up to 1000 characters", "zh_CN":"备注信息，最大限制1000个字符"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"Pass the response header of client IP. The optional values are Cdn-Src-Ip and X-Forwarded-For. The default value is Cdn-Src-Ip.", "zh_CN":"传递客户端ip的响应头部，可选值为Cdn-Src-Ip和X-Forwarded-For，默认值为Cdn-Src-Ip"}
  HeaderOfClientIp *string `json:"header-of-clientip,omitempty" xml:"header-of-clientip,omitempty"`
  // {"en":"Back-to-origin policy setting, which is used to set the origin site information and the back-to-origin policy of the accelerated domain name", "zh_CN":"回源策略设置，用于设置加速域名的源站信息和回源策略"}
  OriginConfig *AddCdnDomainRequestOriginConfig `json:"origin-config,omitempty" xml:"origin-config,omitempty" require:"true" type:"Struct"`
  // {"en":"Live domain configuration, used to set the livestream acceleration domain origin config.
  // Note: In addition to the API call permission, you need to contact the dedicated customer service to apply for the corresponding API client template.", "zh_CN":"直播域名配置，用于设置直播加速域名的推拉流（使用需申请）
  // 注意：该节点下的相关参数配置，除开通API调用权限外，还需要联系专属客服申请开通对应的API客户模板"}
  LiveConfig *AddCdnDomainRequestLiveConfig `json:"live-config,omitempty" xml:"live-config,omitempty" type:"Struct"`
  // {"en":"Livestream domain settings. Set the publishing point of the live push-pull domain.
  // note:
  // 1. The pull stream and the corresponding push stream domain must be configured with the same publishing point.
  // 2. If you are not going to modify the publishing point, please do not pass this param.
  // 3. The publishing point adopts the overlay update. Each time you modify, you need to submit all the publishing points. You cannot submit only the parts that need to be modified.", "zh_CN":"设置直播推拉流域名的发布点
  // 注意：
  // 1、拉流和对应的推流域名，必须配置相同的发布点；
  // 2、不想修改发布点时，不要传入该节点及以下入参；
  // 3、发布点采用覆盖式更新，每次修改时，需要提交全部发布点，不能仅提交需要修改的部分。"}
  PublishPoints []*AddCdnDomainRequestPublishPoints `json:"publish-points,omitempty" xml:"publish-points,omitempty" type:"Repeated"`
}

func (s AddCdnDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s AddCdnDomainRequest) GoString() string {
  return s.String()
}

func (s *AddCdnDomainRequest) SetVersion(v string) *AddCdnDomainRequest {
  s.Version = &v
  return s
}

func (s *AddCdnDomainRequest) SetDomainName(v string) *AddCdnDomainRequest {
  s.DomainName = &v
  return s
}

func (s *AddCdnDomainRequest) SetReferencedDomainName(v string) *AddCdnDomainRequest {
  s.ReferencedDomainName = &v
  return s
}

func (s *AddCdnDomainRequest) SetConfigFormId(v string) *AddCdnDomainRequest {
  s.ConfigFormId = &v
  return s
}

func (s *AddCdnDomainRequest) SetAccelerateNoChina(v string) *AddCdnDomainRequest {
  s.AccelerateNoChina = &v
  return s
}

func (s *AddCdnDomainRequest) SetContractId(v string) *AddCdnDomainRequest {
  s.ContractId = &v
  return s
}

func (s *AddCdnDomainRequest) SetItemId(v string) *AddCdnDomainRequest {
  s.ItemId = &v
  return s
}

func (s *AddCdnDomainRequest) SetCnameLabel(v string) *AddCdnDomainRequest {
  s.CnameLabel = &v
  return s
}

func (s *AddCdnDomainRequest) SetComment(v string) *AddCdnDomainRequest {
  s.Comment = &v
  return s
}

func (s *AddCdnDomainRequest) SetHeaderOfClientIp(v string) *AddCdnDomainRequest {
  s.HeaderOfClientIp = &v
  return s
}

func (s *AddCdnDomainRequest) SetOriginConfig(v *AddCdnDomainRequestOriginConfig) *AddCdnDomainRequest {
  s.OriginConfig = v
  return s
}

func (s *AddCdnDomainRequest) SetLiveConfig(v *AddCdnDomainRequestLiveConfig) *AddCdnDomainRequest {
  s.LiveConfig = v
  return s
}

func (s *AddCdnDomainRequest) SetPublishPoints(v []*AddCdnDomainRequestPublishPoints) *AddCdnDomainRequest {
  s.PublishPoints = v
  return s
}

type AddCdnDomainRequestOriginConfig struct {
  // {"en":"Origin site address, which can be an IP or a domain name.
  // 1. IP is separated by semicolons and multiple IPs are supported.
  // 2. Only one domain name can be entered. IP and domain names cannot be entered at the same time.
  // 3. Maximum character limit is 500.", "zh_CN":"回源地址，可以是IP或域名。 1. IP以分号分隔，支持多个。 2. 域名只能输入一个。IP与域名不能同时输入。 3.限制最大不能超过500个字符长度。"}
  OriginIps *string `json:"origin-ips,omitempty" xml:"origin-ips,omitempty"`
  // {"en":"The Origin HOST for changing the HOST field in the return source HTTP request header.
  // Note: It should be domain or IP format. For domain name format, each segement separated by a dot, does not exceed 62 characters, the total length should not exceed 128 characters.", "zh_CN":"回源HOST，用于更改回源HTTP请求头中的HOST字段。支持格式为: 域名或ip。
  // 注意：必须符合ip/域名格式规范。如果是域名，则域名每段（点号分隔）长度小于等于62，域名总长度小于等于128。"}
  DefaultOriginHostHeader *string `json:"default-origin-host-header,omitempty" xml:"default-origin-host-header,omitempty"`
  // {"en":"Origin port", "zh_CN":"回源请求端口"}
  OriginPort *string `json:"origin-port,omitempty" xml:"origin-port,omitempty"`
}

func (s AddCdnDomainRequestOriginConfig) String() string {
  return tea.Prettify(s)
}

func (s AddCdnDomainRequestOriginConfig) GoString() string {
  return s.String()
}

func (s *AddCdnDomainRequestOriginConfig) SetOriginIps(v string) *AddCdnDomainRequestOriginConfig {
  s.OriginIps = &v
  return s
}

func (s *AddCdnDomainRequestOriginConfig) SetDefaultOriginHostHeader(v string) *AddCdnDomainRequestOriginConfig {
  s.DefaultOriginHostHeader = &v
  return s
}

func (s *AddCdnDomainRequestOriginConfig) SetOriginPort(v string) *AddCdnDomainRequestOriginConfig {
  s.OriginPort = &v
  return s
}

type AddCdnDomainRequestLiveConfig struct {
  // {"en":"The live push-pull stream type, the optional values are pull and push, pull means pull flow; push means push flow.", "zh_CN":"直播推拉流类型，可选值为pull和push，pull表示拉流；   push表示推流。"}
  StreamType *string `json:"stream-type,omitempty" xml:"stream-type,omitempty"`
  // {"en":"The push-pull domain name is used to set the push-flow domain name corresponding to the live streaming domain name. When the stream-type is pull, at least one of the source IP address and the corresponding push-stream domain name is not empty. When the stream-type is push, Incoming.", "zh_CN":"配套推流域名，用于设置直播拉流域名对应的推流域名，当stream-type为pull时，源站IP和配套推流域名至少一个不为空；当stream-type为push时，无需传入。"}
  OriginPushHost *string `json:"origin-push-host,omitempty" xml:"origin-push-host,omitempty"`
  // {"en":"Source station IP. When the stream-type is pull, at least one of the source station IP and the companion push stream domain name is not empty.
  // 1. If it is a push-pull flow package, fill in 127.0.0.1, and the system will also default to 127.0.0.1.
  // 2. If it is directly returning to the source, fill in the source IP of the source pull stream.
  // ", "zh_CN":"源站IP，当stream-type为pull时，源站IP和配套推流域名至少一个不为空。
  // 1、如果是推拉流配套，则填写127.0.0.1，不传系统也默认为127.0.0.1
  // 2、如果是直接回源拉流，则填写回源拉流的源站IP"}
  LiveConfigOriginIps *string `json:"origin-ips,omitempty" xml:"origin-ips,omitempty"`
}

func (s AddCdnDomainRequestLiveConfig) String() string {
  return tea.Prettify(s)
}

func (s AddCdnDomainRequestLiveConfig) GoString() string {
  return s.String()
}

func (s *AddCdnDomainRequestLiveConfig) SetStreamType(v string) *AddCdnDomainRequestLiveConfig {
  s.StreamType = &v
  return s
}

func (s *AddCdnDomainRequestLiveConfig) SetOriginPushHost(v string) *AddCdnDomainRequestLiveConfig {
  s.OriginPushHost = &v
  return s
}

func (s *AddCdnDomainRequestLiveConfig) SetLiveConfigOriginIps(v string) *AddCdnDomainRequestLiveConfig {
  s.LiveConfigOriginIps = &v
  return s
}

type AddCdnDomainRequestPublishPoints struct     {
  // {"en":"Publish point, support multiple, do not pass the system by default to generate a publishing point uri for slash /", "zh_CN":"发布点，支持多个，不传系统默认生成一条发布点uri为“/”"}
  Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s AddCdnDomainRequestPublishPoints) String() string {
  return tea.Prettify(s)
}

func (s AddCdnDomainRequestPublishPoints) GoString() string {
  return s.String()
}

func (s *AddCdnDomainRequestPublishPoints) SetUri(v string) *AddCdnDomainRequestPublishPoints {
  s.Uri = &v
  return s
}

type AddCdnDomainResponse struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"The URL used to access the domain information, where domain-id is the unique token generated by our cloud platform for the domain name and whose value is a string.", "zh_CN":"响应信用于访问该域名信息的URL，其中domain-id为我司云平台为该域名生成的唯一标示，其值为字符串。"}
  Location *string `json:"location,omitempty" xml:"location,omitempty" require:"true"`
  // {"en":"The name of the service domain automatically generated by the My company, for example: xxxx.cdn30.com", "zh_CN":"由我司自动生成的服务域名名称，例如：xxxx.cdn30.com"}
  Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
  // {"en":"Request result code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s AddCdnDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s AddCdnDomainResponse) GoString() string {
  return s.String()
}

func (s *AddCdnDomainResponse) SetHttpStatus(v int) *AddCdnDomainResponse {
  s.HttpStatus = &v
  return s
}

func (s *AddCdnDomainResponse) SetXCncRequestId(v string) *AddCdnDomainResponse {
  s.XCncRequestId = &v
  return s
}

func (s *AddCdnDomainResponse) SetLocation(v string) *AddCdnDomainResponse {
  s.Location = &v
  return s
}

func (s *AddCdnDomainResponse) SetCname(v string) *AddCdnDomainResponse {
  s.Cname = &v
  return s
}

func (s *AddCdnDomainResponse) SetCode(v string) *AddCdnDomainResponse {
  s.Code = &v
  return s
}

func (s *AddCdnDomainResponse) SetMessage(v string) *AddCdnDomainResponse {
  s.Message = &v
  return s
}

type Paths struct {
}

func (s Paths) String() string {
  return tea.Prettify(s)
}

func (s Paths) GoString() string {
  return s.String()
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


