package root_image_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMethods(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RootImage Suite")
}
