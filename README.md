gopinyin Go语言中文转拼音库
===========================

gopinyin说明
-----------

* 支持现有的常用3W多个中文词
* 英文字母保留不变
* 不是中文或英文的字符（如符号）会转成下划线_
* 支持使用首字母模式

gopinyin 用法
-------------
```
go get github.com/lauyoume/gopinyin
```
例子

```go
import(
  "github.com/lauyoume/gopinyin"
  "fmt"
)

func main() {
  fmt.Println(gopinyin.Convert(“Hello, 世界！”, false))
  fmt.Println(gopinyin.Convert(“您好世界”, true))
}

```

