package out

type Items struct {
    Total int     `json:"total"`
    Items []Items `json:"items"`
}
