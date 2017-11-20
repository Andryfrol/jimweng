## 基本類型
- Go的基本類型有basic types
1. bool
2. string
3. int  int8  int16  int32  int64
4. uint uint8 uint16 uint32 uint64 uintptr
5. byte // uint8 的別名
6. rune // int32 的別名
>     // 代表一個Unicode碼
7. float32 float64
8. complex64 complex128

## Reference website
https://openhome.cc/Gossip/Go/Package.html
https://golang.org/doc/code.html

### GO project overview
1. Go programmers typically keep all their Go code in a single workspace.
> Go 程式編譯者只有一個工作環境
2. A workspace contains many version control repositories (managed by Git, for example).
> 該工作環境可以透過Git管控達到擁有多個專案的版本管控
3. Each repository contains one or more packages.
> 每一個專案都可以涵括一個以上的套件包
4. Each package consists of one or more Go source files in a single directory.
> 每一個套件包內可能都是許多的 .go 檔
5. The path to a package's directory determines its import path.
> 要 import 的套件包必須放在他指定的路徑下