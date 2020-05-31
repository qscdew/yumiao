# yumiao
鱼描，扫描文字，创建你的读书笔记与感想。后端基于Golang的gin框架,实现了带有身份认证(jwt)、权限管理(casbin)等功能的RESTful API。
本项目为后端框架，前端开发中。

## 鱼描可以做
* 扫描书籍条码或手动输入信息来创建只属于你的读书笔记区域
* 在该区域内摘录书本内容
* 对每一条内容写下你的心得体会
* 在任何平台修改你的内容
* 一键导出所有笔记

## 项目结构
```
├── app                 //第三方包
│   ├── casbin
│   │   └── casbin.go
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   └── jwt
│       └── jwt.go
├── conf               //配置信息
│   ├── server.ini
│   └── setting.go
├── controllers        //控制器
│   ├── auth.go
│   ├── books.go
│   ├── notes.go
│   ├── notespaces.go
│   └── users.go
├── go.mod
├── go.sum
├── main.go             //程序入口
├── middlewares         //中间件
│   ├── auth.go         //认证相关
│   └── jwt
│       └── jwt.go
├── models              //模型
│   ├── Book.go
│   ├── models.go
│   ├── Note.go
│   ├── NoteSpace.go
│   └── User.go
└── routers
    ├── api
    │   ├── auth.go
    │   ├── books.go
    │   ├── notes.go
    │   ├── notespaces.go
    │   └── users.go
    └── router.go        //路由入口

```
