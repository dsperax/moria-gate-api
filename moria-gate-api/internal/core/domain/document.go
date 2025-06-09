package domain

// Document represents a sanitized CPF or 8-digit CNPJ
type Document struct {
	Value string
}

// NewDocument creates a new Document instance
func NewDocument(value string) Document {
	return Document{Value: value}
}

// IsValid checks if the document is a valid CPF or short CNPJ
func (d Document) IsValid() bool {
	length := len(d.Value)
	return length == 11 || length == 8
}

// Key returns a consistent key for document lookups
func (d Document) Key() string {
	return d.Value
}
