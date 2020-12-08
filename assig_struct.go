package main 
  
import "fmt"
  
// Defining a struct type 
type Address struct { 
    Name    string 
    city    string 
    MobileNO int
} 
  
func main() { 
  
    // Declaring a variable of a `struct` type 
    // All the struct fields are initialized  
    // with their zero value 
    var a Address  
    fmt.Println("Address0: ",a) 
  
    // Declaring and initializing a 
    // struct using a struct literal 
    a1 := Address{"Amir", "Mumbai", 8745213698} 
  
    fmt.Println("Address1: ", a1) 
  
    // Naming fields while  
    // initializing a struct 
    a2 := Address{Name: "shidhu", city: "Delhi", 
                                 MobileNO: 2514789584} 
  
    fmt.Println("Address2: ", a2) 
  
    // Uninitialized fields are set to 
    // their corresponding zero-value 
    a3 := Address{Name: "Rehan",city:"Lucknow",MobileNO:7388303699} 
    fmt.Println("Address3: ", a3) 
} 