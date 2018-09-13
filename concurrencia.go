package main

import "fmt"

var numberA int
var numberB int
var number int
var n1, i1, c1 int
var n, i, c int
func busqueda1(){
  c1=0
  for q:=numberA; q<number; q++ {
    n1=q-1
   for i1=1; i1<=n1; i1++{
     if(n1%i1==0) {
     c1++
   }
   }
   if(c1==2){
     fmt.Print(n1)
     fmt.Print("-")
   }
     c1=0
   }
 }


func busqueda2(){
  c=0
for q:=number; q<=numberB; q++{
  n=q
 for i=1; i<=n; i++{
   if(n%i==0) {
   c++
 }
 }
 if(c==2){
   fmt.Print(n)
   fmt.Print("-")
 }
    c=0
}

}

func main(){
  fmt.Println("Realiza una busqueda de rango a: ")
  fmt.Scan(&numberA)
  fmt.Println("al rango b: ")
  fmt.Scan(&numberB)
  number= (numberB/2)
  busqueda1()
 busqueda2()
}
