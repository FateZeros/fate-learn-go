# 项目初始化

## 初始化

使用 [cobra](https://github.com/spf13/cobra) 初始化项目.

安装 `cobra cli`

```bash
go install github.com/spf13/cobra-cli@latest
```

```bash
mkdir [projectName] && cd [projectName]

go mod init [projectName]

// cobra init
cobra-cli init
```

打包

```bash
go build -o [projectName]
```

## 参考

[create-cli-app-with-cobra](https://www.qikqiak.com/post/create-cli-app-with-cobra/) </br>
