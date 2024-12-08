package resources

var (
	Get_all_clients = "SELECT * FROM clients"
	Get_client_by_id = "SELECT * FROM clients WHERE id = $1"
	Get_client_by_phone = "SELECT * FROM clients WHERE phone = $1"
	Insert_new_client = "INSERT INTO clients(first_name, last_name, phone) VALUES($1, $2, $3) RETURNING id, first_name, last_name, phone"
	Delete_client_by_id = "DELETE FROM clients WHERE id = $1"
	Update_client = "UPDATE clients SET first_name = $1, last_name = $2, phone = $3 WHERE id = $4"

	Get_all_goods = "SELECT * FROM goods"
	Get_good_by_id = "SELECT * FROM goods WHERE id = $1"
	Get_good_by_name_and_price = "SELECT * FROM goods WHERE name = $1 and price = $2"
	Insert_new_good = "INSERT INTO goods(name, price) VALUES($1, $2) RETURNING id, name, price"
	Delete_good_by_id = "DELETE FROM goods WHERE id = $1"
	Update_good = "UPDATE goods SET name = $1, price = $2 WHERE id = $3"
)