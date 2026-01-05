FROM mysql:8.4

# 定义时区构建参数，增加默认值
ARG TZ=Asia/Shanghai

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
  && echo $TZ > /etc/timezone

EXPOSE 3306
