package environments

type DBConfig struct {
	URI     string
	Options OptionsConfig
}

type OptionsConfig struct {
	useUnifiedTopology bool
	useNewUrlParser    bool
	useCreateIndex     bool
	useFindAndModify   bool
}

var (
	BaseURL        string = "http://localhost:3000"
	EncodingString string = "0BPhrNcUJ1"
	Mongo                 = map[string]DBConfig{
		"main": {
			URI: "mongodb://localhost/urlshortener",
			Options: OptionsConfig{
				useUnifiedTopology: true,
				useNewUrlParser:    true,
				useCreateIndex:     true,
				useFindAndModify:   false,
			},
		},
	}
)
