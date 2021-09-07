
```bash
docker run -d --restart=unless-stopped \
-p 80:80 -p 443:443 \
-v ./rancher:/var/lib/rancher/ \
-v ./log/auditlog:/var/log/auditlog \
-e CATTLE_SYSTEM_CATALOG=bundled \
-e AUDIT_LEVEL=3 \
rancher/rancher:2.5.4
```


生成k8s文件
```bash
 goctl kube deploy  -name shorturl -namespace default -image registry.cn-shanghai.aliyuncs.com/ltinyho-web/shorturl:latest -port 8080 -o shorturl.yml
goctl kube deploy  -name transform -namespace default -image registry.cn-shanghai.aliyuncs.com/ltinyho-web/transform:latest -port 8080 -o transform.yml
```

使用k3d安装k3s集群,配置docker-registry
```bash
https://k3d.io/#install-current-latest-release
```


```bash
curl --insecure -sfL https://192.168.3.100/v3/import/bg686mpv6qgc9ttz8sdw4pxpg6dpn46sh5rdbfsxgrprfqd688nrck_c-5b4vp.yaml | kubectl apply -f -
```
