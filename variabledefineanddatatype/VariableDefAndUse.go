package main

import "fmt"
//定义变量的形式为 var 变量名 变量类型  在定义的时候，如果没有进行初始化操作，会有一个默认值
func main1(){
	var a int
	a = 10
	a = a+25

	fmt.Print(a)
}

/**
计算圆的周长和面积
 */
func main2(){

	//使用;=的方式定义变量
	PI:=3.1415926
	r:=2.5
	l:=2*PI*r;

	fmt.Println(l)


	s:=PI*r*r
	fmt.Print(s)
}


func main3(){
	//自动推断类型
	//w:=5.0   //float64
	//p:=2;    //int
/*
	在go语言中不同的数据类型不能通过计算
	fmt.Println(w*p)//mismatched types float64 and int
*/

}

func main()  {
	a:=false
	b:=10
	c:=3.14
	d:='a'
	e:="abc"

	fmt.Println(a,b,c,d,e)
}