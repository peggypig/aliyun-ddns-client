# aliyun-ddns-client
go语言编写的阿里云的DDNS客户端，解决没有静态公网IP的痛点。<br>
喜欢折腾的那些谁看到这里就应该知道这个项目的功能了。<br>

### 说明
参数类信息详见阿里云云解析官网


### 槽点
代码质量不是很高，后期重构一下


### 使用指南
需要【./ddns.conf】 》》》》使用ini配置格式

\[AccessKey\]<br>
AccessKeyId=testAccessKeyId   <br>
AccessKeySecret=testAccessKeySecret<br>
\[Local\]<br>
CycleTime=300<br>
DomainName=niubi.com<br>

上面是推荐配置，当然，你可以根据【[https://help.aliyun.com/product/29697.html]】
官方指南的参数配置来设置参数，在\[Param\]下面。


### 关于aliyun-ddns-client 二进制文件的使用：
-h  -help --h --help 都是显示帮助信息<br>
-start 开启ddns的功能<br>
-ls 显示配置中域名现有的解析记录<br>



