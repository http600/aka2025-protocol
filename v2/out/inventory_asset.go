package out

type ViewAssetResponse struct {
    Total  int     `json:"total,omitempty"`
    Assets []Asset `json:"assets,omitempty"`
}

type Asset struct {
    ImgLink      string `json:"img_link,omitempty"`
    NameDisplay  string `json:"name_display,omitempty"`
    TypeDisplay  string `json:"type_display,omitempty"`
    ExteriorWear string `json:"exterior_wear,omitempty"`
}
