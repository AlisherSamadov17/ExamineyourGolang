package main

import (
	"encoding/json"
	"fmt"
	"os"
	// "sort"
	// "math"
)

type Ticket struct {
	Id       int    `json:"id"`
	From     string `json:"from"`
	To       string `json:"to"`
	Type     string `json:"type"`
	Price    int    `json:"price"`
	Currency string `json:"currency"`
}

type Custumer struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Ticket Ticket `json:"ticket"`
}

func main() {
	custumer := []Custumer{}
	file, err := os.Open("file.json")
	if err != nil {
		fmt.Println(err)
	}

	if err := json.NewDecoder(file).Decode(&custumer); err != nil {
		fmt.Println(err)
	}


	//-------------------------_1-ex_---------------------------------------
	//    var AllPriceIMP int
	// for _,val :=range custumer{
	// 	AllPriceIMP = val.Ticket.Price
	//  if AllPriceIMP>150 && AllPriceIMP<300 {
	// 	fmt.Println("Price:",AllPriceIMP,"$","","Name:",val.Name,"","PhoneNumber:",val.Phone)
	//  }
	// }

	//----------------------------_2-ex_--------------------------------------
	// var AllPrice int
	// for _,val :=range custumer{
	//  	AllPrice = val.Ticket.Price
	// 	fmt.Println("the price of ticket:",AllPrice,"$","","from:",val.Ticket.From,"","to:",val.Ticket.To)
	// }

	//-----------------------------_3-ex_-----------------------------------------
	// totalPrice := 0
	// countBusinessTickets := 0

	// for _, val := range custumer {
	// 	if val.Ticket.Type == "business" {
	// 		totalPrice += val.Ticket.Price
	// 		countBusinessTickets++
	// 	}
	// }

	// if countBusinessTickets > 0 {
	// 	averagePrice := float64(totalPrice) / float64(countBusinessTickets)
	// 	fmt.Printf("O`rtacha narxi: %.2f %s\n", averagePrice,"USD")
	// }

	//----------------------------_4-ex_---------------------------------
	// totalALLBusness:=0
	//  total:=0
	// total2AllEconomic:=0
	//  total2:=0

	// for _,val:=range custumer{
	// if val.Ticket.Type=="business" {
	// total+=val.Ticket.Price
	// totalALLBusness++
	// }
	// if val.Ticket.Type=="economic" {
	// total2+=val.Ticket.Price
	// total2AllEconomic++
	// }
	// }
	// fmt.Println("All prices of business:",total,custumer[0].Ticket.Currency)
	// fmt.Println("All ordered business class numbers:",totalALLBusness)
	// fmt.Println("---------------------------------------------------------")
	// fmt.Println("---------------------------------------------------------")
	// fmt.Println("All prices of economic",total2,custumer[0].Ticket.Currency)
	// fmt.Println("All ordered econimic class numbers:",total2AllEconomic)

	// ----------------------------_5-ex_-------------------------------------
	// for _,val:=range custumer{
	// 	fmt.Println("Going to:",val.Ticket.To,"City")
	// }

	//----------------------------_6-ex_-------------------------------------

	// for i := range custumer {
	// 	sort.Slice(custumer[i:], func(j, k int) bool {
	// 		return custumer[i+j].Ticket.Price < custumer[i+k].Ticket.Price
	// 	})
		
	// }

	// for _, customer := range custumer {
	// 	fmt.Println("custumer:", customer.Name, "price:", customer.Ticket.Price, "USD")
	// }

	//----------------------------_7-ex_-------------------------------------
	// ----------------eng uzun email ni topish--------------------------
	// 	longestEmail := findLongestEmail(custumer)
	// fmt.Println("Longest email:", longestEmail)


	//----------------------------_8-ex_-------------------------------------
 
// 	totalPrice := 0
// 	minPrice := math.MaxInt
// 	maxPrice := math.MinInt

// 	for _, c := range custumer {
// 		price := c.Ticket.Price
// 		totalPrice += price

// 		if price < minPrice {
// 			minPrice = price
// 		}

// 		if price > maxPrice {
// 			maxPrice = price
// 		}
// 	}

// 	averagePrice := float64(totalPrice) / float64(len(custumer))


// 	fmt.Printf("O`rtacha narx: %.2f %s\n", averagePrice,"USD")
// 	fmt.Printf("Eng kichik narx: %d %s\n", minPrice,"USD" )
// 	fmt.Printf("Eng katta narx: %d %s\n", maxPrice, "USD")



//----------------------------_9-ex_-------------------------------------    
	// for _,val:=range custumer{
	// if val.Ticket.To =="Sydney" {
	// 	fmt.Println(val)
	// }	
	// }
//----------------------------_10-ex_-------------------------------------    
// totalPrice:=0
// countBusinessTickets:=0

// for _,val:= range custumer{
// if val.Ticket.Type=="economic"{
//   totalPrice+=val.Ticket.Price
//   countBusinessTickets++
// }	
// }
// if countBusinessTickets > 0  {
// 	averagePrice:=float64(totalPrice)/float64(countBusinessTickets)
// 	fmt.Printf("O`rtacha narxi:%.2f %s\n",averagePrice,"USD")

// }


}

//----------------------------_7-ex_-------------------------------------
// // eng uzun email ni topish
// func findLongestEmail(circling []Custumer) string {
// 	longestEmail := ""
// 	maxLength := 0

// 	for _, val := range circling {
// 		emailLength := len(val.Email)
// 		if emailLength > maxLength { 
// 			maxLength = emailLength 
// 			longestEmail = val.Email
// 		}
// 	}

// 	return longestEmail
// }
