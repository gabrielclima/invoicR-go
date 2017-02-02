package main

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
