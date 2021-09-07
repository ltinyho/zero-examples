FROM centos

ENV TZ Asia/Shanghai

COPY api/etc /app/etc
COPY api/shorturl /app/shorturl
EXPOSE 8080
WORKDIR /app
CMD ["./shorturl", "-f", "etc/shorturl-api.yaml"]
