package model

type Story struct {
	Available     int         `json:"available"`
	CollectionURI string      `json:"collectionURI"`
	Items         []ItemStory `json:"items"`
	Returned      int         `json:"returned"`
}

type ItemStory struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}
