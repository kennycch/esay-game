[TOC]

# 基于GIN的单服务游戏骨架，支持Http和Websocket服务，默认自带Pprof服务。

## 前言：
该骨架基于Gin框架搭建的游戏骨架，由于主要服务与单服务，因此没有Grpc服务。<br/>
数据库基于Redis，已具备本地日志功能。<br/>
另外为了处理请求并发问题，骨架内自带任务分发器对请求进行线性处理，编写业务时几乎无需考虑上锁问题。<br/>

## 框架说明：

1.  建议go版本不要低于1.20且GO111MODULE=on
2.  下载骨架后复制根目录的env.ini.example文件并命名为env.ini，配置参数后即可运行main.go
3.  骨架使用了观察者模式对需要的各种服务进行注册，如果后面需要创建加载更多服务，服务结构体只需实现lifecycle包中的Lifecycle interface并注册到lifecycles中即可
4.  骨架中有自己独特的一套任务分发器，处理任务注册将在game包中的taskMap完成。
5.  任务分配策略分为两种，分别是玩家策略和公共策略。玩家策略将根据玩家ID固定某个Channel进行业务处理从而达到玩家请求线性，公共策略需要有开发人员定义关键词，分配器再根据关键词固定某个Channel进行业务处理从而达到请求线性。
6.  为了保证背包系统统一线性处理，在game/logic/bag包中还配备了AsyncChangeItems方法，如果请求经由公共策略处理且有道具需要进入背包，请全部进背包操作使用AsyncChangeItems方法来确保背包更变的线性。
7.  此骨架假设每个玩家数据就是一条Redis的Hash数据，Hash中的每个字段代表不同的战力系统，因此在设计玩家逻辑包时使用了Hmget来获取本次逻辑处理需要的字段，调用SaveValues时将只更新Hmget的字段来保证性能。
8.  由于proto buf文件为多项目共用文件，因此骨架中并没有包含proto buf源文件目录，使用骨架时请自行构建。
9.  骨架中还缺少了游戏配置目录，由于每个游戏此部分差异较大因此也没有提供相应目录。若游戏配置为json文件，推荐使用github.com/kennycch/gotools/game_config来自动生成游戏配置包。
10. 骨架已包含websocket连接事件和断开事件，需要处理更多业务直接在game/model.go中的connectEvents和disconnectEvents注册即可。
11. 骨架中几乎所有包都有model.go文件，建议包中所有变量、常量以及结构体都编写在此文件中。

## 一、目录结构：
├─tools<br/>
│  ├─logger<br/>
│  ├─net<br/>
│  │  ├─http<br/>
│  │  ├─middleware<br/>
│  │  ├─pprof<br/>
│  │  └─websocket<br/>
│  └─redis<br/>
├─config<br/>
├─game<br/>
│  ├─client<br/>
│  ├─errors<br/>
│  └─logic<br/>
│      ├─bag<br/>
│      └─player<br/>
├─lifecycle<br/>
└─pb<br/>
.gitignore<br/>
env.ini.example<br/>
go.mod<br/>
go.sum<br/>
main.go<br/>
README.md<br/>

### 目录说明：
1.  tools: 工具包，所有需要用到的工具都建议摆放在此目录<br/>
2.  tools.logger：日志包，完成注册后直接调用github.com/kennycch/gotools/log即可<br/>
3.  tools.net：网络服务包<br/>
4.  tools.net.http：Http网络服务包，Http接口路由注册在此包完成<br/>
5.  tools.net.middleware：中间件包<br/>
6.  tools.net.pprof：pprof性能监测包<br/>
7.  tools.net.websocket：websocket路由注册包<br/>
8.  tools.redis：Redis包<br/>
9.  config：程序配置包<br/>
10. game：游戏主业务包<br/>
11. game.client：websocket客户端管理包<br/>
12. game.errors：错误信息管理包<br/>
13. game.logic：业务逻辑编写包，所有业务均按照模块填写在此目录<br/>
14. lifecycle：服务注册器，需要注册的服务全部在main.go中使用此包注册<br/>
15. pb：由proto buf生成的go文件目录<br/>
16. env.ini.example：程序配置示例文件<br/>