package pkg

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func SpecialKode() string {
	id := uuid.New()
	idUpper := strings.ToUpper(id.String()[:4])
	uuid := fmt.Sprintf("ANDL%s", idUpper)
	return uuid
}
