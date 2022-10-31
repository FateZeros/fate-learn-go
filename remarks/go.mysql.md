# GO 操作 mysql 数据库

## Mac 安装 mysql 数据库

### Mac 配置 mysql 的环境变量

```bash
vi .bash_profile

# mysql
export MYSQL_HOME=/usr/local/mysql/bin
export PATH=$PATH:$MYSQL_HOME

source .bash_profile
```

### Mac 重置 mysql 密码

- 1.关闭 mysql 服务器

```bash
sudo /usr/local/mysql/support-files/mysql.server stop
```

2.

```bash
sudo /usr/local/mysql/bin/mysqld_safe --skip-grant-tables
```

3.打开另外一个终端

```bash
mysql -u root -p

UPDATE mysql.user SET authentication_string=PASSWORD('YOUR NEW MYSQL PASSWORD') WHERE User='root';

FLUSH PRIVILEGES;

\q
```

4. 重置密码

```bash
mysql -u root -p

set password = password('YOUR NEW MYSQL PASSWORD');

show databases;
```

### 创建数据库

```bash
// 创建
create database [数据库名]

// 打开数据库
use [数据库名]
```

### 查看 mysql 端口

```bash
show global variables like 'port';
```

[mysql 端口为 0 解决方法](https://www.jianshu.com/p/a2b5a1d4a36a)

## 在 go 使用 mysql

```bash
go get -u github.com/go-sql-driver/mysql
```
