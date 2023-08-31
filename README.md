
 # Tiktok-Simple

![Static Badge](https://img.shields.io/badge/go-b)
![Static Badge](https://img.shields.io/badge/kitex-blue)
![Static Badge](https://img.shields.io/badge/hertz-red?link=https://www.cloudwego.io/zh/docs/hertz/overview/)
![Static Badge](https://img.shields.io/badge/gorm-blue)
![Static Badge](https://img.shields.io/badge/redis-b)
![Static Badge](https://img.shields.io/badge/字节跳动青训营-blue)


> 字节跳动第六届青训营

这是一个基于Go语言，结合`Kitex`、`Gorm`、`Hertz`等框架实现的简易版抖音后台项目。

框架介绍：
- **kitex**：微服务框架
- **Hertz**：Http框架，实现api交互
- **Gorm**：数据库存储
- **Redis**：系统数据缓存

## 一、主要功能
**简易版抖音**的主要功能如下图所示，目前本项目只实现了**基础功能**和**互动方向**的功能。

<div align="center">

![functions.png](pics%2Ffunctions.png)

</div>


## 二、架构设计
本项目由三个层次组成：`http`、`微服务`、`数据持久化`

http层负责响应前端的api请求（采用字节跳动的[hertz框架](https://www.cloudwego.io/zh/docs/hertz/overview/)实现此层）。http层接收到前端传来的api请求后，会通过rpc调用微服务。
本项目的微服务框架使用的是字节跳动的[kitex框架](https://www.cloudwego.io/zh/docs/kitex/getting-started/)。
微服务层保存着业务逻辑代码，并调用数据持久化层为前端提供数据支持。

本项目根据功能分类，初步将服务拆分成了四个微服务：用户、评论、喜欢、视频。

数据持久化层主要使用的是`MySQL`数据库。同时，我们还采用了redis，记录用户的登录状态。

三个层次的调用关系如下：
![call.png](pics%2Fcall.png)

请求处理结束后，会按照箭头的反方向返回。


## 三、目录结构
```shell
tiktok-timple
├── config              存放一些配置文件
├── http                api
    ├─── handlers           api交互函数   
    ├─── rpc                rpc接口，负责调用kitex微服务
    ├─── main.go            hertz/http服务启动入口
├── kitex               微服务
    ├─── kitex_gen          由kitex自动生成的kitex微服务代码
    ├─── *.proto            idl文件
├── pkg                 一些工具文件
├── repository          数据库
    ├── db                  mysql代码
    ├── redis               redis代码
├── service             服务代码（微服务业务逻辑实现）
    ├── comment             “评论”微服务
    ├── like                “喜欢”微服务
    ├── user                “用户”微服务
    ├── video               “视频”微服务
├── readme.md
```

## 四、运行
本项目在windows环境下运行，保姆级启动步骤如下：
### 0、预处理
```shell
go mod tidy
```
在config文件中填写好自己的数据库信息
### 1、启动redis
```shell
.\repository\redis\redis-server\redis-server.exe
```

### 2、启动微服务
新建一个控制窗口，然后输入以下命令。（每一行都需要一个新的控制窗口）
```shell
go run .\service\comment\main.go
go run .\service\like\main.go
go run .\service\user\main.go
go run .\service\comment\main.go
```
### 3、启动http服务
新建一个控制窗口，然后输入以下命令
```shell
go run .\http\main.go
```

## 后记
本项目由两个go后端小白实现，开发时并没有统一好代码风格，所以大家可能会发现两种不一样的代码风格hhhh

由于时间有限，我们仅实现了在`win`下的系统。项目中也存在一些遗憾和设计上的缺陷。欢迎大家友好批评指正！

预计在23年年底，我们会再实现一份更加优质的`linux`版本，预计会加入服务注册与发现、消息队列、docker部署等功能。
