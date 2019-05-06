# Deprecated , 请访问新项目 [zjxpcyc/encoder](https://github.com/zjxpcyc/gen)

# Gen 简易数据加解密与转解码集合
Golang 语言本身间接或者直接提供了很多加解密或者转解码的 API, 本库是做了一些我常用的整合 
暂时本库只是作为一个辅助工具集使用，因此并不具备复杂的扩展功能，所有接口以快速使用为主。
基本所有的参数都是以 `string` 类型为主体, 并没有采用 `[]byte` .


其中 `CBC7Decrypt` 函数是为微信小程序做准备
> CBC7Decrypt 方法是错误的, 请不要使用

## 安装与使用
```golang
go get -u "github.com/zjxpcyc/gen"
```

```golang
import "github.com/zjxpcyc/gen"

// md5
result := gen.MD5("xxxxxx")

// sha1
result := gen.SHA1("xxxxxx")

// sha256
result := gen.SHA256("xxxxxx")

// hamc sha256
result := gen.HmacSHA256("xxxxxx", "this is key")

// base64 转码
// 注意, 入参是 []byte
result := gen.Base64([]byte("xxxxxx"))

// base64 解码
// 注意, 返回的是 []byte
result, err := gen.Base64Decode("xxxxxxxxx")

// xml 与 map 互转
// 只支持最简单的有一级子节点嵌套的xml转换
// 可能转出的 xml 会有 cdata, 不是你想要的, 这种情况请参考 xml.go 文件自己修改
xml, err := gen.Map2XML(map[string]string{xxx})
mp, err := gen.XML2Map([]byte(`<xml><foo>foo</foo><bar><![CDATA[bar]]></bar></xml>`))
```