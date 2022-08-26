package Config

type ConfigModel struct {
	Host            string
	Port            int
	User            string
	Password        string
	Dbname          string
	Sslmode         string
	SenderURL       string
	ConsumerURL     string
	CountPerRoutine int
}
