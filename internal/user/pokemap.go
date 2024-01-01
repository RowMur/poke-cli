package user

type pokemap struct {
	PrevLocationURL string `json:"prevLocationURL"`
	NextLocationURL string `json:"nextLocationURL"`
}

func (ms *pokemap) UpdateURLs(prev, next string) {
	ms.PrevLocationURL = prev
	ms.NextLocationURL = next
}