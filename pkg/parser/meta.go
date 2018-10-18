package parser

// Meta is a container for registry grid.
type Meta struct {
	Page    string `xml:"page"`
	Total   string `xml:"total"`
	Records string `xml:"records"`
	Rows    []Row  `xml:"row"`
}

// Row is registry record in grid.
type Row struct {
	ID string `xml:"id,attr"`
}
