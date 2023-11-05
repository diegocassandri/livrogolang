package sqlexample

import (
	"database/sql"
	"fmt"
)

// Whem I need to mock a function

// Monkey Patching - Declarar uma variável que recebe a função que se deseja mockar

// Caso sem o mock
// func OpenDB(user, password, addr, db string) (*sql.DB, error) {
// 	conn := fmt.Sprintf("%s:%s@%s/%s", user, password, addr, db)
// 	return sql.Open("mysql", conn)
// }

// Declarar uma variável que recebe a função que se deseja mockar
var SQLOpen = sql.Open

// Injetando a hig order function no lugar de chamar a função original
func OpenDB(user, password, addr, db string) (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@%s/%s", user, password, addr, db)
	return SQLOpen("mysql", conn)
}
