JCRandomProxy(随机代理)

功能：
http代理
动态代理
代理有效性校验

安装：
（一）源码安装
git clone https://github.com/Chroblert/JCRandomProxy.git
go run main.go
(二) 使用二进制文件
git clone https://github.com/Chroblert/JCRandomProxy.git
将适合自己系统的二进制文件拷贝到clone下来的目录下
执行：

使用说明：
该工具需与proxypool配合使用，或者将代理IP写在proxy.lst文件中

命令说明：

目录说明：
Conf：
- config.go
- config.ini
- proxy.lst
Proxy：
- GetAProxy.go
- GetAProxyA.go
- GetAProxyB.go
JCLog:
- JCLog.go
main.go