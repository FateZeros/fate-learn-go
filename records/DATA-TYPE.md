# GO 数据类型

## 基本类型

```bash
i := 1234    					int
j := int32(1)     		int32
f := float32(3.14)    float32
bytes := [5]byte{'h', 'e', 'l', 'l', 'o'}  [5]byte
primes := [4]int{2, 3, 5, 7}   [4]int
```

## 结构体和指针
GO 语言让程序员决定何时使用指针，数据存储在内存上。

```bash
type Point struct { x, y int }

p := Point{10, 20}

pp := &Point{10, 20}
```

## 字符串
字符串在Go语言内存模型中用一个2字长的数据结构表示。它包含一个指向字符串存储数据的指针和一个长度数据。
因为 string 类型是不可变的。切分操作 `str[i:j]` 会得到一个新的2字长结构，一个可能不同的但仍指向同一个字节序列
的指针和长度数据。这意味着字符串切分可以在不涉及内存分配或复制操作。这使得字符串切分的效率等同于传递下标。


```bash
s := "hello"

t := s[2:3]
```

## slice
一个slice是一个数组某个部分的引用。在内存中，它是一个包含3个域的结构体：指向slice中第一个元素的指针，slice的长度，以及slice的容量



[参考](https://tiancaiamao.gitbooks.io/go-internals/content/zh/02.1.html)