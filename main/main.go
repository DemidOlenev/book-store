package main

import (
	"book-store/model"
	"book-store/repo"
	"fmt"
)

func main() {
	
	connection := repo.GetConnDBInstance()
	connection.Ping()
	defer connection.Close()
	fmt.Println("*******************Test for clients*******************")
	clientRepository := repo.NewClientRepo(connection)
	client, _ := clientRepository.GetClientById(20)
	model.ClientsInfo(client)
	newClient, err := clientRepository.InsertNewClient("Томас", "Аллелуев", "+79991234455")
	if err == nil {
		model.ClientsInfo(newClient)
	}

	clients, _ := clientRepository.GetClients()
	model.ClientsInfo(clients...)
	client2, _ := clientRepository.GetClientById(2)
	model.ClientsInfo(client2)
	clientRepository.DeleteClientBy(8)

	client2.First_name = "Адам"
	clientRepository.UpdateClient(client2)
	updatedClient, _ := clientRepository.GetClientById(client2.Id)
	model.ClientsInfo(updatedClient)
	fmt.Println("*******************Test for goods*******************")
	goodsRepository := repo.NewGoodsRepo(connection)
	good, _ := goodsRepository.GetGoodById(5)
	model.GoodsInfo(good)
	newGood, err := goodsRepository.InsertNewGood("Телевизор", 102560.68)
	if err == nil {
		model.GoodsInfo(newGood)
	}

	goods, _ := goodsRepository.GetGoods()
	model.GoodsInfo(goods...)
	good5, _ := goodsRepository.GetGoodById(5)
	model.GoodsInfo(good5)
	goodsRepository.DeleteGoodById(9)

	good5.Name = "Мышка нарушка"
	goodsRepository.UpdateGood(good5)
	updatedGood, _ := goodsRepository.GetGoodById(good5.Id)
	model.GoodsInfo(updatedGood)

}
