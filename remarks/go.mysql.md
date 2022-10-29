# GO 操作 mysql 数据库

## Mac 配置 mysql 的环境变量

```bash
vi .bash_profile

# mysql
export MYSQL_HOME=/usr/local/mysql/bin
export PATH=$PATH:$MYSQL_HOME

source .bash_profile
```

## Mac 重置 mysql 密码

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
