package main

import "fmt"

type person struct {
	firstname        string
	lastname         string
	favouritecountry []string
}

func main() {
	p1 := person{"dinesh", "bobba", []string{"USA", "austrilia", "italy"}}
	p2 := person{"ajith", "bolu", []string{"austria", "turkey", "dubai"}}
	fmt.Println("first person struct", p1)
	fmt.Println("person fist name", p1.firstname)
	for i := range p1.favouritecountry {
		fmt.Println(i, p1.favouritecountry[i])
	}
	fmt.Println("second person struct", p2)
	fmt.Println("person fist name", p2.firstname)
	for i := range p2.favouritecountry {
		fmt.Println(i, p2.favouritecountry[i])
	}

	persan := make(map[string]person)
	persan[p1.lastname] = p1
	fmt.Println(persan)

	//3

	honda: truck{
		a: vehicle{doors: "4door",
		colour: red,
		fourwheel: true
		}
	}
	kia: sedan{
		a: vehicle{doors: "2door",
		colour: white,
		fourwheel: true
		}
	}
	


}
type vehicle struct{
	doors string
	colour string
}

type truck struct{
 fourwheel bool
 	a vehicle

}

type sedan struct{
	luxury bool
a vehicle
}


