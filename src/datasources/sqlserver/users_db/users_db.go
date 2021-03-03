package users_db

import (
	"database/sql"
	"flag"
	"fmt"
	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/mnasruul/bookstore_utils-go/logger"
	"log"
	"os"
)

const (
	mssqlServer   = "mssqlServer"
	mssqlPort     = "mssqlPort"
	mssqlUsername = "mssqlUsername"
	mssqlPassword = "mssqlPassword"
)

var (
	Client *sql.DB

	debug         = flag.Bool("debug", true, "enable debugging")
	password      = flag.String("password", os.Getenv(mssqlPassword), "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", os.Getenv(mssqlServer), "the database server")
	user          = flag.String("user", os.Getenv(mssqlUsername), "the database user")
	database      = flag.String("database", "db_users", "the database name")
)

type UserS struct {
	UserId   int64
	UserName string
}

func init() {
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;", *server, *user, *password, *port)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	var err error
	Client, err = sql.Open("mssql", connString)

	if err != nil {
		log.Println("databases error", err)
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		log.Println("databases error", err)
		panic(err)
	}
	mssql.SetLogger(logger.GetLogger())
	log.Println("database successfully configured")
}
