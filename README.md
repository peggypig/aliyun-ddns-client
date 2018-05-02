# aliyun-ddns-client
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;go语言编写的阿里云的DDNS客户端，解决没有静态公网IP的痛点。<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;喜欢折腾的那些谁看到这里就应该知道这个项目的功能了。<br>

### 说明
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;参数类信息详见阿里云云解析官网<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;解析生效时间和TTL时间有关，免费的DNS的TTL为10分钟，也就是说这个DDNS的时效性可能并不高，<br>
但是对于个人来说已经足够完美，嫌弃的同学，可以购买aliyun相关的服务来解决这个问题。<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;一般来说，最差的生效时间为下面设置检测时间加上TTL的10分钟，默认就是15分钟。





### 使用指南
需要【./ddns.conf】 》》》》使用ini配置格式

\[Must\]<br>
AccessKeyId=testAccessKeyId   <br>
AccessKeySecret=testAccessKeySecret<br>
DomainName=niubi.com<br>

\[Optional\]<br>
CycleTime=300 #时间单位是秒   <br>


上面是推荐配置，当然，你可以根据【[https://help.aliyun.com/product/29697.html]】
官方指南的参数配置来设置参数，在\[Optional\]下面。


linux/mac启动：<br>
./aliyun-ddns-client -ls -start 2>>/var/log/ddns.log &    <br>
可以将日志输出到/var/log/ddns.log中


### 关于aliyun-ddns-client 二进制文件的使用：
-h  -help --h --help 都是显示帮助信息<br>
-start 开启ddns的功能<br>
-ls 显示配置中域名现有的解析记录<br>


### 关于交叉编译
编译成Linux下可执行文件<br>
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

编译成Linux下可执行文件<br>
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

编译成Linux下可执行文件<br>
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build


详见【[https://javasgl.github.io/go-cross-complie/]】



# BIG PS
使用中如果有问题，邮件：zjhfyq@vip.qq.com<br>
好用的话，你应该知道点击哪里，啊哈哈