#
## Cannot connect to the Docker daemon at unix:///var/run/docker.sock
```bash
sudo service docker restart
```
基本是因为没有启动docker
## 配置镜像加速
1. 可选的地址
    - 科大镜像：https://docker.mirrors.ustc.edu.cn/
    - 网易：https://hub-mirror.c.163.com/
    - 阿里云：https://<你的ID>.mirror.aliyuncs.com
    - 七牛云加速器：https://reg-mirror.qiniu.com

2.  在/etc/docker/daemon.json内写入
这里使用了七牛云
```
{"registry-mirrors":["https://reg-mirror.qiniu.com/"]}
```

3. 重启服务
```
$ sudo systemctl daemon-reload
$ sudo systemctl restart docker
```

4. 查看效果
```
docker info
# 查看Registry Mirrors:
#     https://reg-mirror.qiniu.com
```

### 顺便记一笔
#### CloundShell vim 退出insert模式
因为浏览器会劫持esc。我的情况是：
`shift + esc` 即可。
网络上有人使用 `shift + tab`