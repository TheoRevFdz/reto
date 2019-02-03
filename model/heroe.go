package model

type Marvel struct {
	Code            int    `json:"code"`
	Status          string `json:"status"`
	Copyright       string `json:"copyright"`
	AttributionText string `json:"attributionText"`
	AttributionHTML string `json:"attributionHTML"`
	Etag            string `json:"etag"`
	Data            Data   `json:"data"`
}

type Data struct {
	Offset  int     `json:"offset"`
	Limit   int     `json:"limit"`
	Total   int     `json:"total"`
	Count   int     `json:"count"`
	Results []Heroe `json:"results"`
}

type Heroe struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Modified    string    `json:"modified"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	ResourceURI string    `json:"resourceURI"`
	Comics      Common    `json:"comics"`
	Series      Common    `json:"series"`
	Stories     Story     `json:"stories"`
	Events      Common    `json:"events"`
	Urls        Routes    `json:"urls"`
}

type Heroes []Heroe
