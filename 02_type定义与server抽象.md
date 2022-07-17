1. http库-Request Body
1）Body:只能读取一次，意味着你读了别人就不能读了
（额外的:流的概念就是从头读到尾，你第二次读的时候就是尾到尾，不会报错，但是也读不到数据）
2）getBody:原则上是可以多次读取，但是原生的http.Request里面，这个是nil
3）在读取到body之后，我们就可以用于反序列化，比如说将json格式转化为一个对象等

2. http库--Request Query
1）除了Body,我们还可能传递参数的地方是Query
2)也就是http://xxx.com/your/path?id=123&b=456
3)所有的值都被解释为字符串，所以需要自己解析为数字等
```go
func queryParams(w http.ResponseWriter, r *http.Request) {
    values := r.URL,Query()
    fmt.Fprintf(w,"query is %v\n",values)
}
```
3. http库--Request RUL
包含路径方面的所有信息和一些很有用的操作

URL里面Host不一定有值
r.Host一般都有值，是Host这个header的值、
RawPath也是不一定有
Path肯定有
Tip:实际中记得自己数出来看一下，确认有没有

4. http库-Request Header
1）header大体上是两类，一类是http预定义的，一类是自己定义的
2）Go会自动将header名字转为标准名字——其实就是大小写调整
3）一般用X开头来表明是自己定义的，比如说X-mycompany-you=header

5. http库--Form
Form和ParseForm
要先调用ParseForm
建议加上Content-Type:application/x-www-form-urlencoded

6. http库使用要点
1）Body和GetBody：重点在于Body是一次性的，而GetBody默认情况下是没有，一般中间件会考虑帮你注入这个方法
2）URL:注意URL里面的字段的含义可能并不是如你期望的那样
3）Form：记得调用前先用ParseForm,别忘记请求里加上http头

7. type定义
type 名字 interface
type 名字 struct
type 名字 别的类型
type 别名=别的类型
结构体初始化
指针与方法接收器
结构体如何实现接口

8. interface定义
1）基本语法:type 名字 interface{}
2）里面只能有方法，方法也不需要func关键字
3）啥是接口：接口就是一组行为的抽象
4）尽量用接口，以实现面向接口编程

Tip:当你怀疑要不要用接口的时候，加上去总是很保险的

9. struct定义
1）基本语法:
type Name struct{
    fieldName FieldType
    //....
}
2）结构体和结构体的字段都遵循大小写控制访问性的原则

10. type A B
基本语法：type TypeA typeB

我一般是，在我使用第三方库有没有办法修改源码的情况下，又想在扩展这个库的结构体的方法，就会用这个
```go
// 1.
type Fish struct{
}
func(f Fish) Swim() {
    fmt.Printf("我是鱼，假装自己是一只鸭子\n")
}

//2.
//定义一个新类型
type FakeFish Fish
func (f FakeFish) FakeSwim() {
    fmt.Printf("我是山寨鱼,嘎嘎嘎\n")
}

fake := FakeFish{}
//fake.Swim() 无法调用原来Fish的方法
fake.FakeSwim()

//转换为Fish
td := Fish(fake)
//真的变成了鱼
td.Swim()

//3. 以上那种没必要那样用
```

11. 结构体初始化
Go没有构造函数
初始化语法：Struct{}
获取指针：&Struct
获取指针2：new(Struct)
new可以理解为Go会为你的变量分配内存，并且把内存都置为0值
```go
//duck1是 *ToyDuck
dock1 := &ToyDuck{}
dock1.Swim()

douck2 := ToyDuck{}
douck2.Swim()

//duck3是 *ToyDuck
duck3 := new(ToyDuck)
duck.Swim()
```
12. 指针
和C,C++一样，*表示指针，&区地址
如果什声明了一个指针，但是没有赋值，那么他是nil
```go
//指针用*表示
var p *ToyDuck = &ToyDuck{}
//解引用，得到结构体
var duck ToyDuck = *p
dock.Swim()

//只是声明了，但是没有使用
var nilDuck *ToyDuck
if nilDuck == nil {
    fmt.Printf("nilDuck is nil")
}
```
13. 结构体的自引用
结构体内部引用自己，只能使用指针
准确来说，在整个引用链上，如果构成循环，那就只能用指针
```go
type Node struct {
    left  *Node
    right *Node
}
```

14. 结构体如何实现接口
当看到一只鸟走起来像鸭子、游起来像鸭子，叫起来像鸭子，那么这只鸟就可以被称为鸭子。
当一个结构体具备了接口的所有的方法的时候。它就实现了这个接口

15. 要点总结：type
1）type定义熟记。其中type A=B 这种别名，一般只用于兼容性处理
所以不需要过多关注；
2）先有抽象再有实现，所以要先定义接口
3）鸭子类型：一个结构体有某个接口的所有方法，它就实现了这个接口；
4）指针：方法接收器，遇事不决用指针

16. 空接口
空接口interface{}不包含任何方法，所以任何结构体都实现了该接口，
类似jave的object,即所谓的继承树根节点

17. json库
1)用于处理json格式的字符串
2)字段后面的内容被称为Tag，即标签，运行期间可以发射拿到
3)json库依据json数据到结构体的映射
4）典型的申明式API设计

