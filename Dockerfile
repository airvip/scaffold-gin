# 依赖最新版 alpine 镜像，5M 左右
FROM alpine:latest

# 在容器根目录创建 app 目录
WORKDIR /app

# 挂载容器目录
VOLUME ["/app/etc","/app/log"]

# 拷贝当前目录下 go-app 可执行文件
COPY go-app /app/go-app

# 拷贝当前目录下配置文件
COPY etc /app/etc

# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

# 设置编码
ENV LANG C.UTF-8

# 暴漏端口
EXPOSE 9999

# 运行 golang 程序命令
# ENTRYPOINT [ "pwd" ]
# ENTRYPOINT [ "ls", "-al" ]
# ENTRYPOINT [ "ls", "-al", "etc/" ]
ENTRYPOINT ["./go-app"] 