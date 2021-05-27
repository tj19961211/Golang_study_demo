# gRPC实践理解

grpc是一种用于两个进程间通信的方式，是基于HTTP/2协议实现的

## HTTP/2

HTTP/2是一个二进制协议，也就意味着它的可读性几乎为0(是写给机器看的协议)，但可以通过抓包工具对其进行解析

HTTP/2中一些通用术语：

- Stream： 一种双向流，一条链接可以有多个streams

- Message： 也就是逻辑上的request、response

- Frame(帧)： 数据传输的最小单位。每个Frame都属于一个特定的stream或者整个链接。一个message可能有多个frame组成

### Frame Format

Frame 是 HTTP/2 里面最小的数据传输单位，一个 Frame 定义如下（[直接从官网 copy 的](http://httpwg.org/specs/rfc7540.html#rfc.section.4.1)）：

```html
+-----------------------------------------------+
|                 Length (24)                   |
+---------------+---------------+---------------+
|   Type (8)    |   Flags (8)   |
+-+-------------+---------------+-------------------------------+
|R|                 Stream Identifier (31)                      |
+=+=============================================================+
|                   Frame Payload (0...)                      ...
+---------------------------------------------------------------+
```

Length：也就是 Frame 的长度，默认最大长度是 16KB，如果要发送更大的 Frame，需要显示的设置 max frame size。

Type：Frame 的类型，譬如有 DATA，HEADERS，PRIORITY 等。

Flag 和 R：保留位，可以先不管。

Stream Identifier：标识所属的 stream，如果为 0，则表示这个 frame 属于整条连接。Frame 

Payload：根据不同 Type 有不同的格式。
