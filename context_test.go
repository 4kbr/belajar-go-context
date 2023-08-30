package belajar_go_context

import (
	"context"
	"fmt"
	"testing"
)

// /context
func TestContext(t *testing.T) {
	//buat context secara manual
	background := context.Background()
	fmt.Println(background)
}
