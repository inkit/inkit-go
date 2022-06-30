package render

type CreateRenderOptions struct {
	FileName          string            `json:"-"`
	Html              string            `json:"html"`
	Width             float64           `json:"width"`
	Height            float64           `json:"height"`
	FolderId          string            `json:"folder_id,omitempty"`
	Unit              string            `json:"unit"`
	MergeParameters   map[string]string `json:"merge_parameters,omitempty"`
	ExpireAfterNViews int               `json:"expire_after_n_views,omitempty"`
	TemplateId        string            `json:"template_id,omitempty"`
}

type RendersList struct {
	Items    []Render `json:"items"`
	Metadata Metadata `json:"metadata"`
}

type Render struct {
	Id                string            `json:"id"`
	Html              string            `json:"html"`
	TemplateId        string            `json:"template_id"`
	FolderId          string            `json:"folder_id"`
	Status            string            `json:"status"`
	Width             float64           `json:"width"`
	Height            float64           `json:"height"`
	Unit              string            `json:"unit"`
	MergeParameters   map[string]string `json:"merge_parameters"`
	PdfUrl            string            `json:"pdf_url"`
	HtmlUrl           string            `json:"html_url"`
	ExpireAfterNViews int               `json:"expire_after_n_views"`
	ExpireAt          string            `json:"expire_at"`
	UpdatedAt         string            `json:"updated_at"`
	CreatedAt         string            `json:"created_at"`
}

type Metadata struct {
	Pagination Pagination `json:"pagination"`
	Sort       Sort       `json:"sort"`
}

type Pagination struct {
	PageSize    int `json:"page_size"`
	CurrentPage int `json:"current_page"`
	NextPage    int `json:"next_page"`
	PrevPage    int `json:"prev_page"`
	PageCount   int `json:"page_count"`
	TotalCount  int `json:"total_count"`
}

type Sort struct {
	Key   string `json:"key"`
	Order int    `json:"order"`
}

type RenderListOptions struct {
	Page     int
	PageSize int
}
