# 红岩寒假考核

## 接口文档地址
https://documenter.getpostman.com/view/24381916/2s935oKifa

## 技术栈
gin，mysql，redis

## 用到的第三方库
1.gin框架
go语言的一个web框架

2.go-redis
Go-redis 是一个类型安全的 Go Redis 客户端库，支持 Pub/Sub、sentinel 和 pipelining 等功能。它是一个能够支持 Redis 集群的 Redis 客户端，旨在通过集群更改自动存储和更新插槽信息。

3.cron
cron一个用于管理定时任务的库，用 Go 实现 Linux 中crontab这个命令的效果。

4.zap
zap是Uber开发的非常快的、结构化的，分日志级别的Go日志库。根据Uber-go Zap的文档，它的性能比类似的结构化日志包更好，也比标准库更快。具体的性能测试可以去github上看到。

5.viper
Viper是适用于Go应用程序的完整配置解决方案。它被设计用于在应用程序中工作，并且可以处理所有类型的配置需求和格式。

6.Gomail
Gomail 是发送电子邮件的简单高效的包，本项目使用其实现发送验证码接口

7.jwt
JSON Web Token (JWT)是一个开放标准(RFC 7519)，它定义了一种紧凑的、自包含的方式，用于作为JSON对象在各方之间安全地传输信息。该信息可以被验证和信任，因为它是数字签名的。