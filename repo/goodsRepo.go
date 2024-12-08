package repo

import (
	"book-store/model"
	"book-store/resources"
	"database/sql"
	"fmt"
)

type GoodsRepo struct {
	conn *sql.DB
}

func NewGoodsRepo(c *connectionDB) *GoodsRepo {
	return &GoodsRepo{
		conn: c.conn,
	}
}

func (gr GoodsRepo) GetGoods() ([]model.Good, error) {
	Goods := make([]model.Good, 0)
	fmt.Println("Get all Goods query")
	rows, err := gr.conn.Query(resources.Get_all_goods)
	if err != nil {
		fmt.Println("Error occured during retrieving Goods: ", err)
		return nil, err
	}
	defer rows.Close()
	g := model.Good{}
	for rows.Next() {
		err := rows.Scan(&g.Id, &g.Name, &g.Price)
		if err != nil {
			fmt.Println("Error occured during parsing DB answer: ", err)
			return nil, err
		}
		Goods = append(Goods, g)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println("Error occured after DB answer: ", err)
		return nil, err
	}
	fmt.Println("All Goods were retrieved and parsed")
	return Goods, nil
}

func (gr GoodsRepo) GetGoodById(id int) (model.Good, error) {
	g := model.Good{}
	fmt.Println("Get good with id: ", id)
	err := gr.conn.QueryRow(resources.Get_good_by_id, id).Scan(&g.Id, &g.Name, &g.Price)
	if err != nil {
		fmt.Println("Error ocured during retreiving good with id ", err)
		return model.Good{}, err
	}

	fmt.Println("Retrieved good with id ", id)
	return g, nil
}

func (gr GoodsRepo) InsertNewGood(name string, price float32) (model.Good, error) {
	fmt.Println("Try to insert new good data into Goods ", name, price)
	stmt, err := gr.conn.Prepare(resources.Insert_new_good)
	if err != nil {
		fmt.Println("Error ocured during preparing query to DB ", err)
		return model.Good{}, err
	}
	defer stmt.Close()
	var newGood model.Good

	stmt.QueryRow(name, price).Scan(&newGood.Id, &newGood.Name, &newGood.Price)
	fmt.Println("Good was inserted with id: ", newGood.Id)
	return newGood, nil
}

func (gr GoodsRepo) DeleteGoodById(id int) error {
	fmt.Println("Try to delete good ", id)
	_, err := gr.conn.Exec(resources.Delete_good_by_id, id)
	if err != nil {
		fmt.Println("Error ocured during deleting good with id ", id, err)
		return err
	}
	fmt.Println("Good was deleted with id: ", id)
	return nil
}

func (gr GoodsRepo) UpdateGood(good model.Good) {
	fmt.Println("Try to update client with new data: ", good.Info())
	_, err := gr.GetGoodById(good.Id)
	if err != nil {
		fmt.Println("Good not found with id: ", good.Id, err)
	} else {
		_, err = gr.conn.Exec(resources.Update_good, good.Name, good.Price, good.Id)
		if err != nil {
			fmt.Println("Error ocured during updating good with data", good.Info(), err)
		}

		fmt.Println("Good was updated with data: ", good.Info())
	}
}