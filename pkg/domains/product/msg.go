package product

type msg struct {
	Failed    string
	InvalidID string
	NotFound  string
}

var messages = msg{
	Failed:    "product_create_failed",
	InvalidID: "product_invalid_id",
	NotFound:  "product_not_found",
}
