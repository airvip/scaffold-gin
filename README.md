### Swagger-UI

生成 swagger 文件 `swag init`

Swagger-UI Url `http://ip:9999/swagger-ui/index.html` 


编译运行
```
go build -tags netgo -ldflags '-s -w' -o go-app
./go-app
```


### 多环境配置

#### 打包成二进制文件
```
go run .\main.go
go run .\main.go -c="app-dev.yml"
```



### 服务器部署

#### 1. build 项目

```
root@airvip:~/workspace/go/scaffold-gin# go build -o go-app
root@airvip:~/workspace/go/scaffold-gin# ls
Dockerfile  README.md  common  docs  etc  go-app  go.mod  go.sum  internal  log  main.go  util

多了一个 go-app 可执行文件

root@airvip:~# cd ~
root@airvip:~# mkdir go_app
root@airvip:~# mv workspace/go/scaffold-gin/go-app go_app/
root@airvip:~# cd go_app/
root@airvip:~/go_app# ls
go-app
root@airvip:~/go_app# cp -R /root/workspace/go/scaffold-gin/etc .
```

#### 2. 写执行脚本 server.sh

```
root@airvip:~/go_app# vim server.sh
```

内容如下
```
#!/bin/bash
# 设置 gin 为 release 生产模式
export GIN_MODE=release
# 切换到 app 路径下
cd /root/go_app
# 启动
./go-app
```

新增可执行权限
```
root@airvip:~/go_app# chmod +x server.sh
```

#### 3. 创建系统 service

```
因为我的  /etc/systemd/system/ 目录下文件少
root@airvip:~/go_app# cd /etc/systemd/system/
或者
root@airvip:~/go_app# cd /lib/systemd/system
创建文件
root@airvip:/etc/systemd/system# vim go-app.service
```

内容如下
```
[Unit]
Description=go-app

[Service]
Type=simple
Restart=always
RestartSec=3s
ExecStart=/root/go_app/server.sh

[Install]
WantedBy=multi-user.target
```

释义

* Description  对这个服务的描述
* Restart=always 服务异常退出时会重启
* RestartSec=3s  设置重启间隔为3秒
* ExecStart=/root/go_app/server.sh  文件的完整路径这个服务会执行这个文件
* WantedBy=multi-user.target    所有用户都可以执行

#### 运行

```
查看状态
root@airvip:/etc/systemd/system# systemctl status go-app
● go-app.service - go-app
   Loaded: loaded (/etc/systemd/system/go-app.service; disabled; vendor preset: enabled)
   Active: inactive (dead)

启动
root@airvip:/etc/systemd/system# systemctl start go-app

重启
root@airvip:/etc/systemd/system# systemctl restart go-app

停止
root@airvip:/etc/systemd/system# systemctl stop go-app

或者使用下面的命令
service go-app status
service go-app start
service go-app restart
service go-app stop
```

#### 直接运行服务
```
root@airvip:~/go_app# nohup ./go-app&
```


#### 配合nginx 


修改nginx配置文件

```
server {
        listen       80;
        server_name  ginvideo.diff.wang;

        location / {
            proxy_pass http://127.0.0.1:8080/;
        }
}
```
重启nginx


# docker 部署
```
# 打包
root@dyt:~/workspace/go/scaffold-gin# go build -tags netgo -ldflags '-s -w' -o go-app

# 生成镜像
root@dyt:~/workspace/go/scaffold-gin# docker build -t go-app:v1 .
root@dyt:~/workspace/go/scaffold-gin# docker images
REPOSITORY               TAG       IMAGE ID       CREATED         SIZE
go-app                   v1        bd2f7fb24a69   3 minutes ago   31.1MB

#运行镜像
root@dyt:~/workspace/go/scaffold-gin# docker run -d -p 9999:9999 --name gin-demo go-app:v1
2db985002277d0643f290c56f6b95508ef3ec1e04a1f7ae559a1ebe54ad6aca5
root@dyt:~/workspace/go/scaffold-gin# docker ps
CONTAINER ID   IMAGE          COMMAND                  CREATED          STATUS          PORTS                                       NAMES
2db985002277   go-app:v1      "./go-app"               43 seconds ago   Up 43 seconds   0.0.0.0:9999->9999/tcp, :::9999->9999/tcp   gin-demo


#将文件目录挂载出来
root@dyt:~# docker run -d -p 9999:9999 --name gin-demo -v ./conf:/app/etc -v ./log:/app/log go-app:v1
32d9f09cd46fd535dab831218acf723b4d03c893a242da6b2882f760715adc79
```

## 发布与拉取镜像
```

# 推送到镜像仓库   https://cr.console.aliyun.com/cn-beijing/instance/repositories
root@dyt:~# docker login --username=sdq*****@163.com registry.cn-beijing.aliyuncs.com
Password: 
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded

# 使用"docker tag"命令重命名镜像，并将它通过专有网络地址推送至Registry。
root@dyt:~# docker tag 457a308adf4c registry.cn-beijing.aliyuncs.com/airvip/gin-demo:v1
# 使用 "docker push" 命令将该镜像推送至远程。
root@dyt:~# docker push registry.cn-beijing.aliyuncs.com/airvip/gin-demo:v1


# 新服务器拉取镜像
root@airvip:~# docker pull registry.cn-beijing.aliyuncs.com/airvip/gin-demo:v1

```

