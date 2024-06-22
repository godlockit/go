package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)
func add(x int, y int)/* (x,y int)*/int {
	y = rand.Intn(10)
	fmt.Println("It is your numder", x+y, math.Sqrt(4))

	return x+y
}
func swap(y,x string) (string,string){
	return x, y
}
func split(sum int)(x,y int){
	x = sum * 5
	y = sum - x
	return x , y
}
func leetters(name string) string{
	if (len(name) <= 8) {
		fmt.Println("word if short for this script")
		fmt.Println("1ая буква - ", name[0])
	}else{
		fmt.Println("9ая буква - ", name[9])	
	}
	return name
}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
type Vertex struct {
	Lat  int
}

var m = map[string]Vertex{
	"I": Vertex{
		1,
	},
	"II": Vertex{
		2,
	},
	"III": Vertex{
		3,
	},
	"IV": Vertex{
		4,
	},
	"V": Vertex{
		5,
	},
	"VI": Vertex{
		6,
	},
	"VII": Vertex{
		7,
	},
	"VIII": Vertex{
		8,
	},
	"IX": Vertex{
		9,
	},
	"X": Vertex{
		10,
	},
}


 var i,j = 1,2
func main() {
	fmt.Println("Hello World")

	add(42,7)

	a1:=6
	b1:=4
	c1 := a1 % b1
	fmt.Println("line = ", c1)

	a,b := swap("lol","papich")
	fmt.Println(a,b)

	x1,y1 := split(50)
	fmt.Println("split sum, x=" ,x1, "y=",y1)

	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	// fmt.Println(add(42,7))

	leetters("hubabu")

	sum:=0
	for i:=0;i<10;i++{
		sum +=i
	}
	fmt.Println(sum)

	sum = 1
	for ; sum<1000;{
		sum+=sum
	}

	fmt.Println(sum)
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println(sqrt(2),sqrt(-4))

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}
