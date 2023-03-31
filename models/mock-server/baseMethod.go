package mock_server

import "github.com/google/uuid"

type MockResponse struct {
	Path       string `json:"path"`
	Method     string `json:"method"`
	StatusCode int    `json:"status_code"`

	CollectionId uuid.UUID   `json:"collection_id"`
	ResponseBody interface{} `json:"response_body"`
}
type InitJsonFile struct {
	Get    []MockResponse `json:"GET"`
	Post   []MockResponse `json:"POST"`
	Put    []MockResponse `json:"PUT"`
	Patch  []MockResponse `json:"PATCH"`
	Delete []MockResponse `json:"DELETE"`
}
