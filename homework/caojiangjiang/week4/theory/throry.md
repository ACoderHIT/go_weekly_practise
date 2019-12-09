
曾经学习python的时候，记得书上说 dict 是 python的 horsepower（动力）。然后,Slice 和 Map 又何尝不是 golang 的 workhorse 呢？
Array 是值类型，Slice 和 Map 是引用类型。他们是有很大区别的，尤其是在参数传递的时候。
另外，Slice 和 Map 的变量 仅仅声明是不行的，必须还要分配空间（也就是初始化，initialization） 才可以使用。
第三，Slice 和 Map 这些引用变量 的 内存分配，不需要你操心，因为 golang 是存在 gc 机制的（垃圾回收机制）
Array 的用法
数组的声明（这里就是定义，给数据存储分配了空间）：
var arrayName [arraySize] dataType

如果数组定义好之后， 没有给数组元素指定值，那么所有元素被自动初始化为零值。

数组的初始化
var a = [10]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //定义数组的时候，直接初始化

var b = [10]int {1, 2, 3, 4}  //部分元素初始化， 其余元素零值

var c = [...]int {1, 2, 3, 4, 5}   //由初始化列表决定数组长度，不可省去标识符 "..."，否则将变成切片Slice

var d = [10]{2:4, 5:7}   //可以按照下标来进行初始化
数组的访问，可以直接按照下标进行访问
数组的遍历：
package main
import(
    "fmt"
)
func main() {
    var f = [20]int {1, 1}
    for i := 2; i < 20; i++ {
        f[i] = f[i-1] + f[i-2]
    }
    for i := 0; i < 20; i++ {   //采用下标进行遍历
        if i % 5 == 0 {
            fmt.Printf("\n")
        }
        fmt.Printf("f[%2d] = %4d",i , f[i])
    }
}
也可以采用 range 关键字进行遍历：

func main() {
    var f = [20]int {1, 1}
    for i := 2; i < 20; i++ {
        f[i] = f[i-1] + f[i -2]
    }
    for i , v := range f {   //采用 range 关键字 进行遍历
        fmt.Printf("f[%2d] = %4d", i, v)
    }
}
多维数组
var a [3][4]int
初始化
var a = [3][4]int {{1,2}, {1,2,3,4}, {2,3, 4}}
多维数组遍历
/*找到二维数组中的最大元素*/
     package main
     import "fmt"
    func main() {
        var i, j, row, col, max int
        var a = [3][4]int {{1, 3, 7, 3}, {2, 3, 7 , 9}, {22, 3, 5, 10}}
        max = a[0][0]
        for i := 0; i < = 2; i ++ {
                for j := 0; j <= 3; j++ {
                        if a[i][j] > max {
                            max = a[i][j]
                            row = i
                            col = j
                        }
                }
        }
        fmt.Println("max = %d, row = %d, col = %d\n", max, row, col)
    }
Slice 的用法
Slice 的声明（没有分配内存）
 `var s1 []int`
在创建切片的时候，不要指定切片的长度。（否则就成了数组）
切片的类型可以是Go 语言的任何基本数据类型（也包括 引用类型和 Struct 类型）
当一个切片被声明之后，没有初始化的时候，这个 s1 默认的值是 nil。切片的长度是0。可以使用内建函数 len() 获得切片的长度，使用内建函数 cap() 获得切片的容量。
Slice 的创建 (分配了内存)
三种创建方式： 基于底层数组创建，直接创建，或者 make() 函数创建

基于底层数组创建 slice
var slice1 []int   //声明但是不分配空间
slice1 = array[start:end]  //这里不包含 end
slice2 := array[:]         // 引用全部的元素
slice3 := array[0:len(array)]
var slice4 []int
sliec34 = array    //引用全部的元素
直接创建 slice
在声明的时候，直接初始化。

var slice1 = []int {1 ,2, 3, 4, 5}
make() 函数创建 slice
var slice1 = make([]int, 5)  //长度和容量都是 5
var slice2 = make([]int, 5, 10)  //容量是5.
Slice 的 访问和遍历
采用下标进行访问，采用 range 进行遍历。

packge main
import "fmt"
func main() {
    var slice1 = []int {1, 2,3 , 4, 5}
    //使用下标访问 slice
    for i := 0; i <=4; i++ {
        fmt.Println("slice[%d] = %d", i, slice[i])
    }
    fmt.Println()
    //使用range 进行遍历
    for i, v := range slice {
        fmt.Println("slice[%d] = %d", i, v)
    }
 }
Slice 的操作
Slice 中的切片的元素，可以动态的添加和删除，所以操作起来要比数组更加方便。

切片元素的增加
采用内建函数 append() 向切片尾部，增加新的元素， 这些元素保存到底层的数组。

append() 并不会影响原来的切片的属性，（原来切片的长度和cap）
append() 将会返回更新后的切片的对象。
append() 是个变参函数，可以一次性添加多个对象。
append() 添加元素的个数超过 切片的 cap() 的时候，那么底层会 重新分配一个 “足够大” 的内存，一般来说是将原来的内存空间扩大二倍，然后将数据复制到新的内存中去， 原来的空间会保留 （供原先切片使用）（底层数组变化这个问题，应该关注一下）
举例如下：
package main

import "fmt"

func main() {
    //使用make 创建 切片
    var slice1 = make([]int, 3, 6)
    // 使用 append 添加元素，并且未超出 cap
    slice2 := append(slice1, 1, 2, 3)
    // 使用 append 添加元素，并且超出 cap. 这个时候底层数组会变化，新增加的元素只会添加到新的底层数组，不会覆盖旧的底层数组。
    slice3 := append(slice1, 4, 5, 6, 7)
    slice1[0] = 10
    fmt.Printf("len = %d cap = %d %v\n", len(slice1), cap(slice1), slice1)
    fmt.Printf("len = %d cap = %d %v\n", len(slice2), cap(slice2), slice2)
    fmt.Printf("len = %d cap = %d %v\n", len(slice3), cap(slice3), slice3)

}

程序输出是：
len = 3 cap = 6 [10 0 0]
len = 6 cap = 6 [10 0 0 1 2 3]     // 这里的[1, 2, 3] 没有被 [4, 5, 6]覆盖
len = 7 cap = 12 [0 0 0 4 5 6 7]   //这里第一个元素没有变成10，并且容量变成原来的2倍。
切片元素的复制
使用切片长时间引用超大的底层数组，会导致严重的内存浪费现象。 可以新建一个小的slice 对象，然后将所需要的数据复制过去，这样子就不会引用底层数组，直接拷贝了数据，这就是需求。函数 copy()可以 在切片之间复制元素。

copy() 可以复制的元素数量取决于 复制方 和 被复制方的最小长度。
同一个底层数组之间的 元素复制，会导致元素重叠问题。
package main

import "fmt"

func main() {
    var slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    var slice2 = make([]int, 3, 5)
    var n int
    n = copy(slice2, slice1) // just copy three elements
    fmt.Println(n, slice2, len(slice2), cap(slice2))

    slice3 := slice1[3:6]   //二者引用同一个底层数组
    n = copy(slice3, slice1[1:5])  //所以，copy的时候发生元素重叠
    fmt.Println(n, slice1, slice3)

}
程序输出为：
3 [1 2 3] 3 5
3 [1 2 3 2 3 4 7 8 9 10] [2 3 4]   //可以看到元素重叠
Map 的用法
map 存储的是 键值对(key-value)。是一个无序的数据的集合，通过键来进行索引得到对应的值。 这种方式可以加快查找速度。Map 通常称为 字典（dictionary） 或者哈希表(Hash table)。Map 现在是很多语言的标配。

字典的声明
字典名称，“键”类型， “值”类型
var mapName map[keyType]valueType
注意：

不需要给字典指定长度，字典的长度会在初始化或者创建的过程中动态增长
Key 必须是能支持 比较运算符（==, !=）的数据类型，比如 整数，浮点数，指针，数组，结构体，接口等。 而不能是 函数，字典，切片这些类型。
Value 类型 可以是Go语言的任何基本数据类型。
var map1 map[string]int
字典的初始化 和 创建
字典 声明好之后，必须经过初始化或者创建 才能使用。未初始化或者创建的字典为 nil
可以使用“{}”来在声明的时候进行初始化。
可是使用 make()来创建字典。
创建或者初始化之后，就可以使用 “=”操作符来动态的向字典中添加数据项了。
下面使用方式错误，编译不通过：
var map1 map[string]int
map1["key1"] = 2   //编译不通过，字典没有初始化或者创建
下面使用方式正确

var map1 map[string]int {}  //字典的初始化
map1["key1"] = 1

var map2 map[string]int
map2 = make(map[string]int)  //字典的创建
map2["key2"] = 2    //使用 等号 添加数据项
字典元素的查找
v, OK := mapName[Key]    //注意这里是 := 
如果Key存在，将Key对应的Value赋值给v，OK== true. 否则 v 是0，OK==false.

package main

import "fmt"

func main() {
    var map1 = map[string]int{"key1": 100, "key2": 200}
    //
    v, OK := map1["key1"]
    if OK {
        fmt.Println(v, OK)
    } else {
        fmt.Println(v)
    }
    // 这里 不是 :=，是 = ，因为这些变量已经定义过了。
    v, OK = map1["key3"]
    if OK {
        fmt.Println(v, OK)
    } else {
        fmt.Println(v)
    }

}
输出为：
100 true
0
字典项的删除
go 提供了内置函数 delete() 来删除容器内的元素。

delete(map1, "key1")
如果key1值不存在，那么调用将什么也不发生，也不会产生副作用。 但是，如果传入的map 是一个 nil，那么将导致程序出现异常，这一点在写程序的时候特别注意。

package main

import (
    "fmt"
)

func main() {
    var map1 = map[string]int{"key1": 100, "key2": 200, "key3": 300}
    for k, v := range map1 {
        fmt.Println(k, v)
        if k == "key2" {
            delete(map1, k)
        }
        if k == "key3" {
            map1["key4"] = 400
        }
    }

    fmt.Println(map1)
}
程序输出：
key2 200
key3 300
key1 100
map[key1:100 key4:400 key3:300]  //可以看到 map 是无序的。

https://studygolang.com/articles/2685