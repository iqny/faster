FROM golang:alpine
# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
RUN go mod download

# 移动到工作目录：cmd/myjob
WORKDIR cmd/myjob

# 将我们的代码编译成二进制可执行文件app
RUN go build -o app .

# 启动容器时运行的命令
CMD ["/build/cmd/myjob/app","-conf=job.toml"]