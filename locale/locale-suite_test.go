package locale_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2" //nolint:revive // ginkgo ok
	. "github.com/onsi/gomega"    //nolint:revive // gomega ok
)

func TestLocale(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Locale Suite")
}