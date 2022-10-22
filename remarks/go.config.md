# go 环境配置

`go env` 查看环境变量，`go version` 查看版本

go mod 项目管理工具

## 配置

### 修改 go env 配置

```bash
go env -w GO111MODULE=on
```

## go 项目初始化

使用 go mod 作为项目管理工具

```bash
mkdir [projectName]

cd [projectName] && go mod init [projectName]

```

## go 项目运行

```bash
go run [file].go
```
