## GO 语法问题

> addressable(CanAddr) （可以进行想象，进行值复制的是无法address的）
1. 可寻址 x, *x, slice[1], x.age, 可寻址array[1](什么意思呢？ 因为array在go中就是一个值对象，和变量是一回事，其中的元素都是副本，因此需要对其取地址才是可寻址的)
2. 不可寻址 string[0], map[x], literal值, 方法method（而非函数）,const(包括零值)

> CanSet 只加了一个限制 struct 的私有变量无法set

> 反射

1. static/concrete type(类型断言是否成功)
2. interface = [type -> concrete真实类型, value -> 实际对象]
 ```` GO  
   Value -> interface{} 
   value设置值，保证可取地址 CanSet
   SetXXX
   Elem()    对于指针来说 v.Elem() == reflect.Indirect(v), reflect.Indirect(v)其他情况都返回原值
             对于interface{} 返回接口动态值
             其他panic
   
   
   
   var r io.Reader
   tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
   
   // r(value tty, type Os.File)
   var w io.Writer
   w = r.(io.Writer) // 断言做了什么呢？ 判断concrete 中是否有断言中的方法集合
   interface{} 
 ````
   
   
Type
1. 可以获取结构体字段方法
2. 