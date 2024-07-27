package res

type (
	S3 struct {
		Bucket    string `json:"bucket"`
		Region    string `json:"region"`
		AccessKey string `json:"accessKey"`
		SecretKey string `json:"secretKey"`
		PathStyle bool   `json:"pathStyle"`
		EndPoint  string `json:"endPoint"`
		UseSSL    bool   `json:"useSSL"`
		Port      int    `json:"port"`
	}

	S3Public struct {
		S3
		// BucketURL is the public URL of the bucket's root
		BucketURL string `json:"bucketURL"`
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

	OpenAPI struct {
		ApiKey string
		// OrganizationId for users with multiple organizations
		OrganizationId string
	}
)
