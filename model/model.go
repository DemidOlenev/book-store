package model

import "fmt"

type Client struct {
	Id         int
	First_name string
	Last_name  string
	Phone      string
}

func (c Client) Info() string {
	return fmt.Sprintf("ID [%d] | NAME [%s] | LAST NAME [%s] | PHONE [%s]",
		c.Id, c.First_name, c.Last_name, c.Phone)
}

func ClientsInfo(clients ...Client) {
	for _, c := range clients {
		fmt.Println(c.Info())
	}
}

type Good struct {
	Id    int
	Name  string
	Price float32
}

func (g Good) Info() string {
	return fmt.Sprintf("ID [%d] | NAME [%s] | PRICE [%f]",
		g.Id, g.Name, g.Price)
}

func GoodsInfo(goods ...Good) {
	for _, g := range goods {
		fmt.Println(g.Info())
	}
}
