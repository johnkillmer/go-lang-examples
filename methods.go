package main

import "fmt"

type rect struct {
	width, height, x int
}

//pointer receiver - use when wanting to avoid a copy or when you want to mutate the receiving
func (r *rect) area() int {
   r.x=5;
   return r.width * r.height
}

//value receiver 
func (r rect) perim() int {
   r.x=2
   return 2 * r.width + 2 * r.height
}

func main() {
   r := rect{width:10, height:5}
   fmt.Println(r)	

   fmt.Println("area: ", r.area())
   fmt.Println(r.x)
   fmt.Println("perim: ", r.perim())
   fmt.Println(r.x)

   rp := &r
   fmt.Println("area: ", rp.area())
   fmt.Println("perim: ", rp.perim())
}