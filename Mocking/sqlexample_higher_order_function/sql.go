package sqlexample

import (
	"database/sql"
	"fmt"
)

// Whem I need to mock a function

// A estratégia é criar uma high order function, com a assinatura da funçcão que se deseja mockar

// Caso sem o mock
// func OpenDB(user, password, addr, db string) (*sql.DB, error) {
// 	conn := fmt.Sprintf("%s:%s@%s/%s", user, password, addr, db)
// 	return sql.Open("mysql", conn)
// }

// Assinatura da função a ser mockada
type sqlOpener func(string, string) (*sql.DB, error)

// Injetando a hig order function no lugar de chamar a função original
func OpenDB(user, password, addr, db string, sqlOpener sqlOpener) (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@%s/%s", user, password, addr, db)
	return sqlOpener("mysql", conn)
}
