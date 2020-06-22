## 对thoth的要求

1. 我期望thoth不只是能够获取到markdown格式, 也期望其支持不同的格式
2. 我期望thoth能使我只编写规则代码, 即可获取到相应的内容
3. 我期望thoth有个能有一个良好的代码质量, 和规范的文档
4. 我期望thoth是简单且高效的
5. 我期望thoth能有一个易用的web界面和一个简单直观的cli工具
6. 我期望thoth能先支持单页面, 再支持多页面
7. 我期望thoth能先支持静态页面, 后支持动态页面, 和api调用

## 名词定义

* 目标: 需要爬取的目标页面, 当前仅支持单页面
* 规则: 爬取目标的规则, 比如需要定义目标, 使用css还是正则匹配 
* 文章: 对一篇文章进行匹配(target)


## 规则定义(暂且)
```json5
{
	"target": "https://", // 目标
	"regex": ".*?", // 使用正则匹配
	"css": "css", // 使用css
	"show": "markdown", // html,
	"tag": "", // 内容的target
	"category": "", // 内容的分类
	"selector": "regex", // 使用何种选择器
	"template": "github", // 期望使用内置的模板
}
```

## 支持的爬虫解析

* css
* regex(待支持)

## 支持的网站

* github wiki
```json
{
  "target": "https://github.com/ohmyzsh/ohmyzsh/wiki/Cheatsheet",
  "selector": "#wiki-wrapper"
}
```

* github readme

```json
{
  "target": "https://github.com/ohmyzsh/ohmyzsh",
  "selector": "article"
}
```

## 存储定义

```json
{
  "origin_url": "https://m.38kanshu.cc/191854/mulu.html",
  "book": "总有人打扰我的挂机生活",
  "name": "第1章. 我是一个只喜欢兽耳娘的咸鱼挂哔", 
  "chapters": ["https://m.38kanshu.cc/191854/79113644.html", "https://m.38kanshu.cc/191854/79113644.html"]
}
```
