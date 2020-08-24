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

使用说明：
该工具需与proxypool配合使用，或者将确定可以使用的代理IP写在proxy.lst文件中

配置说明：
参见config.ini

目录说明：
Conf：
- config.go # 配置相关
- config.ini # 配置文件
- proxy.lst # 确定可用的代理
Proxy：
- GetAProxy.go # 获取代理
- GetAProxyA.go # 从proxypool中获取代理
- GetAProxyB.go # 从proxy.lst中获取代理
main.go
