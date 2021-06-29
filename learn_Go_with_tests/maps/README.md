# map

## what is map?

维基百科里这样定义 map：


> In computer science, an associative array, map, symbol table, or dictionary is an abstract data type composed of a collection of (key, value) pairs, such that each possible key appears at most once in the collection.
>
> 在计算机科学中，关联数组、映射、符号表或字典是一种抽象数据类型，由（键、值）对的集合组成，这样每个可能的键在集合中最多出现一次。

简单来说，map 是由 key-value 对组成的；并且 key 只会出现一次

## map 内存模型

在源码中，表示 map 的结构体是 hmap，它是 hashmap 的“缩写”：

```go
// A header for a Go map.

type hmap struct{

    // 元素个数，调用 len(map) 时，直接返回此值
    count     int

    flags     uint8

    // buckets 的对数 log_2
    B         uint8

    // overflow 的 bucket 近似数
    noverflow uint16
    
    // 计算 key 的哈希的时候会传入哈希函数
    hash0     uint32

    // 指向 buckets 数组，大小为 2^B    
    // 如果元素个数为0，就为 nil
    buckets    unsafe.Pointer

    // 扩容的时候，buckets 长度会是 oldbuckets 的两倍
    oldbuckets    unsafe.Pointer

    // 指示扩容进度，小于此地址的 buckets 迁移完成
    nevacuate  uintptr

    extra *mapextra  // optional fields
}
```

B 是 buckets 数组的长度的对数，也就是说 buckets 数组的长度就是 2^B 

bucket 里面存储了 key 和 value

buckets 是一个指针，最终它指向的是一个结构体：

```go
type bmap struct {
    tophash [bucketCnt]uint8
}
```

但这只是表面(src/runtime/hashmap.go)的结构，编译期间会给它加料，动态地创建一个新的结构：

```go
type bmap struct {
    topbits [8]uint8
    keys    [8]keytype
    values  [8]valuetype
    pad     uintptr
    overflow  uintptr
}
```

bmap 就是我们常说的 "桶" ，桶里面会最多装 8 个 key，这些 key 之所以会落入同一个桶，是因为它们经过哈希计算后，哈希结果是 "一类" 的。在桶内，又会根据 key 计算出来的 hash 值的高 8 位来决定 key 到底落入桶的哪个位置(一个桶内最多有 8 个位置)

![](https://mmbiz.qpic.cn/mmbiz_png/ASQrEXvmx61pib1iaeK6CYYicjtlSl0HrycEvYofWxQWP0fnXSqqfwRFKt8HSJ7HP2qic0mqfEv9w82B0Qvpg1OJNg/640?wx_fmt=png&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

当 map 的 key 和 value 都不是指针，并且 size 都小于 128 字节的情况下，会把 bmap 标记为不含指针，这样可以避免 gc 时扫描整个 hmap。但是，我们看 bmap 其实有一个 overflow 的字段，是指针类型的，破坏了 bmap 不含指针的设想，这时会把 overflow 移动到 extra 字段来。

```go
// mapextra holds fields that are not present on all maps.
type mapextra struct {
	// If both key and elem do not contain pointers and are inline, then we mark bucket
	// type as containing no pointers. This avoids scanning such maps.
	// However, bmap.overflow is a pointer. In order to keep overflow buckets
	// alive, we store pointers to all overflow buckets in hmap.extra.overflow and hmap.extra.oldoverflow.
	// overflow and oldoverflow are only used if key and elem do not contain pointers.
	// overflow contains overflow buckets for hmap.buckets.
	// oldoverflow contains overflow buckets for hmap.oldbuckets.
	// The indirection allows to store a pointer to the slice in hiter.
	overflow    *[]*bmap
	oldoverflow *[]*bmap

	// nextOverflow holds a pointer to a free overflow bucket.
	nextOverflow *bmap
}
```

bmap 是存放 k-v 的地方，以下是 bmap 的内部组成：

![](https://mmbiz.qpic.cn/mmbiz_png/ASQrEXvmx61pib1iaeK6CYYicjtlSl0HrycRIjnUcLIJJSRzDeGXQW7eFbcsfIF69fHIyy8RgHj7f9ibI4pQVUwyHA/640?wx_fmt=png&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

```go
// A bucket for a Go map.
type bmap struct {
	// tophash generally contains the top byte of the hash value
	// for each key in this bucket. If tophash[0] < minTopHash,
	// tophash[0] is a bucket evacuation state instead.
	tophash [bucketCnt]uint8
	// Followed by bucketCnt keys and then bucketCnt elems.
	// NOTE: packing all the keys together and then all the elems together makes the
	// code a bit more complicated than alternating key/elem/key/elem/... but it allows
	// us to eliminate padding which would be needed for, e.g., map[int64]int8.
	// Followed by an overflow pointer.
}
```

从注释中可以看到，bmap 并不是以 key/value/key/value/... 这样的形式放在一起的，而是将 key 和 value 各自放在一起。源码里说明这样的好处是在某些情况下可以省略掉 padding 字段，节省内存空间

例如，有这样一个类型的 map：

`map[int64]int8`

如果按照 `key/value/key/value/...` 这样的模式存储，那在每一个 key/value 对之后都要额外 padding 7 个字节；而将所有的 key，value 分别绑定到一起，这种形式 `key/key/.../value/value/...`，则只需要在最后添加 padding。

每个 bucket 设计成最多只能放 8 个 key-value 对，如果有第 9 个 key-value 落入当前的 bucket，那就需要再构建一个 bucket ，通过 `overflow` 指针连接起来。


## 创建 map

```go
ageMap := make(map[string]int)

// 指定 map 长度
ageMp := make(map[string]int, 8)

// ageMp 为 nil, 不能向其添加元素，会直接 panic
var ageMp map[string]int
```

通过汇编语言可以看到，实际上底层调用的是 makemap 函数，主要做的工作就是初始化 hmap 结构体的各种字段，例如计算 B 的大小，设置哈希种子 hash0 等等。

```go
// makemap implements Go map creation for make(map[k]v, hint).
// If the compiler has determined that the map or the first bucket
// can be created on the stack, h and/or bucket may be non-nil.
// If h != nil, the map can be created directly in h.
// If h.buckets != nil, bucket pointed to can be used as the first bucket.
func makemap(t *maptype, hint int, h *hmap) *hmap {
	mem, overflow := math.MulUintptr(uintptr(hint), t.bucket.size)
	if overflow || mem > maxAlloc {
		hint = 0
	}

	// initialize Hmap
	if h == nil {
		h = new(hmap)
	}
	h.hash0 = fastrand()

	// Find the size parameter B which will hold the requested # of elements.
	// For hint < 0 overLoadFactor returns false since hint < bucketCnt.
	B := uint8(0)
	for overLoadFactor(hint, B) {
		B++
	}
	h.B = B

	// allocate initial hash table
	// if B == 0, the buckets field is allocated lazily later (in mapassign)
	// If hint is large zeroing this memory could take a while.
	if h.B != 0 {
		var nextOverflow *bmap
		h.buckets, nextOverflow = makeBucketArray(t, h.B, nil)
		if nextOverflow != nil {
			h.extra = new(mapextra)
			h.extra.nextOverflow = nextOverflow
		}
	}

	return h
}
```

> 注意，这个函数返回的结果：*hmap，它是一个指针, 而 makeslice 函数返回的是 Slice 结构体

```go
type slice struct {
	array unsafe.Pointer  // 元素指针
	len   int             // 长度
	cap   int             // 容量
}

func makeslice(et *_type, len, cap int) unsafe.Pointer {
	mem, overflow := math.MulUintptr(et.size, uintptr(cap))
	if overflow || mem > maxAlloc || len < 0 || len > cap {
		// NOTE: Produce a 'len out of range' error instead of a
		// 'cap out of range' error when someone does make([]T, bignumber).
		// 'cap out of range' is true too, but since the cap is only being
		// supplied implicitly, saying len is clearer.
		// See golang.org/issue/4085.
		mem, overflow := math.MulUintptr(et.size, uintptr(len))
		if overflow || mem > maxAlloc || len < 0 {
			panicmakeslicelen()
		}
		panicmakeslicecap()
	}

	return mallocgc(mem, et, true)
}
```

makemap 和 makeslice 的区别，带来一个不同点：当 map 和 slice 作为函数参数时，在函数参数内部对 map 的操作会影响 map 自身；而对 slice 却不会。

主要原因： 一个是指针(`*hmap`), 一个是结构体(`slice`)。Go 语言中的函数传参都是值传递，在函数内部，参数会被 copy 到本地。 `*hmap`  指针 copy 完之后，仍然指向同一个 map ，因此函数内部对 map 的操作会影响实参。而 slice 被 copy 后，会成为一个新的 slice， 对它进行的操作不会影响到实参。

## 哈希函数

map 的一个关键点在于，哈希函数的选择。在程序启动时，会检测 cpu 是否支持 aes，如果支持，则使用 aes hash，否则使用 memhash。 这是在函数 alginit() 中完成，位于路径：`src/runtime/alg.go` 下

> hash 函数，有加密型和非加密型。加密型的一般用于加密数据、数字摘要等，典型代表就是 md5、sha1、sha256、aes256 这种；非加密型的一般就是查找。在 map 的应用场景中，用的是查找。
> 
> 选择 hash 函数主要考察的是两点：性能、碰撞概率。

表示类型的结构体

```go
type _type struct {
    size         uintptr
    ptrdata      uintptr   // size of memory prefix holding all pointers
    hash         uint32
    tflag        tflag
    fieldalign   uint8
    kind         uint8
    alg          *typeAlg
    gcdata       *byte
    str          nameOff
    ptrToThis    typeOff
}
```

其中 `alg` 字段就和哈希相关，它是指向如下结构体的指针：

```go
// src/runtime/alg.go
type typeAlg struct {
    // (ptr to object, send) -> hash
	hash func(unsafe.Pointer, uintptr) uintptr

	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
}
```

typeAlg 包含两个函数，hash 函数计算类型的哈希值，而 equal 函数则计算两个类型是否 "哈希相等"。 对于 string 类型，它的 hash 、 equal 函数如下：

```go
func strhash(a unsafe.Pointer, h uintptr) uintptr {
	x := (*stringStruct)(a)
	return memhash(x.str, h, uintptr(x.len))
}

func strequal(p, q unsafe.Pointer) bool {
	return *(string)(p) == *(*string)(q)
}
```

根据 key 的类型，_type 结构体的 alg 字段会被设置对应类型的 hash 和 equal 函数。

## key 定位过程

key 经过哈希计算后得到哈希值，共 64 个 bit 位（64位机，32位机就不讨论了，现在主流都是64位机），计算它到底要落在哪个桶时，只会用到最后 B 个 bit 位。还记得前面提到过的 B 吗？如果 B = 5，那么桶的数量，也就是 buckets 数组的长度是 2^5 = 32。

例如，现在有一个 key 经过哈希函数计算后，得到的哈希结果是：
`
10010111 | 000011110110110010001111001010100010010110010101010 │ 01010`

 - 获取放置桶的编号： 用最后的 5 个 bit 位，也就是 `01010`，值为 10，也就是 10 号桶。这个操作实际上就是取余操作，但是取余开销太大，所以代码实现上用的位操作代替。

 - 获取 key 在桶的链表中的位置： 再用哈希值的高 8 位，找到此 key 在 bucket 中的位置，这是在寻找已有的 key。最开始桶内还没有 key，新加入的 key 会找到第一个空位，放入。

buckets 编号就是桶编号，当两个不同的 key 落在同一个桶中，也就是发生了哈希冲突。冲突的解决手段是用链表法：在 bucket 中，从前往后找到第一个空位。这样，在查找某个 key 时，先找到对应的桶，再去遍历 bucket 中的 key。

![](https://mmbiz.qpic.cn/mmbiz_png/ASQrEXvmx61pib1iaeK6CYYicjtlSl0HrycFFpgwNjSpHLP9sTiaPTrGe9icBPkycO2pbKvibTddsnjk5YrDe0VicwGjA/640?wx_fmt=png&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

上图中，假定 B = 5，所以 bucket 总数就是 2^5 = 32。首先计算出待查找 key 的哈希，使用低 5 位 00110，找到对应的 6 号 bucket，使用高 8 位 `10010111`，对应十进制 151，在 6 号 bucket 中寻找 tophash 值（HOB hash）为 151 的 key，找到了 2 号槽位，这样整个查找过程就结束了。

如果在 bucket 中没找到，并且 overflow 不为空，还要继续去 overflow bucket 中寻找，直到找到或是所有的 key 槽位都找遍了，包括所有的 overflow bucket。

通过汇编语言可以看到，查找某个 key 的底层函数是 `mapacess` 系列函数，函数的作用类似，区别在下一节会讲到。这里我们直接看 `mapacess1` 函数：

```go
// mapaccess1 returns a pointer to h[key].  Never returns nil, instead
// it will return a reference to the zero object for the elem type if
// the key is not in the map.
// NOTE: The returned pointer may keep the whole map live, so don't
// hold onto it for very long.
func mapaccess1(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
	if raceenabled && h != nil {
		callerpc := getcallerpc()
		pc := funcPC(mapaccess1)
		racereadpc(unsafe.Pointer(h), callerpc, pc)
		raceReadObjectPC(t.key, key, callerpc, pc)
	}
	if msanenabled && h != nil {
		msanread(key, t.key.size)
	}
	if h == nil || h.count == 0 {
		if t.hashMightPanic() {
			t.hasher(key, 0) // see issue 23734
		}
		return unsafe.Pointer(&zeroVal[0])
	}
	if h.flags&hashWriting != 0 {
		throw("concurrent map read and map write")
	}
	hash := t.hasher(key, uintptr(h.hash0))
	m := bucketMask(h.B)
	b := (*bmap)(add(h.buckets, (hash&m)*uintptr(t.bucketsize)))
	if c := h.oldbuckets; c != nil {
		if !h.sameSizeGrow() {
			// There used to be half as many buckets; mask down one more power of two.
			m >>= 1
		}
		oldb := (*bmap)(add(c, (hash&m)*uintptr(t.bucketsize)))
		if !evacuated(oldb) {
			b = oldb
		}
	}
	top := tophash(hash)
bucketloop:
	for ; b != nil; b = b.overflow(t) {
		for i := uintptr(0); i < bucketCnt; i++ {
			if b.tophash[i] != top {
				if b.tophash[i] == emptyRest {
					break bucketloop
				}
				continue
			}
			k := add(unsafe.Pointer(b), dataOffset+i*uintptr(t.keysize))
			if t.indirectkey() {
				k = *((*unsafe.Pointer)(k))
			}
			if t.key.equal(key, k) {
				e := add(unsafe.Pointer(b), dataOffset+bucketCnt*uintptr(t.keysize)+i*uintptr(t.elemsize))
				if t.indirectelem() {
					e = *((*unsafe.Pointer)(e))
				}
				return e
			}
		}
	}
	return unsafe.Pointer(&zeroVal[0])
}
```
