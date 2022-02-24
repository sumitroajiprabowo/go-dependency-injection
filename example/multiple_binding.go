/*
Example of Multiple Binding for Dependency Injector using Multiple Binding
in this case we have a same type data dependency and we need to inject it in different places
*/
package example

// Create database struct
type Database struct {
	Name string // Name of the database
}

type DatabasePostgresSQL Database //Create Alias for Database
type DatabaseMySQL Database       //Create Alias for Database

//Create a new database postgresql instance and return it as a pointer
func NewDatabasePostgresSQL() *DatabasePostgresSQL {
	return &DatabasePostgresSQL{
		Name: "PostgresSQL",
	}
}

// Create a new database mysql instance and return it as a pointer
func NewDatabaseMySQL() *DatabaseMySQL {
	return &DatabaseMySQL{
		Name: "MySQL",
	}
}

/*
Create a new struct database repository instance and return it as a pointer of database with alias
*/
type DatabaseRepository struct {
	PostgresSQL   *DatabasePostgresSQL
	DatabaseMySQL *DatabaseMySQL
}

/*
Create a new database repository with parameters postgresql and mysql and return it as a pointer
*/
func NewDatabaseRepository(postgresql *DatabasePostgresSQL, mysql *DatabaseMySQL) *DatabaseRepository {
	return &DatabaseRepository{
		PostgresSQL:   postgresql,
		DatabaseMySQL: mysql,
	}
}
