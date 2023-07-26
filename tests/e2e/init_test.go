package e2e

import (
	"context"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var ctx context.Context
var token string

func TestRHCSProvider(t *testing.T) {
	token = os.Getenv("RHCS_TOKEN")
	ctx = context.Background()
	RegisterFailHandler(Fail)
	RunSpecs(t, "RHCS Provider Test")

}
