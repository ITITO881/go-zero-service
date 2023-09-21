# go-zero-rpc

基于go zero框架的rpc服务

通过proto文件生成pb文件，指定路径(使用绝对路径)

```
protoc filename.proto --go_out=D:\aaa_project\goprojects\github.com\ITITO881\go-zero-service\rpc\itopb --go-grpc_out=D:\aaa_project\goprojects\github.com\ITITO881\go-zero-service\rpc\itopb
```

本地引用  
生成本地module，添加go.mod文件

```
module github.com/ITITO881/go-zero-rpc
```

在引入项目的go.mod文件中添加 *replace pkg => ./pkg*  
ps:使用相对路径

```
require github.com/ITITO881/go-zero-rpc v0.0.0
replace github.com/ITITO881/go-zero-rpc => ./github.com/ITITO881/go-zero-rpc
```
