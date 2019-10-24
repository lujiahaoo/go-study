

# go-study-api

使用中间件前先在`bootstrap/app.go`文件下按它的范围注册，然后在`app/routers/router.go`文件下对应不同的范围调用中间件组

路由方法直接添加在 `app/routers/api.go`文件下，根据路由前缀来填

- [ ] 基于Docker运行(dockerfile)

- [x] Use go modules to init the project
- [ ] When start, stop or restart the application  makes the console output some logo
- [x] Validation of the request parameter
  - [ ] model validate
- [x] 路由中间件
  - [x] JWT
  - [ ] 日志消息
  - [x] Exception 
  - [ ] Tracing(链路追踪)
- [ ] 自定义消息预警
  - [x] mail (gomail)
  - [ ] wechat
  - [ ] SMS
- [ ] 消息队列(RabbitMQ)
- [x] Mysql读写分离
- [ ] 负载均衡