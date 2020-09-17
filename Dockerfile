FROM centos

# By 2020/09/10 LiWang

# 作者信息
LABEL MAINTAINER 'liwang <2859413527@QQ.COM>'

# 环境变量
ENV ListenRoute 4u6385IP
ENV LimitTime 60
ENV LimitCount 60
ENV MongoOn 0
ENV MongoHost 127.0.0.1
ENV MongoAuthDB ''
ENV MongoUser ''
ENV MongoPass ''

WORKDIR /

# 安装Golang 工具（正儿八经应该是在本地编译好后，使用COPY 或者 ADD 进容器中，这里为了省事，在容器里面进行编译）
RUN yum install golang git -y

# 克隆地址，然后进行编译
RUN set -x \
        && git clone https://github.com/LoongLiWang/DockerFIle_ReturnOutIP.git \
        && cd DockerFIle_ReturnOutIP \
        && go get gopkg.in/mgo.v2 \
        && go build

# 切换工作目录
WORKDIR /ReturnOutIP

# 对外开放端口
EXPOSE 93

# 执行命令
CMD  ./ReturnOutIP "-ListenRoute" $ListenRoute "-LimitTime" $LimitTime "-LimitCount" $LimitCount "-MongoOn" $MongoOn "-MongoHost" $MongoHost "-MongoAuthDB" $MongoAuthDB "-MongoUser" $MongoUser "-MongoPass" $MongoPass
