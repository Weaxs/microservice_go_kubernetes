# microservice_go_kubernetes

以Go+Kubernetes微服务实现

使用Go语言的微服务实现，具体的业务代码参考[《凤凰架构》Kubernetes微服务](https://github.com/fenixsoft/microservice_arch_kubernetes)

## 框架

服务间调用使用 keiex [[文档](https://www.cloudwego.io/docs/kitex) | [源码](https://github.com/cloudwego/kitex) | [示例](https://github.com/cloudwego/kitex-examples)]

网关服务使用 hertz [[文档](https://www.cloudwego.io/docs/hertz) | [源码](https://github.com/cloudwego/hertz) | [示例](https://github.com/cloudwego/hertz-examples)]
