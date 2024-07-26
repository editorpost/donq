package res

type (
	S3 struct {
		Host               string `json:"host"`
		Port               int    `json:"port"`
		User               string `json:"user"`
		Dbname             string `json:"dbname"`
		SSLMode            string `json:"sslmode"`
		Password           string `json:"password"`
		RootCertificatePEM string `json:"root_certificate_pem"`
	}

	Postgresql struct {
		Host               string `json:"host"`
		Port               int    `json:"port"`
		User               string `json:"user"`
		Dbname             string `json:"dbname"`
		SSLMode            string `json:"sslmode"`
		Password           string `json:"password"`
		RootCertificatePEM string `json:"root_certificate_pem"`
	}

	Metrics struct {
		URL string `json:"url"`
	}

	Logs struct {
		URL string `json:"url"`
	}
)
