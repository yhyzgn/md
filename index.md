<p align = "center">
<img alt="Lute" src="https://b3log.org/images/brand/lute-128.png">
<br><br>
一款对中文语境优化的 Markdown 引擎，支持 Go 和 JavaScript<br><br>
<em>
千呼万唤始出来 犹抱琵琶半遮面<br>
转轴拨弦三两声 未成曲调先有情
</em>
<br><br>
<a title="Build Status" target="_blank" href="https://travis-ci.org/88250/lute"><img src="https://img.shields.io/travis/88250/lute.svg?style=flat-square"></a>
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/88250/lute"><img src="https://goreportcard.com/badge/github.com/88250/lute?style=flat-square"></a>
<a title="Coverage Status" target="_blank" href="https://coveralls.io/repos/github/88250/lute/badge.svg?branch=master"><img src="https://img.shields.io/coveralls/github/88250/lute.svg?style=flat-square&color=CC9933"></a>
<a title="Code Size" target="_blank" href="https://github.com/88250/lute"><img src="https://img.shields.io/github/languages/code-size/88250/lute.svg?style=flat-square"></a>
<a title="MulanPSL" target="_blank" href="https://github.com/88250/lute/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MulanPSL-orange.svg?style=flat-square"></a>
<br>
<a title="GitHub Commits" target="_blank" href="https://github.com/88250/lute/commits/master"><img src="https://img.shields.io/github/commit-activity/m/88250/lute.svg?style=flat-square"></a>
<a title="Last Commit" target="_blank" href="https://github.com/88250/lute/commits/master"><img src="https://img.shields.io/github/last-commit/88250/lute.svg?style=flat-square&color=FF9900"></a>
<a title="GitHub Pull Requests" target="_blank" href="https://github.com/88250/lute/pulls"><img src="https://img.shields.io/github/issues-pr-closed/88250/lute.svg?style=flat-square&color=FF9966"></a>
<a title="Hits" target="_blank" href="https://github.com/88250/hits"><img src="https://hits.b3log.org/88250/lute.svg"></a>
<br><br>
<a title="GitHub Watchers" target="_blank" href="https://github.com/88250/lute/watchers"><img src="https://img.shields.io/github/watchers/88250/lute.svg?label=Watchers&style=social"></a>  
<a title="GitHub Stars" target="_blank" href="https://github.com/88250/lute/stargazers"><img src="https://img.shields.io/github/stars/88250/lute.svg?label=Stars&style=social"></a>  
<a title="GitHub Forks" target="_blank" href="https://github.com/88250/lute/network/members"><img src="https://img.shields.io/github/forks/88250/lute.svg?label=Forks&style=social"></a>  
<a title="Author GitHub Followers" target="_blank" href="https://github.com/88250"><img src="https://img.shields.io/github/followers/88250.svg?label=Followers&style=social"></a>
</p>

<p align="center">
<a href="https://github.com/88250/lute/blob/master/README_en_US.md">English</a>
</p>

## 💡 简介

[Lute](https://github.com/88250/lute) 是一款结构化的 Markdown 引擎，完整实现了最新的 [GFM](https://github.github.com/gfm/)/[CommonMark](https://commonmark.org) 规范，对中文语境支持更好。

欢迎到 [Lute 官方讨论区](https://ld246.com/tag/lute)了解更多。同时也欢迎关注 B3log 开源社区微信公众号 `B3log开源`：

![b3logos.jpg](https://b3logfile.com/file/2020/08/b3logos-032af045.jpg)

## 📽️ 背景

> sdsdf

之前我一直在使用其他 Markdown 引擎，它们或多或少都有些“瑕疵”：

* 对标准规范的支持不一致
* 对“怪异”文本处理非常耗时，甚至挂死
* 对中文支持不够好

Lute 的目标是构建一个结构化的 Markdown 引擎，实现 GFM/CM 规范并对中文提供更好的支持。所谓的“结构化”指的是从输入的 MD 文本构建抽象语法树，通过操作树来进行 HTML 输出、原文格式化等。
实现规范是为了保证 Markdown 渲染不存在二义性，让同一份 Markdown 文本可以在实现规范的 Markdown 引擎处理后得到一样的结果，这一点非常重要。

实现规范的引擎并不多，我想试试看自己能不能写上一个，这也是 Lute 的动机之一。关于如何实现一个 Markdown 引擎，网上众说纷纭：

* 有的人说 Markdown 适合用正则解析，因为文法规则太简单
* 也有的人说 Markdown 可以用编译原理来处理，正则太难维护

我赞同后者，因为正则确实太难维护而且运行效率较低。最重要的原因是符合 GFM/CM 规范的 Markdown 引擎的核心解析算法不可能用正则写出来，因为规范定义的规则实在是太复杂了。

最后，还有一个很重要的动机就是 B3log 开源社区需要一款自己的 Markdown 引擎：

* [Solo](https://github.com/88250/solo)、[Pipe](https://github.com/88250/pipe)、[Sym](https://github.com/88250/symphony) 需要效果统一的 Markdown 渲染，并且性能非常重要
* [Vditor](https://github.com/Vanessa219/vditor) 需要一款结构化的引擎作为支撑以实现下一代的 Markdown 编辑器

## ✨  特性

* 实现最新版 GFM/CM 规范
* 零正则，非常快
* 内置代码块语法高亮
* 对中文语境支持更好
* 术语拼写修正
* Markdown 格式化
* Emoji 解析
* HTML 转换 Markdown
* 自定义渲染函数
* 支持 JavaScript 端使用

## 📚 数学公式

行内公式： $ \Gamma(z) = \int_0^\infty t^{z-1}e^{-t}dt\,. $

行间公式：
$$
\Gamma(z) = \int_0^\infty t^{z-1}e^{-t}dt\,.
$$

多行公式块：

$$
\frac{1}{
\Bigl(\sqrt{\phi \sqrt{5}}-\phi\Bigr) e^{
\frac25 \pi}} = 1+\frac{e^{-2\pi}} {1+\frac{e^{-4\pi}} {
1+\frac{e^{-6\pi}}
{1+\frac{e^{-8\pi}}{1+\cdots}}
}
}
$$

普通表达式：$(2+3)[4+4]$

如开4次方：$\sqrt[4]{\frac xy}$

连分数：$$x=a_0 + \cfrac {1^2}{a_1 + \cfrac {2^2}{a_2 + \cfrac {3^2}{a_3 + \cfrac {4^2}{a_4 + ...}}}}$$

分类表达式：
$$
f(n)
\begin{cases}
\cfrac n2, &if\ n\ is\ even\\
3n + 1, &if\  n\ is\ odd
\end{cases}
$$
积分：$\int_{r=1}^\infty$

方程组：
$$
\left \{
\begin{array}{c}
a_1x+b_1y+c_1z=d_1 \\
a_2x+b_2y+c_2z=d_2 \\
a_3x+b_3y+c_3z=d_3
\end{array}
\right.
$$
三角函数：

矩阵：
$$
\begin{pmatrix}
1&a_1&a_1^2&\cdots&a_1^n\\
1&a_2&a_2^2&\cdots&a_2^n\\
\vdots&\vdots&\vdots&\ddots&\vdots\\
1&a_m&a_m^2&\cdots&a_m^n\\
\end{pmatrix}
$$

## 🈚️ 流程图

```mermaid
graph TB
    c1-->a2
    subgraph one
    a1-->a2
    end
    subgraph two
    b1-->b2
    end
    subgraph three
    c1-->c2
    end
```

## 🔥 时序图

```mermaid
sequenceDiagram
    Alice->>John: Hello John, how are you?
    loop Every minute
        John-->>Alice: Great!
    end
```

## 😁 甘特图

```mermaid
gantt
    title A Gantt Diagram
    dateFormat  YYYY-MM-DD
    section Section
    A task           :a1, 2019-01-01, 30d
    Another task     :after a1  , 20d
    section Another
    Task in sec      :2019-01-12  , 12d
    another task      : 24d
```

## 📈 图表

```echarts
{
  "title": { "text": "最近 30 天" },
  "tooltip": { "trigger": "axis", "axisPointer": { "lineStyle": { "width": 0 } } },
  "legend": { "data": ["帖子", "用户", "回帖"] },
  "xAxis": [{
      "type": "category",
      "boundaryGap": false,
      "data": ["2019-05-08","2019-05-09","2019-05-10","2019-05-11","2019-05-12","2019-05-13","2019-05-14","2019-05-15","2019-05-16","2019-05-17","2019-05-18","2019-05-19","2019-05-20","2019-05-21","2019-05-22","2019-05-23","2019-05-24","2019-05-25","2019-05-26","2019-05-27","2019-05-28","2019-05-29","2019-05-30","2019-05-31","2019-06-01","2019-06-02","2019-06-03","2019-06-04","2019-06-05","2019-06-06","2019-06-07"],
      "axisTick": { "show": false },
      "axisLine": { "show": false }
  }],
  "yAxis": [{ "type": "value", "axisTick": { "show": false }, "axisLine": { "show": false }, "splitLine": { "lineStyle": { "color": "rgba(0, 0, 0, .38)", "type": "dashed" } } }],
  "series": [
    {
      "name": "帖子", "type": "line", "smooth": true, "itemStyle": { "color": "#d23f31" }, "areaStyle": { "normal": {} }, "z": 3,
      "data": ["18","14","22","9","7","18","10","12","13","16","6","9","15","15","12","15","8","14","9","10","29","22","14","22","9","10","15","9","9","15","0"]
    },
    {
      "name": "用户", "type": "line", "smooth": true, "itemStyle": { "color": "#f1e05a" }, "areaStyle": { "normal": {} }, "z": 2,
      "data": ["31","33","30","23","16","29","23","37","41","29","16","13","39","23","38","136","89","35","22","50","57","47","36","59","14","23","46","44","51","43","0"]
    },
    {
      "name": "回帖", "type": "line", "smooth": true, "itemStyle": { "color": "#4285f4" }, "areaStyle": { "normal": {} }, "z": 1,
      "data": ["35","42","73","15","43","58","55","35","46","87","36","15","44","76","130","73","50","20","21","54","48","73","60","89","26","27","70","63","55","37","0"]
    }
  ]
}
```

## 🎵 五线谱

```abc
X: 24
T: Clouds Thicken
C: Paul Rosen
S: Copyright 2005, Paul Rosen
M: 6/8
L: 1/8
Q: 3/8=116
R: Creepy Jig
K: Em
|:"Em"EEE E2G|"C7"_B2A G2F|"Em"EEE E2G|\
"C7"_B2A "B7"=B3|"Em"EEE E2G|
"C7"_B2A G2F|"Em"GFE "D (Bm7)"F2D|\
1"Em"E3-E3:|2"Em"E3-E2B|:"Em"e2e gfe|
"G"g2ab3|"Em"gfeg2e|"D"fedB2A|"Em"e2e gfe|\
"G"g2ab3|"Em"gfe"D"f2d|"Em"e3-e3:|
```

## 🦟 普通内容

这段内容展示了在内容里面一些排版格式，比如：

- **加粗** - `**加粗**`
- *倾斜* - `*倾斜*`
- ~~删除线~~ - `~~删除线~~`
- `Code 标记` - `` `Code 标记` ``
- [超级链接](https://ld246.com) - `[超级链接](https://ld246.com)`
- [username@gmail.com](mailto:username@gmail.com) - `[username@gmail.com](mailto:username@gmail.com)`

## 👨🏻 提及用户

@Vanessa 通过 `@User` 可以在内容中提及用户，被提及的用户将会收到系统通知。

> NOTE:
>
> 1. @用户名之后需要有一个空格
> 2. 新手没有艾特的功能权限

## 🔨 表格

如果需要展示数据什么的，可以选择使用表格。

| header 1 | header 3 |
| -------- | -------- |
| cell 1   | cell 2   |
| cell 3   | cell 4   |
| cell 5   | cell 6   |

| n    | `value`               | 描述        |
| ---- | ------------------- | ----------- |
| 😃   | 短                  | description |
| x 短 | long long long long | short       |

## 🗃 案例

* [链滴](https://ld246.com)
* [思源笔记](https://github.com/siyuan-note/siyuan)

## 有序列表

10. 1
    1. 11
        1. 111
    2. 12
11. 2

## 无序列表

* 1
    * 11
        * 111
    * 12
* 2

## 🇨🇳 中文语境优化

* 自动链接识别加强
* 在中西文间自动插入空格
* 英文标点替换为中文标点

## ♍ 格式化

格式化功能可将“不整洁”的 `Markdown` 文本格式化为统一风格，在需要公共编辑的场景下，统一的排版风格能让大家更容易协作。

<details>
<summary>点此展开格式化示例。</summary>
<br>
Markdown 原文：

````````markdown
# ATX 标题也有可能需要格式化的 ##
一个简短的段落。

Setext 说实话我不喜欢 Setext 标题
----
0. 有序列表可以从 0 开始
0. 应该自增序号的
1.   对齐对齐对齐

我们再来看看另一个有序列表。
1. 没空行的情况下序号要从 1 开始才能打断段落开始一个新列表
3. 虽然乱序不影响渲染
2. 但是随意写序号容易引起误解

试下贴段代码：
```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, 世界")
}
```
对了，缩进代码块建议换成围栏代码块：

    缩进代码块太隐晦了
    也没法指定编程语言，容易导致代码高亮失效
    所以建议大家用 ``` 围栏代码块
试下围栏代码块匹配场景：
````markdown
围栏代码块只要开头的 ` 和结束的 ` 数量匹配即可，这样可以实现在围栏代码块中显示围栏代码块：
```
这里只有 3 个 `，所以不会匹配markdown代码块结束
```
下面匹配到就真的结束了。
````
以上块级内容都挤在一坨了，插入合理的空行也很有必要。


但是过多的空行分段也不好啊，用来分段的话一个空行就够了。



接下来让我们试试稍微复杂点的场景，比如列表项包含多个段落的情况：
1. 列表项中的第一段

   这里是第二个段落，贴段代码：
   ```markdown
   要成为Markdown程序员并不容易，同理PPT架构师也是。
   注意代码块中的中西文间并没有插入空格。
   ```
   这里是最后一段了。
1. 整个有序列表是“松散”的：列表项内容要用 `<p>` 标签

最后，我们试下对 GFM 的格式化支持：

|col1|col2  |       col3   |
---           |---------------|--
col1 without left pipe      |   this is col2   | col3 without right pipe
                                 ||need align cell|

**以上就是为什么我们需要Markdown Format，而且是带中西文自动空格的格式化。**
````````

格式化后：

````````markdown
# ATX 标题也有可能需要格式化的

一个简短的段落。

## Setext 说实话我不喜欢 Setext 标题

0. 有序列表可以从 0 开始
1. 应该自增序号的
2. 对齐对齐对齐

我们再来看看另一个有序列表。

1. 没空行的情况下序号要从 1 开始才能打断段落开始一个新列表
2. 虽然乱序不影响渲染
3. 但是随意写序号容易引起误解

试下贴段代码：

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, 世界")
}
```

对了，缩进代码块建议换成围栏代码块：

```
缩进代码块太隐晦了
也没法指定编程语言，容易导致代码高亮失效
所以建议大家用 ``` 围栏代码块
```

试下围栏代码块匹配场景：

````markdown
围栏代码块只要开头的 ` 和结束的 ` 数量匹配即可，这样可以实现在围栏代码块中显示围栏代码块：
```
这里只有 3 个 `，所以不会匹配markdown代码块结束
```
下面匹配到就真的结束了。
````

以上块级内容都挤在一坨了，插入合理的空行也很有必要。

但是过多的空行分段也不好啊，用来分段的话一个空行就够了。

接下来让我们试试稍微复杂点的场景，比如列表项包含多个段落的情况：

1. 列表项中的第一段

   这里是第二个段落，贴段代码：

   ```markdown
   要成为Markdown程序员并不容易，同理PPT架构师也是。
   注意代码块中的中西文间并没有插入空格。
   ```

   这里是最后一段了。
2. 整个有序列表是“松散”的：列表项内容要用 `<p>` 标签

最后，我们试下对 GFM 的格式化支持：

| col1                   | col2            | col3                    |
| ---------------------- | --------------- | ----------------------- |
| col1 without left pipe | this is col2    | col3 without right pipe |
|                        | need align cell |                         |

**以上就是为什么我们需要 Markdown Format，而且是带中西文自动空格的格式化。**
````````

</details>

## ✍️ 术语修正

Markdown 原文：

```markdown
在github上做开源项目是一件很开心的事情，请不要把Github拼写成`github`哦！

特别是简历中千万不要出现这样的情况：

> 熟练使用JAVA、Javascript、GIT，对android、ios开发有一定了解，熟练使用Mysql、postgresql数据库。
```

修正后：

```markdown
在 GitHub 上做开源项目是一件很开心的事情，请不要把 GitHub 拼写成`github`哦！

特别是简历中千万不要出现这样的情况：

> 熟练使用 Java、JavaScript、Git，对 Android、iOS 开发有一定了解，熟练使用 MySQL、PostgreSQL 数据库。
```

## ⚡ 性能

请看 [Golang markdown 引擎性能基准测试](https://ld246.com/article/1574570835061)。

## 💪 健壮性

Lute 承载了[链滴](https://ld246.com)上的所有 Markdown 处理，每天处理数百万解析渲染请求，运行表现稳定。

## 🔒 安全性

Lute 没有实现实现 GFM 中的 [Disallowed Raw HTML (extension)](https://github.github.com/gfm/#disallowed-raw-html-extension-)，因为该扩展还是存在一定漏洞（比如没有处理 `<input>` ）。
建议通过其他库（比如 [bluemonday](https://github.com/microcosm-cc/bluemonday)）来进行 HTML 安全过滤，这样也能更好地适配应用场景。

## 🛠️ 使用

有三种方式使用 Lute：

1. 后端：用 Go 语言的话引入 `github.com/88250/lute` 包
2. 后端：将 Lute 启动为一个 HTTP 服务进程供其他进程调用，具体请参考[这里](https://github.com/88250/lute-http)
3. 前端：引入 js 目录下的 lute.min.js，支持 Node.js

### Go

引入 Lute 库：

```shell
go get -u github.com/88250/lute
```

最小化可工作示例：

```go
package main

import (
	"fmt"

	"github.com/88250/lute"
)

func main() {
	luteEngine := lute.New() // 默认已经启用 GFM 支持以及中文语境优化
	html:= luteEngine.MarkdownStr("demo", "**Lute** - A structured markdown engine.")
	fmt.Println(html)
	// <p><strong>Lute</strong> - A structured Markdown engine.</p>
}
```

关于代码块语法高亮：

* 默认使用外部样式表，主题为 github.css，可从 chroma-styles 目录下拷贝该样式文件到项目中引入
* 可通过 `lutenEngine.SetCodeSyntaxHighlightXXX()` 来指定高亮相关参数，比如是否启用内联样式、行号以及主题

### JavaScript

简单示例可参考 JavaScript 目录下的 demo，结合前端编辑器的完整用法请参考 [Vditor 中的示例](https://github.com/Vanessa219/vditor/tree/master/demo)。

![Vditor](https://b3logfile.com/file/2020/02/%E6%88%AA%E5%9B%BE%E4%B8%93%E7%94%A8-ef21ef12.png)

#### 关于 `lute.RenderJSON()` 的使用

```typescript
// JSONRenderer的类型
type JSONRendererType = Array<JSONRendererItemType>

// Flag节点
type FlagType = |
    | "Paragraph"
    | "Emphasis"
    | "Strong"
    | "Blockquote"
    | "ListItem"
    | "Strikethrough"
    | "TableHead"
    | "Table"
    | "TableRow"
    | "Mark"
    | "Sub"
    | "Sup"
    | "Tag"
    | "BlockRef"

// 非Flag节点
type NotFlagType = |
    | "Heading"
    | "ThematicBreak"
    | "List"
    | "HTMLBlock"
    | "InlineHTML"
    | "CodeBlock"
    | "Text"
    | "CodeSpan"
    | "HardBreak"     
    | "SoftBreak"
    | "Link"
    | "Image"
    | "HTMLEntity"
    | "TaskListItemMarker"
    | "TableCell"
    | "EmojiUnicode"
    | "EmojiImg"
    | "MathBlock"
    | "InlineMath"
    | "YamlFrontMatter"
    | "Backslash"
    | "BlockEmbed"
    | "BlockQueryEmbed"

interface JSONRendererItemType {
    type?: string
    value?: string
    flag?: string
    title?: string
    language?: string
    mindmap?: string
    children?: Array<JSONRendererItemType>
}

// 节点分为四类：常规节点、flag节点、链接节点、代码块节点
interface NormalNodeType {
    type: string
    value: string
    children?: Array<JSONRendererItemType>
}

interface FlagNodeType {
    flag: string
    children?: Array<JSONRendererItemType>
}

// 链接或者图片
interface LinkNodeType {
    type: string
    value: string
    title: string
    children?: Array<JSONRendererItemType>
}

interface CodeBlockType {
    type: string
    value: string
    language: string
    mindmap?: string // 如果language为mingmap
}
```

一些细节：

1. lute.js 没有内置语法高亮特性
2. lute.js 编译后大小为 ~2MB，通过 `brotli -o lute.min.js.br lute.min.js` 压缩后大小 ~200KB，常规 GZip 压缩后大小 ~300KB

## 📜 文档

* [《提问的智慧》精读注解版](https://ld246.com/article/1536377163156)
* [CommonMark 规范要点解读](https://ld246.com/article/1566893557720)
* [Lute 实现后记](https://ld246.com/article/1567062979327)
* [Markdown 解析原理详解和 Markdown AST 描述](https://ld246.com/article/1587637426085)

## 🏘️ 社区

* [讨论区](https://ld246.com/tag/lute)
* [报告问题](https://github.com/88250/lute/issues/new/choose)

## 📄 授权

Lute 使用 [木兰宽松许可证, 第2版](http://license.coscl.org.cn/MulanPSL2) 开源协议。

## 🙏 鸣谢

Lute 的诞生离不开以下开源项目，在此对这些项目的贡献者们致敬！

* [commonmark.js](https://github.com/commonmark/commonmark.js)：该项目是 CommonMark 官方参考实现的 JavaScript 版，Lute 参考了其解析器实现部分
* [goldmark](https://github.com/yuin/goldmark)：另一款用 golang 写的 Markdown 引擎，Lute 参考了其树遍历实现部分
* [golang-commonmark](https://gitlab.com/golang-commonmark/markdown)：另一款用 golang 写的 Markdown 引擎，Lute 参考了其 URL 编码以及 HTML 转义算法
* [Chroma](https://github.com/alecthomas/chroma)：用 golang 写的语法高亮引擎
* [中文文案排版指北](https://github.com/sparanoid/chinese-copywriting-guidelines)：统一中文文案、排版的相关用法，降低团队成员之间的沟通成本，增强网站气质
* [GopherJS](https://github.com/gopherjs/gopherjs)：将 Go 代码编译成 JavaScript 代码
