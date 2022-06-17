package main

import (
	"fmt"
	"math"
)

func main() {

	//&运算符，两个都为1才为1
	//2的补码：0000 0010  3的补码：0000 0011 得到结果 0000 0010
	fmt.Println(2 & 3) //2
	//|运算符，一个为1就为1
	//2的补码：0000 0010  3的补码：0000 0011 得到结果 0000 0011
	fmt.Println(2 | 3) //3
	//^运算符，两个不同就为1
	//2的补码：0000 0010  3的补码：0000 0011 得到结果 0000 0001
	fmt.Println(2 ^ 3) //1
	//-2的补码：（原码：0000 0010 反码：0111 1101 补码：0111 1110）
	//2的补码：0000 0010 结果为：0111 1100 此时取到的还是一个补码
	//取这个补码的原码：先取反码：0111 1011 则得到原码：0000 0100
	fmt.Println(-2 ^ 2) //4

	//>>表示向右移动两位，舍弃溢出的位，且高位用符号位补
	//即为：0000 0001 => 0000 0000
	fmt.Println(1 >> 2) //0
	//<<表示向左移动两位，符号位不变，低位补0
	//即为：0000 0001 => 0000 0100
	fmt.Println(1 << 2) //4
	//即为：0000 0101 => 0001 0100
	fmt.Println(5 << 2) // 20
	//即为：0000 0101 => 0000 0001
	fmt.Println(5 >> 2) //1

}

func aaa(a, b, c float64) (x1, x2 float64) {

	if b*b-4*a*c == 0 {
		x1 = (math.Sqrt(b*b-4*a*c) - b) / (2 * a)
		x2 = x1
	}

	return x1, x2
}
