## もしcharsetがutf-8じゃない場合
```gotemplate
go get golang.org/x/text
```

### 使い方
```gotemplate
utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
```

## 自動でcharsetを判断するライブラリ

```gotemplate
go get golang.org/x/net/html
```
