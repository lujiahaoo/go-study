

### —使用中间件前先在`config/app.go`文件下按它的范围注册，然后在`routers/router.go`文件下调用中间件组

### —路由方法直接添加在 `routers/routes.go`文件下，根据路由前缀来填

