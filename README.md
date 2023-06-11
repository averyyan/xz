# golang 学习项目

从0开始复现rk系列的产品，很喜欢rk的结构设计。
自己复现一遍更好的学习golang
推荐rk 
github: https://github.com/rookie-ninja
主页: rkdev.info

# 和RK相同
用了我自己的理解，boot改为register，entry改为service可以对照参考。
部分代码基于RK实现

# 和RK的不同
未实现监控相关的代码
采用显式注册的方式来控制服务的注册和启动，按照注册顺序启动服务，注册倒序注销服务保障服务的依赖。

## 文件夹说明
core 存放核心服务文件 register service logger cert等核心依赖 
plugin 存放基础服务文件 mongodb jwt等基础依赖
web 存放网络类服务文件 gin grpc 等
user 存放自定义服务 wxwork dingding 等依赖其他服务的延伸服务

### 服务文件夹说明
    config 文件用于对配置文件解析
    service 文件用于启动服务和配置
    service子文件夹 用于服务快捷延伸 参考mongo