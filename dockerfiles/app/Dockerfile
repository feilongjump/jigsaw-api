FROM golang:alpine AS builder

# 设置必要的环境变量
ENV GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPATH=/go \
    GOCACHE=/go-build

# 移动到工作目录：/opt/app
WORKDIR /opt/app

# 此步骤相当于进行缓存，在 go.mod 以及 go.sum 文件没有进行修改时，可以直接从缓存中获取
# 复制 go.mod 和 go.sum 并下载依赖
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go mod download

# 将代码复制到容器中
COPY . .

# 将代码编译成二进制可执行文件 jigsaw
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go build -o jigsaw .


########## 创建一个小镜像 ########
FROM alpine:latest

# 构建环境时 - 环境变量
ARG TZ
ARG PORT

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 安装基础依赖
RUN apk add --no-cache tzdata

# 设置时区
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 移动到工作目录：/opt/app
WORKDIR /opt/app

# 从 builder 镜像中进行复制
COPY --from=builder /opt/app/.jigsaw.toml ./.jigsaw.toml
COPY --from=builder /opt/app/jigsaw ./jigsaw

RUN chmod +x jigsaw

# 声明服务端口
EXPOSE $PORT

# 启动容器时运行的命令
CMD ["./jigsaw", "--env=production"]
