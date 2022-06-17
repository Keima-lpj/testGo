package main

import (
	"fmt"
	project1 "testGo/project01"
)

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

type TestA struct {
	Name string
	age  int
}

type TestB struct {
	TestA
	Sex   int
	hobby string
}

func main() {
	//sliceAndArr();
	//structTest()
	//structTestPub()
	//structTestExtends()
	testReference()
}

//猜想，既然slice是数组的引用，那么定义一个slice引用某个数组，
// 修改了slice的值之后，数组的值也会修改吗？
func sliceAndArr() {
	//定义一个数组
	intArr := [5]int{1, 2, 3, 4, 5}

	intSlice := intArr[1:3]

	fmt.Println(intArr, intSlice)

	fmt.Printf("数组的类型:%T，切片的类型:%T\n", intArr, intSlice)

	//输出数组各个值的地址
	fmt.Printf("%p ", &intArr[0])
	fmt.Printf("%p ", &intArr[1])
	fmt.Printf("%p ", &intArr[2])
	fmt.Printf("%p ", &intArr[3])
	fmt.Printf("%p ", &intArr[4])
	fmt.Println()

	//输出切片各个值的地址
	fmt.Printf("%p ", &intSlice[0])
	fmt.Printf("%p ", &intSlice[1])
	fmt.Println()

	//这里可以看到，切片的第一个元素的地址和数组的第二个元素的地址是一样的
	//切片的第二个元素和数组的第三个元素的地址也是一样的，可以证实切片确实是数组的引用

	//这里修改切片的第一个值，看是否影响数组
	intSlice[0] = 10
	fmt.Println(intArr, intSlice)
	//可以看到确实影响了数组。

	//切片还可以继续切片，此时第二次产生的切片是否也指向第一次切片引用的数组？、
	intSlice2 := intSlice[:]
	//输出切片各个值的地址
	fmt.Printf("%p ", &intSlice2[0])
	fmt.Printf("%p ", &intSlice2[1])
	fmt.Println()
	//修改切片2的值，观察数组和切片1的值，可以看到数组和两个切片的值都变更了
	intSlice2[0] = 20
	fmt.Println(intArr, intSlice, intSlice2)

	//append函数的一些想法，声明一个切片，长度为3，容量为6
	intSlice3 := make([]int, 3, 6)
	fmt.Printf("%p ", &intSlice3[0])
	fmt.Printf("%p ", &intSlice3[1])
	fmt.Printf("%p ", &intSlice3[2])
	fmt.Println()

	//这个时候append一个小于等于三位的切片，观察地址变化，发现前三位的地址并未发生变化
	intSlice3 = append(intSlice3, intSlice2...)
	fmt.Printf("%p ", &intSlice3[0])
	fmt.Printf("%p ", &intSlice3[1])
	fmt.Printf("%p ", &intSlice3[2])
	fmt.Printf("%p ", &intSlice3[3])
	fmt.Printf("%p ", &intSlice3[4])
	fmt.Println()

	//此时再次追加一个切片，发现所有的地址都变化了，证明如果给切片追加超过其容量的值后，
	//系统会在内存中重新分配一块地址，并将旧地址的值copy到新的地址上去。
	intSlice3 = append(intSlice3, intSlice...)
	fmt.Printf("%p ", &intSlice3[0])
	fmt.Printf("%p ", &intSlice3[1])
	fmt.Printf("%p ", &intSlice3[2])
	fmt.Printf("%p ", &intSlice3[3])
	fmt.Printf("%p ", &intSlice3[4])
	fmt.Printf("%p ", &intSlice3[5])
	fmt.Printf("%p ", &intSlice3[6])
	fmt.Println()

	//如果不使用append而是直接增加值呢？   直接增加值的情况只要是超过当前数组的长度，就会提示越界
	/*	intSlice4 := make([]int, 2, 3)
		fmt.Printf("%p ", &intSlice4[0])
		fmt.Printf("%p ", &intSlice4[1])
		intSlice4[2] = 20
		fmt.Printf("%p ", &intSlice4[0])
		fmt.Printf("%p ", &intSlice4[1])
		fmt.Printf("%p ", &intSlice4[2])
		fmt.Println()*/

	//copy函数猜想
	floatSlice := []float64{1.1, 1.2, 1.3}
	var floatSlice2 = make([]float64, 5)
	copy(floatSlice2, floatSlice)
	//使用copy函数后，查看对应的地址和值
	fmt.Println(floatSlice, floatSlice2)
	fmt.Printf("%p ", &floatSlice[0])
	fmt.Printf("%p ", &floatSlice[1])
	fmt.Printf("%p ", &floatSlice[2])
	fmt.Println()
	fmt.Printf("%p ", &floatSlice2[0])
	fmt.Printf("%p ", &floatSlice2[1])
	fmt.Printf("%p ", &floatSlice2[2])
	//发现地址完全不一样，证明使用了copy将一个切片的值给另一个切片后，两个切片之间是值copy关系
	//并未指向同一内存地址
}

//猜想，一个结构体变量中的属性在内存中的地址是连续的吗？
func structTest() {
	var cat1 Cat

	cat1.Name = "喵喵"
	cat1.Age = 2
	cat1.Color = "虎斑"
	cat1.Hobby = "睡觉"

	//可以看出，一个结构体变量的内存地址为其第一个属性的内存地址
	fmt.Printf("整个结构体变量的内存地址为:%p\n", &cat1)
	fmt.Printf("Name的内存地址:%p\n", &cat1.Name)
	fmt.Printf("Age的内存地址:%p\n", &cat1.Age)
	fmt.Printf("Color的内存地址:%p\n", &cat1.Color)
	fmt.Printf("Hobby的内存地址:%p\n", &cat1.Hobby)
}

//猜想，关于结构体的包引用和访问控制问题
func structTestPub() {
	//这里直接使用project1包的student结构体
	//由于student为大写，这里没有任何问题
	var stu project1.Student
	stu.Name = "feifei"
	stu.Score = 22.2
	fmt.Println(stu)

	//这里使用project1包的teacher结构体，可以看到不能直接使用，因为其为私有
	//var tea project1.teacher
	//这里使用开放出的方法GetNewTeacher来获取teacher的对象指针
	tea := project1.GetNewTeacher()
	tea.Name = "laodie"
	tea.Score = 78.5
	fmt.Println(*tea)

	//那么问题来了，如果Student的对象里面的属性为小写开头呢（私有）
	//分别访问校长和辅导员
	var master = project1.Master{
		Name: "laoma",
		//score:100,    //可以看到申明为小写的属性不能使用
	}
	//这里使用校长的内置方法来设置分数
	master.SetScore(78.4)
	fmt.Println(master)

	//这里访问辅导员
	coun := project1.GetNewCounselor()
	coun.Name = "111"
	//coun.score，这里同样不能使用小写的私有属性
	//这里同样使用指导员的内置方法来设置分数
	coun.SetScore(100.0)
	fmt.Println(*coun)
}

//笔记中说结构体可以使用继承的匿名结构体中的所有字段和方法，即首字母大写或者小写的字段，方法都能使用
//猜想，A为父类，B为子类。看看区分成哪几种情况
func structTestExtends() {
	//第一种情况，A和B在同一个包下，同时在这个包使用，
	// 毫无疑问，无论是公有属性还是私有属性都可以使用
	var b TestB
	b.Name = "a"
	b.age = 18
	b.Sex = 1
	b.hobby = "basketball"

	//第二种情况在这里使用同一个包下的其他的私有结构体，并访问其私有属性，也是可以的
	//这里有一个坑，如果需要使用同一个包的其他文件下的结构体或者方法，则在go run 和go build的时候
	//需要将该文件一起编译
	var cat cat
	cat.name = "miaomiao"
	fmt.Println(cat)

	//第三种情况，在这里使用不同包的结构体。由于student和person是私有，不能实例化
	// 只能实例化公有的worker
	var worker project1.Worker
	//这里看看能不能使用公有的Worker继承自私有person结构体里的属性
	worker.Name = "lpj"
	//显然这里还是需要看属性的大小写。如果是继承自私有结构体，但是属性是大写的，则在外部也是可以用的
	//如果是小写，则即便继承过来也会认定为私有，不能使用
	//worker.age = 18   //这里会报错
	worker.Worker()
}

//猜想，据说golang中main包中不同的文件的代码不能直接互相调用，其他包的可以
func testReference() {

	//这里调用同样是main包的struct.go里的Dog结构体
	dog := Dog{name: "feifei", age: 18}
	//打印之，这里如果直接使用go run testMind.go或者go build testMind.go
	//会提示undefined:Dog。这个时候需要在编译的时候加上struct.go
	//即使用go run testMind.go struct.go或者go build testMind.go struct.go
	//将提供Dog结构体的文件也一并编译，才会正常运行
	fmt.Println(dog)

	//那么如果是不同的包呢？
	//我们在包project1中建立了fish和behavior文件
	fish := project1.Fish{
		Name: "shark",
	}
	//打印之
	fmt.Println(fish)
	//调用fish文件的方法
	project1.FishSay()
	//调用fish结构体的方法，都可以完美执行，且下面的方法是依赖behavior中的方法的
	fish.Say()
	fish.FishAge()

	//这里证明，如果是另一个包，在引用了对应的文件之后，其互相的调用是不受编译影响的
	//证明只有main包受此影响
}
