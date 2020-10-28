package modules

func GetModulesList() []IModule{

	minimega := &Minimega{
		Module{
			Name: "minimega",
			ConfigPath: "",
		},
	}

	opendss := &OpenDSS{
		Module{
			Name: "opendss",
			ConfigPath: "",
		},
	}

	parser := &Parser{
		Module{
			Name: "parser",
			ConfigPath: "",
		},
	}

	var list []IModule = []IModule{ minimega, opendss, parser }

	return list
}

type Module struct {
	Name string
	ConfigPath string
}

type IModule interface {
	Configure(path string) bool
	GetName() string
}
