
# 后端部分

> 基于go+gin+gorm+mysql

## 运行方式
> [!NOTE]\
> 请保证运行路径下有`assets/`和`config/`

`go mod tidy`

`go build main.go`

`./main`



## 数据库
基于mysql，请修改`config/config.yaml`里的账号密码为你的数据库账号密码。

在MySQL中执行`assets/Main.sql`，这可以帮助你建表和导入基本数据。


## 致谢
[phigros-b19-go]()
