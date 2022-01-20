# Nginx
## 概念
- 功能
    + 反向代理
    + 负载均衡
    + 动静分离

## Install
```bash
yum install nginx
```
## command
```bash
nginx

nginx -s stop

nginx -s reload

nginx -s stop
```
## 防火墙
```bash
service firewalld start                         # 开启防火墙
service firewalld restart                       # 重启防火墙
service firewalld stop                          # 关闭防火墙
# 查看防火墙规则
firewall-cmd --list-all
# 查看端口是否开放
firewall-cmd --query-port=8080/tcp
# 开启80端口
firewall-cmd --permanent --add-port=80/tcp
# 移除8080端口
firewall-cmd --permanent --remove-port=8080/tcp

# 重启防火墙（修改配置后重启）
firewall-cmd --reload

# 命令解释
1. firewall-cmd     # Linux操作firewall的工具
2. --permanent      # 表示设置为持久
3. --add-port       # 标识添加的端口，类似同理

```

## note
- `/etc/nginx/conf.d/*.conf` 是找到的nginx的配置文件路径