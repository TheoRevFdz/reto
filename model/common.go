package model

type Common struct {
	Available     int    `json:"available"`
	CollectionURI string `json:"collectionURI"`
	Items         []Item `json:"items"`
	Returned      int    `json:"returned"`
}

type Item struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
}
