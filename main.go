package main

/*store book data in memory => struct tags ie `json:"title"`
specify shape when content is serialized into JSON */
type book struct {
	// capitalization makes it an exported field/public field
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}