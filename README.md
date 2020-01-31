## ePing

由于各个平台的ping命令不支持带协议的host，往往复制的url还需要单独提取host来使用。

ePing是基于Golang写的拓展ping工具，大致流程：**获取参数 --> 解析url --> 依据平台设置次数 --> 执行命令 --> 实时输出**。 

```
[Sewell]: ~/Documents/ePing ✗ master*
➜  ./eping https://baidu.com:443/urlpath/index.html
PING baidu.com (39.156.69.79): 56 data bytes
64 bytes from 39.156.69.79: icmp_seq=0 ttl=50 time=165.996 ms
64 bytes from 39.156.69.79: icmp_seq=1 ttl=50 time=79.287 ms
64 bytes from 39.156.69.79: icmp_seq=2 ttl=50 time=69.949 ms
64 bytes from 39.156.69.79: icmp_seq=3 ttl=50 time=54.314 ms

--- baidu.com ping statistics ---
4 packets transmitted, 4 packets received, 0.0% packet loss
round-trip min/avg/max/stddev = 54.314/92.387/165.996/43.425 ms
```

将编译好的程序移动至环境变量中的程序中，或设置别名`alias ping='/path/eping'`来使用。建议使用前者，命名为eping。