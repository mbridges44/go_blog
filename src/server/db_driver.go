package main

import(
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type Db_Config struct {
	Db_user string
	Db_pass string
	Db_name string
	Db_address string
}

var dbConfig Db_Config
var db *sql.DB

func init(){
	readConfig("db_config.json")
	initDatabase()
	getEntry(1);
}

func readConfig(path string){
	ParseJSON(path, &dbConfig)
}

func getEntry(id int) *Entry{
	statement, err := db.Prepare("SELECT * FROM UserPost WHERE entryID = ?")
	if(err != nil){
		panic(err.Error());
	}
	var blogEntry Entry

	err = nil
	err =  statement.QueryRow(id).Scan(&blogEntry.entryID, &blogEntry.Title, &blogEntry.Body, &blogEntry.Author)
	if(err != nil){
		panic(err.Error());
	}

	return &blogEntry
}

/*func SaveEntry(entry *Entry) error{

}*/



func initDatabase(){
	var err error
	println("Connecting to database")
	db, err = sql.Open("mysql", dbConfig.Db_user + ":" + dbConfig.Db_pass + "@/blog")
	if(err != nil) {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	println("DATABASE CONNECTION SUCCESSFUL")
}
