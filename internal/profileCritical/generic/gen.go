package generic

import (
	"strings"
	"github.com/google/uuid"
)

func GenerateCSRFToken() string {
   csrfToken := uuid.New().String()
   return strings.ReplaceAll(csrfToken, "-", "")
}