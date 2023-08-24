# 公主连结公会战管理

后端使用gin+gorm这俩框架，前端正常三大件

[演示网站](https://43.134.236.220:8080/index)

## 主要功能

对5个阶段的5个boss有分别的展示，并且可以

- [x] 报刀

- [x] 挂树

- [x] 进入

- [x] 进度条显示血量

- [x] 调整boss血量、周目

等功能，协调公会出刀

这基本功能之外，还想做

- [x] 出刀记录，~~方便把摸鱼的揪出来~~

## 部署教程

我这里使用的服务器是 win10 pro 64位

1. 部署mysql

   在官网下载安装，需要版本在  8.0  以上，不然sql执行有问题，还需要设置  root密码为123456

2. 安装mysql连接工具navicat

   连接上数据库后建一个库，名字为  pcr  ，字符集选择  utf8mb4  ，建好后对这个库执行我给的  models  目录下的  pcr.sql  脚本建表造数据

当你把上面的都做完后，就可以双击包里的exe文件启动web服务器了，固定使用的  8080端口  ，访问路径为  “/index”  ，直接访问根路径会报404 ~~因为我忘了写根目录的路由~~

## 演示视频

[演示和部署视频](https://www.bilibili.com/video/BV1Z8411N7h3/)

## 最后

![image-20230104202542779](https://s1.ax1x.com/2023/01/05/pSk4Al4.png)

现在主页大概是这样子:sob:

