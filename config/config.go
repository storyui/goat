package config

type GoatExecConfig struct {
	Steps []GoatStepConfig `json:"steps"`
}

type GoatStepConfig struct {
	Name  string `json:"name"`
	Run   string `json:"run"`
	Delay int32  `json:"delay"`
}

type GoatServerConfig struct {
	User     string          `json:"user"`
	Host     string          `json:"host"`
	Port     int             `json:"port"`
	Password string          `json:"password"`
	Exec     *GoatExecConfig `json:"exec"`
}

type GoatConfig struct {
	Servers []map[string]GoatServerConfig `json:"servers"`
}
