package models

type User struct {
	UID       uint32 `json:"_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    int8   `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewUser(user User) (bool, error) {
	con := Connect()
	defer con.Close()
	transaction, err := con.Begin()
	if err != nil {
		return false, err
	}
	query := "insert into users (username, email, password) values ($1,$2,$3) returning uid"
	{
		statement, err := transaction.Prepare(query)
		if err != nil {
			transaction.Rollback()
			return false, err
		}
		defer statement.Close()
		err = statement.QueryRow(user.Username, user.Email, user.Password).Scan(&user.UID)
		if err != nil {
			transaction.Rollback()
			return false, err
		}
	}
	query = "insert into wallets (usr) values ($1)"
	wallet := Wallet{User: user}
	wallet.GeneratePubKey()
	{
		statement, err := transaction.Prepare(query)
		if err != nil {
			transaction.Rollback()
			return false, err
		}
		_, err = statement.Exec(wallet.PublicKey, wallet.User.UID)
		if err != nil {
			transaction.Rollback()
			return false, err
		}
	}
	return true, transaction.Commit()
}

// GetUsers get users
func GetUsers() ([]User, error) {
	con := Connect()
	defer con.Close()
	query := "select * from users"
	resultSet, err := con.Query(query)
	if err != nil {
		return nil, err
	}
	defer resultSet.Close()
	var users []User
	for resultSet.Next() {
		var user User
		err := resultSet.Scan(&user.UID, &user.Username, &user.Email, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
