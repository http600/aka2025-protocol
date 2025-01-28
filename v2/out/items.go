package out

import "aka2025-protocol.go/aka2025/protocol/v2"

type Items struct {
    Total int             `json:"total"`
    Items []protocol.Item `json:"items"`
}
