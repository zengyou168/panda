package utils

import (
	"github.com/google/uuid"
	"strings"
)

// UUID 生成没有破折号的 UUID
func UUID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}
