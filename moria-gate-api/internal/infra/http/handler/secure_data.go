package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"moria-gate-api/internal/core/domain"
	"moria-gate-api/internal/core/services"
	"moria-gate-api/internal/infra/crypto"
	"moria-gate-api/internal/infra/hash"

	"github.com/gin-gonic/gin"
)

type SecureHandler struct {
	Lookup *services.LookupService
}

type request struct {
	Document string `json:"document"`
}

type response struct {
	Hash        string `json:"hash"`
	Encrypted   string `json:"encrypted"`
	Description string `json:"description,omitempty"`
}

func NewSecureHandler(service *services.LookupService) *SecureHandler {
	return &SecureHandler{Lookup: service}
}

func (h *SecureHandler) SecureData(c *gin.Context) {
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	doc := domain.NewDocument(req.Document)
	if !doc.IsValid() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid document"})
		return
	}

	client, err := h.Lookup.GetSecureClientData(context.Background(), doc)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "client not found"})
		return
	}

	// Marshal client info
	raw, _ := json.Marshal(client)

	// Encrypt using AES-GCM
	key := []byte(os.Getenv("API_SECRET"))
	encrypted, err := crypto.EncryptAESGCM(raw, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "encryption failed"})
		return
	}

	// Generate SHA256 of document
	hashed := hash.SHA256Hex(doc.Key())

	c.JSON(http.StatusOK, response{
		Hash:        hashed,
		Encrypted:   encrypted,
		Description: "Data is hashed and encrypted using AES-GCM",
	})
}
