FROM centos

ENV TZ Asia/Shanghai

COPY etc /app/etc
COPY transform /app/transform
EXPOSE 8080
WORKDIR /app
CMD ["./transform", "-f", "etc/transform.yaml"]
