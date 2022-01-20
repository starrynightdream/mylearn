## 命令
### 开关机
```bash
sync                # 将数据同步至硬盘中
shutdown            # 关机命令
    -h      # 指定时间关机
   -h now       # 立刻关机
   -h 10        # 10 分钟之后关机
    -r      # 重启
reboot              # 重启 shutdown -r now
halt                # 关闭系统  shutdown -h now
poweroff            # 关闭系统  shutdown -h now
```

### 文件
#### 基本命令
```bash
ls                  # 查看文件
    -a      # 查看隐藏文件
    -l      # 查看文件细节信息
cd                  # 切换目录
pwd                 # 查看当前所在目录
mkdir               # 创建文件夹
    -p      # 递归创建（多级目录
rmdir               # 删除目录（默认只删除空目录
    -p      # 递归删除
cp                  # 复制目录
    cp from to
rm                  # 移除目录或文件
    -f      # 强制删除
    -r      # 递归
    -i      # 交互模式，每次删除询问
mv                  # 移动。可进行重命名
    -f      # 强制移动
    -u      # 只进行update（只替换更新的文件
```
- `~` 是用户目录

#### 权限命令
```bash
chgrp               # 更改文件属组
    chgrp 属组名 文件名
    -R      # 递归修改
chown               # 更改文件属主
    chown 属主名 文件名
    chown 属主名:属组名 文件名
    -R      # 递归修改
chmod               # 更改文件9个权限
    chomd 777 文件名
    # r:4   w:2   x:1    和为7
```

#### 查看命令
```bash
cat                 # 由第一行开始查看
tac                 # 由最后一行开始查看
nl                  # 输出行号
more                # 一页一页显示
less                # 与more类似，但可以向前翻页，`q`退出
head                # 仅查看前几行
tail                # 仅查看尾几行
```

### 网络
```bash
ifconfig            # 查看网络配置
ping                # ping
```

### 账户管理

### 拓展
#### 硬链接与软连接
硬链接：指向同一个文件，即允许一个文件有多个文件。可以保证误删

软连接：类似快捷方式
```bash
ln                      # 创建链接，默认硬链接
    -s          # 创建软连接
    ln f1 f2        # 格式
touch                   # 创建文件
echo "char" >> f1           # 输出字符串
```

#### vim 
- set nu 设置行号
- set nonu  取消行号