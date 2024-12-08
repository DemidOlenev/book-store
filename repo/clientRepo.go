package repo

import (
	"book-store/model"
	"book-store/resources"
	"database/sql"
	"fmt"
)

type ClientRepo struct {
	conn *sql.DB
}

func NewClientRepo(c *connectionDB) *ClientRepo {
	return &ClientRepo {
		conn: c.conn,
	}
}

func (cr ClientRepo) GetClients() ([]model.Client, error) {
	clients := make([]model.Client, 0)
	fmt.Println("Get all clients query")
	rows, err := cr.conn.Query(resources.Get_all_clients)
	if err != nil {
		fmt.Println("Error occured during retrieving clients: ", err)
		return nil, err
	}
	defer rows.Close()
	c := model.Client{}
	for rows.Next() {
		err := rows.Scan(&c.Id, &c.First_name, &c.Last_name, &c.Phone)
		if err != nil {
			fmt.Println("Error occured during parsing DB answer: ", err)
			return nil, err
		}
		clients = append(clients, c)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println("Error occured after DB answer: ", err)
		return nil, err
	}
	fmt.Println("All clients were retrieved and parsed")
	return clients, nil
}

func (cr ClientRepo) GetClientById(id int) (model.Client, error) {
	c := model.Client{}
	fmt.Println("Get client with id: ", id)
	err := cr.conn.QueryRow(resources.Get_client_by_id, id).Scan(&c.Id, &c.First_name, &c.Last_name, &c.Phone)
	if err != nil {
		fmt.Println("Error ocured during retreiving client with id ", err)
		return model.Client{}, err
	}

	fmt.Println("Retrieved client with id ", id)
	return c, nil
}

func (cr ClientRepo) GetClientByPhone(phone string) (model.Client, error) {
	c := model.Client{}
	fmt.Println("Get client with phone: ", phone)
	err := cr.conn.QueryRow(resources.Get_client_by_phone, phone).Scan(&c.Id, &c.First_name, &c.Last_name, &c.Phone)
	if err != nil {
		fmt.Println("Error ocured during retreiving client with phone ", err)
		return model.Client{}, err
	}

	fmt.Println("Retrieved client with phone ", phone)
	return c, nil
}

func (cr ClientRepo) InsertNewClient(f_name string, l_name string, phone string) (model.Client, error) {
	fmt.Println("Try to insert new client data into clients ", f_name, l_name, phone)
	stmt, err := cr.conn.Prepare(resources.Insert_new_client)
	if err != nil {
		fmt.Println("Error ocured during preparing query to DB ", err)
		return model.Client{}, err
	}
	defer stmt.Close()
	var newClient model.Client

	stmt.QueryRow(f_name, l_name, phone).Scan(&newClient.Id, &newClient.First_name, &newClient.Last_name, &newClient.Phone)
	if newClient.Id == 0 {
		fmt.Println("Client with this phone is already exists: ", newClient.Phone)
	} else {
		fmt.Println("Client was inserted with id: ", newClient.Id)
	}
	return newClient, nil
}

func (cr ClientRepo) DeleteClientBy(id int) error {
	fmt.Println("Try to delete client ", id)
	_, err := cr.conn.Exec(resources.Delete_client_by_id, id)
	if err != nil {
		fmt.Println("Error ocured during deleting client with id ", id, err)
		return err
	}
	fmt.Println("Client was deleted with id: ", id)
	return nil
}

func (cr ClientRepo) UpdateClient(client model.Client) {
	fmt.Println("Try to update client with new data: ", client.Info())
	_, err := cr.GetClientById(client.Id)
	if err != nil {
		fmt.Println("Client not found with id: ", client.Id)
	} else {
		_, err = cr.conn.Exec(resources.Update_client, client.First_name, client.Last_name, client.Phone, client.Id)
		if err != nil {
			panic(err)
		}

		fmt.Println("Client was updated with data: ", client.Info())
	}
}