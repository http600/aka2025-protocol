package profile

type Link struct {
    URL        string `json:"url,omitempty"`
    Title      string `json:"title,omitempty"`
    PlatCode   int    `json:"plat_code,omitempty"`
    PlatName   string `json:"plat_name,omitempty"`
    SecretCode string `json:"secret_code,omitempty"`
}

type ListLinksResponse struct {
    Links   []Link `json:"links,omitempty"`
    Status  int    `json:"status,omitempty"`
    Message string `json:"message,omitempty"`
}
