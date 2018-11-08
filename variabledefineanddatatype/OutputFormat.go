package main

import "fmt"

func main1()  {
	 a:=3
	 b:=3.1415926
	 fmt.Printf("%3d\n",a);//d% 占位符，表示输出一个整型数据
	 fmt.Printf("%f\n",b)   //%f 默认保留6位小数

	 fmt.Printf("%.3f",b);  //&.3f 保留三位小数，第四位小数进行四舍五入
}

func main2()  {
	var a bool   //bool类型的值默认为false
	a = true
	fmt.Printf("%t",a)   //%t 占位符，表示输出一个布尔类型的数据
}


func main3(){
	var b string
	b = "nihao"

	fmt.Printf("%s",b)   //%s 占位符，表示输出一个字符串类型数据
}

func main()  {
	var c byte
	c='a'
	fmt.Println(c)  //直接输出的话就是对应的ASCII码值
	fmt.Printf("%c",c) //%c表示是输出一个字节类型的数据

}
