功能說明:

由手動Run pipeline輸入不同的變量觸發相關功能
1.創建域名:

createdomainname //輸入欲創建之域名,可支持複數用","分開

originset  //回源cname|ip設置,若不同回源要分次跑

2.獲取域名id: //啟用&禁用&刪除域名皆須取得域名id才有辦法操作

setdomainname //需取得之域名

3.禁用域名:

setdomainname //需禁用之域名

disable == "true"

4.啟用域名: //非創建域名,功能是起用被禁用之域名

setdomainname //需啟用之域名

enable == "true"

5.刪除域名: //刪除域名需要域名狀態為禁用

setdomainname //需刪除之域名

delete == "true"

6.取得證書id:

certname
