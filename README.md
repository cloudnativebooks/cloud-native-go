# 《Go专家编程》随书代码

## 使用方式
本书在编写过程中，为了保证正确性，往往会编写一些示例代码进行验证，书中出现的绝大部分示例代码均可以本仓库找到。

读者在阅读本书过程中，可以亲自运行本示例代码。

#### 运行所有代码
```go
E:\RainbowMango\github.com\cloudnativebooks\cloud-native-go>go test ./...
ok      _/E_/RainbowMango/github.com/cloudnativebooks/cloud-native-go/chan      0.137s
ok      _/E_/RainbowMango/github.com/cloudnativebooks/cloud-native-go/defer     0.122s
ok      _/E_/RainbowMango/github.com/cloudnativebooks/cloud-native-go/errors    0.145s
ok      _/E_/RainbowMango/github.com/cloudnativebooks/cloud-native-go/iota      0.161s
ok      _/E_/RainbowMango/github.com/cloudnativebooks/cloud-native-go/map       0.231s
ok      _/E_/RainbowMango/github.com/cloudnativebooks/cloud-native-go/panic     0.168s
ok      _/E_/RainbowMango/github.com/cloudnativebooks/cloud-native-go/range     0.152s
ok      _/E_/RainbowMango/github.com/cloudnativebooks/cloud-native-go/recover   0.203s
ok      _/E_/RainbowMango/github.com/cloudnativebooks/cloud-native-go/reflection        0.237s
ok      _/E_/RainbowMango/github.com/cloudnativebooks/cloud-native-go/select    0.273s
ok      _/E_/RainbowMango/github.com/cloudnativebooks/cloud-native-go/slice     0.297s
ok      _/E_/RainbowMango/github.com/cloudnativebooks/cloud-native-go/string    0.146s
ok      _/E_/RainbowMango/github.com/cloudnativebooks/cloud-native-go/struct    0.122s
ok      _/E_/RainbowMango/github.com/cloudnativebooks/cloud-native-go/sugar     0.113s
ok      _/E_/RainbowMango/github.com/cloudnativebooks/cloud-native-go/trap/loop 0.085s
ok      _/E_/RainbowMango/github.com/cloudnativebooks/cloud-native-go/trap/loop2        1.067s
```

#### 运行特定包的代码
```go
E:\RainbowMango\github.com\cloudnativebooks\cloud-native-go>go test ./chan/...
ok      _/E_/RainbowMango/github.com/cloudnativebooks/cloud-native-go/chan      0.067s
```

#### 运行特定的用例
```go
E:\RainbowMango\github.com\cloudnativebooks\cloud-native-go>go test ./chan/... -run=ExampleWorker
ok      _/E_/RainbowMango/github.com/cloudnativebooks/cloud-native-go/chan      0.070s
```

## 勘误
由于编者水平水限，读者在阅读本书过程中遇到的任何问题，都欢迎[提交Issue](https://github.com/cloudnativebooks/cloud-native-go/issues)。


## 欢迎关注华为云原生微信公众号
![](./qrcode.png)
