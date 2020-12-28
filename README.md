## Captcha 图形验证码生成

生成图片示例
![](./captcha.gif)


``` golang
builder := captcha.NewBuilder(captcha.DefaultConfig)

file, letters, err := builder.Build()
if err != nil {
  return
}
```