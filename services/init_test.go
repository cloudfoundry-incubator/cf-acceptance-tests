package services

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	ginkgoconfig "github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	. "github.com/vito/cmdtest/matchers"

	. "github.com/cloudfoundry/cf-acceptance-tests/helpers"
	"github.com/cloudfoundry/cf-acceptance-tests/config"
	. "github.com/pivotal-cf-experimental/cf-test-helpers/cf"
)

var IntegrationConfig = config.Load()

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)

	AsUser(RegularUserContext, func () {
		Expect(Cf("target",
			"-o", RegularUserContext.Org,
			"-s", RegularUserContext.Space)).To(ExitWith(0))

		RunSpecsWithDefaultAndCustomReporters(t, "Services", []Reporter{reporters.NewJUnitReporter(fmt.Sprintf("junit_%d.xml", ginkgoconfig.GinkgoConfig.ParallelNode))})
	})
}
