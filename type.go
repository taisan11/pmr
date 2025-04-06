package main

type PM struct {
	//名前
	Name      string  `json:"Name"`
	Install   *string `json:"Install"`
	Uninstall *string `json:"Uninstall"`
	Update    *string `json:"Update"`
	UpdateAll *string `json:"UpdateAll"`
	Run       *string `json:"Run"`
	Admin     bool    `json:"Admin"`
}

type Config struct {
	Level []string `json:"Level"`
}
