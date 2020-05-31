# yumiao
鱼描，扫描文字，创建你的读书笔记与感想。后端基于Golang的gin框架,实现了带有身份认证(jwt)、权限管理(casbin)等功能的RESTful API。
本项目为后端框架，前端开发中。

## 项目结构
```
├── app
│   ├── casbin
│   │   └── casbin.go
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   └── jwt
│       └── jwt.go
├── conf
│   ├── server.ini
│   └── setting.go
├── controllers
│   ├── auth.go
│   ├── books.go
│   ├── notes.go
│   ├── notespaces.go
│   └── users.go
├── go.mod
├── go.sum
├── main.go
├── middlewares
│   ├── auth.go
│   └── jwt
│       └── jwt.go
├── models
│   ├── Book.go
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
    └── router.go

```
