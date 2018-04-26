# aliyun-ddns-client
go语言编写的阿里云的DDNS客户端，解决没有静态公网IP的痛点


### 槽点
代码质量不是很高，后期重构一下

### 使用指南
需要ddns.conf 》》》》使用ini配置格式

\[AccessKey\]
AccessKeyId=testAccessKeyId
AccessKeySecret=testAccessKeySecret
\[Local\]
CycleTime=300
DomainName=niubi.com

上面是推荐配置，当然，你可以根据【[https://help.aliyun.com/product/29697.html]】
官方指南的参数配置来设置参数，在\[Param\]下面。


### 关于aliyun-ddns-client 二进制文件的使用：
-h  -help --h --help 都是显示帮助信息
-start 开启ddns的功能
-ls 显示配置中域名现有的解析记录



