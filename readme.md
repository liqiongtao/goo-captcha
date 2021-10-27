# 获取

```
rst := captcha.Get(240, 80)
fmt.Println(rst)
```

返回数据结构

```
{
    "id": "",
    "base64image": "",
}
```

# 校验

```
err := captcha.Verify("{id}", "{code}")
if err != nil {
}
```
