package coinbase

func hasFormBody(method string) bool {
	return method == "POST" || method == "PUT" || method == "PATCH"
}
