package config

type Mysql struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Path     string `yaml:"path"`
	Post     string `yaml:"post"`
	DBName   string `yaml:"db-name"`
}

func (m *Mysql) Dsn() string {
	return m.UserName + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Post + ")/" + m.DBName
}
