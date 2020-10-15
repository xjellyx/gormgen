# gormgen

gormgen是一种代码生成工具，可以生成更好的API来查询和更新[gorm](https://gorm.io)结构，而无需处理 interface {}或数据库列名称。

## 如何运行

```go
//go:generate gormgen -structs User,Admin -inputDir . -importPkgs gorm.io/gorm -transformErr true
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Age   int
	Email string
}

type Admin struct {
	gorm.Model
	Name  string `json:"name"`
	Age   int
	Email string
}
```

运行 `go generate` 会生成对应的go文件在example目录下 :
`
gen_admin.go
gen_user.go
`

## 安装
 - `git clone github.com/olongfen/gormgen`
 - `cd cmd/gormgen/ `
 - `go build -o gormgen main.go`
 - mv gormgen to GOPATH/bin/
 - `gormgen -structs User,Admin -inputDir ./example -importPkgs gorm.io/gorm -transformErr true`


## 使用

- `go get -u github.com/olongfen/gormgen/...`
-　在你的模型文件中添加 `//go:generate`,如上代码所示.
- 执行命令即可.

## 功能
- [X] 可以转换数据库错误
- [X] 根据模型生成相对应代码，唯一键和主键会生成相对应的方法
- [X] 可以添加日志

## TODO
- 完成test
- 其他

## 贡献
如果你有什么想法或者建议可以提交代码合并请求

## 参考
- github.com/MohamedBassem/gormgen

