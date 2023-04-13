docker build . -f Dockerfile.rootless -t artifacts.iflytek.com/docker-private/atp/gitea:1.19.2-ailab-rootless  --build-arg GOPROXY=https://goproxy.cn
