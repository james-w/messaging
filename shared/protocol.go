// This can define structured message types, JSON encoding/decoding, etc.
// Currently unused but important for future extensibility.
package shared

type Message struct {
	Sender    string `json:"sender"`
	Timestamp string `json:"timestamp"`
	Text      string `json:"text"`
}

