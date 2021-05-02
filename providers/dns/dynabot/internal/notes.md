
https://github.com/joeguo/dropcatch/blob/9097a01e235e382e77ac60c4b53ae3d11cf795b7/dynadot/dynadot.go#L16

https://github.com/jsok/terraform-provider-dynadot/blob/master/dynadot/commands.go

---

https://www.dynadot.com/domain/api3.html#set_dns2

`set_dns2`

Request
```
https://api.dynadot.com/api3.xml?key=mykey&command=set_dns2&domain=domain1.com&main_record_type0=aaaa&main_record0=0:0:0:0:0:0:0:1&main_record_type1=mx&main_record1=mail1.myisp.com&main_recordx1=0&subdomain0=www&sub_record_type0=a&sub_record0=192.168.1.1
```

https://api.dynadot.com/api3.xml?
key=mykey
&command=set_dns2
&domain=domain1.com
&main_record_type0=aaaa
&main_record0=0:0:0:0:0:0:0:1
&main_record_type1=mx
&main_record1=mail1.myisp.com
&main_recordx1=0
&subdomain0=www
&sub_record_type0=a
&sub_record0=192.168.1.1

Response
```xml
<SetDnsResponse><SetDnsHeader><SuccessCode>0</SuccessCode><Status>success</Status></SetDnsHeader></SetDnsResponse>
```

| Set DNS2 Request Parameter                 | Explanation                                                                                                                                                   |
|--------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `domain`                                   | The domain name you want to set, 100 domains can be set per request, make sure they are separated by commas                                                   |
| `main_record_type0` - `main_record_type19` | Main record type, should be "a", "aaaa", "cname", "forward", "txt", "mx", "stealth", "email"                                                                  |
| `main_record0` - `main_record19`           | Specify the DNS record for your domain                                                                                                                        |
| `main_recordx0` - `main_recordx19`         | Mx distance, forward type(301 as "1", 302 as "2"), stealth forward title or email alias, necessary when main_record_type is "forward","mx","stealth","email". |
| `subdomain0` - `subdomain49` (optional)    | Subdomain records (optional)                                                                                                                                  |
| `sub_record_type0` - `sub_record_type49`   | Subdomain record type, should be "a", "aaaa", "cname", "forward", "txt", "srv", "mx", "stealth", "email"                                                      |
| `sub_record0` - `sub_record49`             | Subdomain IP address or target host                                                                                                                           |
| `sub_recordx0` - `sub_recordx49`           | Mx distance, forward type, stealth forward title or email alias, necessary when main_record_type is "forward","mx","stealth","email"                          |
| `ttl` (optional)                           | Time to live                                                                                                                                                  |
| `<SuccessCode></SuccessCode>`              | If the operation is successful, "0" for success, "-1" for failure                                                                                             |
| `<Status></Status>`                        | The status of request                                                                                                                                         |
| `<Error></Error>`                          | Error information about the request, only used when status is "error"                                                                                         |

---

https://www.dynadot.com/domain/api3.html#is_processing

`is_processing`

Request
```
https://api.dynadot.com/api3.xml?key=mykey&command=is_processing
```

Response
```xml
<Response>
  <ResponseHeader>
    <ResponseCode>0</ResponseCode>
    <ResponseMsg>yes</ResponseMsg>
  </ResponseHeader>
</Response>
```

---
