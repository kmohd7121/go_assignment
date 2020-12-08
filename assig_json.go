package main
import ("encoding/json"
	"fmt")
type person struct{
	FirstName string
	LastName string 
	Age  int
}  
func main(){
	p1:=person{
		FirstName:"Mohd",
		LastName:"kalam",
		Age:22,
	}
	p2:=person{
		FirstName:"Miss",
		LastName:"leon",
		Age:26,
	}
	people:=[]person{p1,p2}
	fmt.Println(people)
	bs,err:=json.Marshal(people)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}