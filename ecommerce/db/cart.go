package db

import (
	"ecommerce/models"
	"fmt"
)

type Cart struct {
	ProductName string `json:"product_name" validate:"required,alpha"`
	Quantity    string `json:"quantity" validate:"required"`
}

func GetCart(id string, ch chan []models.CartList) {
	query := "SELECT product_name,quantity,price from cart where user_id=" + id + ";"
	rows, err := Db.Query(query)
	if err != nil {
		ch <- []models.CartList{}
	}

	var AllProduct []models.CartList

	for rows.Next() {
		var Product models.CartList
		err := rows.Scan(&Product.ProductName, &Product.Quantity, &Product.Price)
		if err != nil {
			ch <- []models.CartList{}
		}
		AllProduct = append(AllProduct, Product)
	}
	ch <- AllProduct
}

func InsertToCart(item Cart, id string) error {

	product := GetProduct(item)
	query := "INSERT INTO cart values('" + id + "','" + product.Name + "','" + product.Price + "','" + item.Quantity + "');"
	_, err := Db.Exec(query)
	return err

}
func ShowCart(id string) ([]models.CartList, string) {
	listch := make(chan []models.CartList)
	totalch := make(chan string)

	go GiveMeTotal(id, totalch)
	go GetCart(id, listch)

	list := <-listch
	total := <-totalch
	total = fmt.Sprintf("Total =%s", total)
	return list, total

}

func StringToInt(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		n = (n*10 + int(s[i]-'0'))
	}
	return n
}

// func IntToString(n int) string {
// 	if n == 0 {
// 		return "0"
// 	}
// 	s := ""
// 	for n > 0 {
// 		s += string(n%10 + '0')
// 		n /= 10
// 	}

// 	return Reverse(s)
// }
// func Reverse(s string) string {
// 	runes := []rune(s)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }
