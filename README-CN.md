# go-tinypng

[:book: English Documentation](README.md) | :book: 中文文档

------

## 前言
[Tinypng](https://tinypng.com/)是一个提供图片压缩、修改大小的网站，同时提供了Online、API、PS-Plugin等方式供我们使用，API方面官方提供了许多语言的SDK支持，但遗憾的是并没有golang的，而go-tinypng则是一个`golang sdk for tinypng`，更多信息请参考我的一片博客——[Golang + Tinypng实现图片压缩](http://www.jianshu.com/p/ca8827d8110e).

## 支持（Support）

- 图片压缩（Compress）
- 修改图片大小Resize

## 安装（Installation）
安装：
```
    go get -u github.com/gwpp/go-tinypng
```
导入：
```
    import "github.com/gwpp/go-tinypng"
```