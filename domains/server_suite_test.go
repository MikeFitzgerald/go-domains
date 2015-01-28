// domains/server_suite_test.go

package domains_test

import (
    . "github.com/MikeFitzgerald/go-domains/domains"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"

    "fmt"
    "testing"
)

func TestDomains(t *testing.T) {
    defineFactories()
    RegisterFailHandler(Fail)
    RunSpecs(t, "Domains Suite")
}

/*
Define two factories: one for a valid domain,
and one for an invalid one (too short).
*/
func defineFactories() {
    gory.Define("domain", Domain{},
        func(factory gory.Factory) {
            factory["DomainName"] = "slashthought.com"
        })

    gory.Define("domainTooShort", Domain{},
        func(factory gory.Factory) {
            factory["DomainName"] = "ab"
        })
}