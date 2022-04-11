#源镜像
FROM golang:latest
#设置工作目录
WORKDIR /JCRandomProxy
#将服务器的go工程代码加入到docker容器中
ADD . /JCRandomProxy
#go构建可执行文件
RUN go mod init JCRandomProxy && export GOPROXY=https://goproxy.cn && go mod tidy && go build .
#暴露端口
EXPOSE 8081
#最终运行docker的命令
ENTRYPOINT  ["./JCRandomProxy"]