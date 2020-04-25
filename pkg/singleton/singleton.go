package singleton

type sql struct {
	Connect string
}

var db *sql

func Database(conn string) *sql {
	if db == nil {
		db = &sql{
			Connect: conn,
		}
	}

	return db
}
