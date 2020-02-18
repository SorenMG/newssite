package sites

type Site struct {
	Id        int           `json:"id"`
	Name      string        `json:"name"`
	Domain    string        `json:"domain"`
	Locale    string        `json:"locale"`
	ScrapeDef SiteScrapeDef `json:"scrapedef"`
}

type SiteScrapeDef struct {
	PublishedTime  string    `json:"pubtime"`
	EditedTime     string    `json:"edtime"`
	Description    string    `json:"desc"`
	Title          string    `json:"title"`
	Tag            string    `json:"tag"`
	Section        string    `json:"section"`
	ImageURL       string    `json:"imgurl"`
	ImageDimension Dimension `json:"imgdim"`
}

type Dimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
