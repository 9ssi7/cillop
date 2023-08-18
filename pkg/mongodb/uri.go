package mongodb

type UriParams struct {
	Host  string
	Port  string
	User  string
	Pass  string
	Db    string
	Query string
}

func CalcMongoUri(params UriParams) string {
	uri := "mongodb://"
	if params.User != "" && params.Pass != "" {
		uri += params.User + ":" + params.Pass + "@"
	}
	uri += params.Host + ":" + params.Port + "/" + params.Db
	if params.Query != "" {
		uri += "?" + params.Query
	}
	return uri
}
