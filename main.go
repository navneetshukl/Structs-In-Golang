package main

import "fmt"

type Name struct{
	firstName string
	lastName string
}

func main() {
	/*mp := make(map[int]int)
	for i := 0; i < 10; i++ {
		mp[i] = i * i
	}
	fmt.Println(mp[0])*/
	/*mp:=make(map[int]Name);
	mp[0]=Name{
		firstName: "Navneet",
		lastName: "Shukla",
	}
	mp[1]=Name{
		firstName: "Yatinjal",
		lastName: "Shukla",
	}
	len:=len(mp)
	for i := 0; i < len; i++ {
		fmt.Println(mp[i].firstName+" "+mp[i].lastName)
		
	}*/
	mp:=make(map[int]Address)

	var n int
	fmt.Println("Enter total number of user you want to enter")
	fmt.Scanln(&n);
	for i:=0;i<n;i++{
		var name,city,country string
		fmt.Println("Enter Name : ")
		fmt.Scanln(&name)
		fmt.Println("Enter City : ")
		fmt.Scanln(&city)
		fmt.Println("Enter country : ")
		fmt.Scanln(&country)
		address:=Address{
			Name: name,
			City: city,
			Country: country,
		}
		mp[i]=address

	}
	fmt.Println(mp)
}