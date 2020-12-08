package main

import(
	"fmt"
)

func  display(mychannel chan int ){
	fmt.Println("Display goroutine")
	for(i:=1;i<=5;i++){
		mychannel<-i
	}
	close(mychannel)
}
func main(){
	mychannel:=make(chan int)
	go display(mychannel)
	for{
		num,status:=<-mychannel//read avalue from mychannel
		if status==false {
			break
		}
		fmt.Println("main goroutine:%d\n",num)
	}
}