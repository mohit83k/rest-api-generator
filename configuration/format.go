package configuration

type configuration struct {
	Server    server
	ApiRoutes []apiRoute
	MockDir   string
}

type server struct {
	Port           string   `json:"port"`
	Tls            bool     `json:"tls"`
	Certificate    string   `json:"certificate"`
	PrivateKey     string   `json:"privateKey"`
	CorsException  []string `json:"corsException"`
	AllowedHeaders []string `json:"allowedHeaders"`
}

type apiRoute struct {
	Url  string `json:"url"`
	File string `json:"file"`
}
