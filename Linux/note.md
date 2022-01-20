# Linux 命令笔记
## 一些概念
- 没有消息就是好消息
- 一切皆文件
- 一切文件都挂载在 `/` 下


### 常用目录
- /bin  
    + Binary的缩写，这个目录存放着最经常使用的命令
- /boot 
    + 存放启动Linux 时使用的一些核心文件。包括一些链接文件和镜像文件
- /dev
    + Device的缩写，存放外部设备。（访问方式和普通文件一致
- /etc      <div style="float:right;">**import**</div>
    + 存放所有系统管理所需要的配置文件和子目录
- /home      <div style="float:right;">**import**</div>
    + 用户主目录，每个用户有一个自己的目录，一般以用户账号命名
- /lib
    + 存放系统最基本的动态连接共享库（类似Windows的dll
- /lost+found
    + 一般情况下为空，非法关机会有文件
- /media
    + Linux自动识别一些设备，如U盘，光驱等，识别后挂载到这下面
- /mnt
    + 让用户临时挂载别的文件系统到此
- /opt      <div style="float:right;">**import**</div>
    + 主机额外软件摆放的目录
- /proc
    + 虚拟目录，是内存的映射，可以通过访问这个目录获取系统信息
- /root      <div style="float:right;">**import**</div>
    + 系统管理员（超级权限者）的用户目录
- /sbin
    + s是super user 的s。是系统管理员的系统管理程序
- /srv
    + 存放一些服务（守护进程）启动后需要提取的系统管理程序
- /sys
    + Linux2.6版本的内核大变化，内有新的文件系统sysfs
- /tmp
    + 存放临时文件
- /usr      <div style="float:right;">**import**</div>
    + 用户很多应用程序和文件都在这个目录下面
    + Unix Software Resource 的缩写
    + /usr/bin
        * 系统用户使用的应用程序
    + /usr/sbin
        * 超级系统用户使用的应用程序和系统守护进程
    + /usr/src
        * 内核源代码默认的放置目录
- /var      <div style="float:right;">**import**</div>
    + 存放不断扩充的东西，我们习惯将经常修改的文件和目录放置在此。包括各种日志文件
- /run
    + 临时文件系统，存储系统启动以来的信息。重启时目录下的文件应该被删除或清空
- /www （Linux一般没有，搭建服务器会有）
    + 存放服务器网站相关资源

### 常用目录part2
- /etc/sysconfig/network-scripts   
    + 网络配置目录

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

grep                # 查找文件中符合条件的字符串
```

### 网络
```bash
ifconfig            # 查看网络配置
ping                # ping
```

#### 文件属性
##### 前置10个字母
1. d/l
    - `d` 代表其为目录 dir。
    - `-` 代表其为文件
    - `l` 链接文档 link file
    - `b` 装置文件里面可供存储的接口设备（可随机存取装置。如硬盘）
    - `c` 装置文件里面的串行端口设备（一次性读取装置。如键盘，鼠标）
2. 之后三个为一组
    - 为 `r` 可读 `w` 可写 `x` 可执行。若不具有该权限则为 `-`
    - 三组分别为：属主权限，属组权限，其他用户权限

##### 数字
文件数量

##### 所属用户，所属组
谁的，是那个组的

##### 大小，时间，名称
链接文件的名称会看到链接对象

### 账户管理
```bash
useradd                     # 添加用户
    -m              # 自动创建home下的主目录
    -g              # 指定用户所属用户组
    -G              # 指定用户所属附加组
    -s              # 指定登录的shell文件
    -u              # 用户标识。有-o则可以复用其他用户的用户标识

userdel                     # 删除用户
    -m              # 一并删除home下的主目录

usermod                     # 修改用户
    -d              # 修改用户的主目录
    -g              # 指定用户所属用户组
    -G              # 指定用户所属附加组
    -s              # 指定登录的shell文件
    -u              # 用户标识。有-o则可以复用其他用户的用户标识

su                          # 切换用户

exit                        # 退出登录

hostname                    # 查看/修改主机名

passwd                      # 修改密码（默认自己的，可以指定对象)
    -l              # 锁定，冻结目标账户
    -d              # 移除密码，使得无论输入什么密码都不能登录
```
> 一切皆文件。实际上创建用户就是向文件写入记录。 `/etc/passwd`
>
> 口令:用户标识号:组标识号:注释性描述:主目录:登录shell
>
> 每一行都是一个用户
>
> 口令在 `/etc/shadow`，均加密，保证安全性
>
> 登录名:加密口令:最后一次修改时间:最小时间间隔:最大时间间隔:警告时间:不活动时间:失效时间:标志

### 用户组管理
```bash
groupadd                    # 添加用户组
    -g              # 指定组id

groupdel                    # 删除用户组

groupmod                    # 修改用户组
    -g              # 修改组id
    -n              # 修改名称

newgrp                      # 将当前用户切换入新组
# 这较少用
```
> `/etc/group`

### 磁盘管理
```bash
df                          # 列出所有磁盘情况
    -h              # 单位化(可视化)表示
du                          # 查看文件占用空间
    -a              # 查看所有文件
```

```bash
mount                       # 将外部设备挂载到内部实现访问
umount                      # 取消挂载（卸载
    -f              # 强制
```

### 进程管理
```bash
ps                          # 查看进程信息
    -A              # 查看所有
    -au             # 以用户等较为详细的信息罗列
    -x              # 查看后台运行进程的参数
    -ef             # 父进程信息

pstree                      # 进程树显示

kill                        # 结束进程
    -9              # 根据pid
```
- 每一个进程都有父进程
- 前台运行时，关闭即结束。后台运行可持续

### 安装
```bash
rpm -ivh rpm的软件包名
    -i  # 安装
    -t  # 测试安装
    -p  # 显示安装进度
    -f  # 忽略任何错误
    -U  # 升级安装
    -v  # 验证安装

# 卸载 若使用了
# rpm –ivh software-1.2.3-1.i386.rpm
# 安装
# 而当卸载时，则应执行：
rpm –e software。


tar -xzvf soft.tar.gz //一般会生成一个soft目录 
cd soft 
./configure 
make 
make install
make clean
```

```bash
yum ls
yum install
```

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