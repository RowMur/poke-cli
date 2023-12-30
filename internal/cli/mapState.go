package cli

type mapState struct {
	PrevLocationURL string `json:"prevLocationURL"`
	NextLocationURL string `json:"nextLocationURL"`
}

func (ms *mapState) updateURLs(prev, next string) {
	ms.PrevLocationURL = prev
	ms.NextLocationURL = next
}