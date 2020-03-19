package database

func Drop(table interface{}) {
	Db.DropTableIfExists(&table)
}