package main

import "fmt"

func main1()  {
	a,b,c,d := 1,2,2,3   //多重赋值也是需要:=的赋值运算符
	fmt.Print(a,b,c,d)
}


func  main2(){
	//交换a,b值得形式发生改变，形式与实现一致了
	a,b:=10,20
	a,b=b,a

	fmt.Println(a,b)
}

func  main()  {
	//匿名变量
	a,_,c:=1,2,3
	fmt.Println(a,c)
}


