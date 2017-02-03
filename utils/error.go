package utils

// JsonErr struct for error JSON response
type JsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

// CheckErr - a common function for error check
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
