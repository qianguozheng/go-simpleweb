package airdisk
import "flag"


type Options struct {
	Database string
	Port string
	Logto string
	Loglevel string
}

func ParseArgs() *Options{
	database := flag.String("db", "./airdisk.db", "Assign db used to connect")
	port := flag.String("port", ":38001", "Assign port used to serve")
	logto := flag.String("log","stdout", "Write log messages to this file.")
	loglevel := flag.String("log-level", "DEBUG", "The level of messages to log. one of: DEBUG, INFO, WARNING, ERROR")

	flag.Parse()
	return &Options{
		Database: *database,
		Port : *port,
		Logto: *logto,
		Loglevel: *loglevel,
	}
}
