package main
import ("fmt"
		"errors")
func main(){
	total,err:=sumOf(10,8)
	if err!=nil{
		fmt.Println("there wae an error",err)
	}else{
		fmt.Println(total)
	}
}
func sumOf(start int,end int) (int,error) {
	if start>end{
		return 0,errors.New("start is greater than end")
	}
	total:=0
	for i:=start;i<=end;i++{
		total+=i
	}
	return total,nil
}