// domains/server_test.go
package server_test

import (
    . "github.com/MikeFitzgerald/go-domains/domains"

    "bytes"
    "encoding/json"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "net/http"
    "net/http/httptest"
)

/*
Convert JSON data into a slice.
*/
func sliceFromJSON(data []byte) []interface{} {
    var result interface{}
    json.Unmarshal(data, &result)
    return result.([]interface{})
}

/*
Convert JSON data into a map.
*/
func mapFromJSON(data []byte) map[string]interface{} {
    var result interface{}
    json.Unmarshal(data, &result)
    return result.(map[string]interface{})
}

/*
Server unit tests.
*/
var _ = Describe("Server", func() {
    var dbName string
    var session *DatabaseSession
    var server Server
    var request *http.Request
    var recorder *httptest.ResponseRecorder

    BeforeEach(func() {
        // Set up a new server, connected to a test database,
        // before each test.
        dbName = "domains_test"
        session = NewSession(dbName)
        server = NewServer(session)

        // Record HTTP responses.
        recorder = httptest.NewRecorder()
    })

    AfterEach(func() {
        // Clear the database after each test.
        session.DB(dbName).DropDatabase()
    })

    Describe("GET /domains", func() {

        // Set up a new GET request before every test
        // in this describe block.
        BeforeEach(func() {
            request, _ = http.NewRequest("GET", "/domains", nil)
        })

        Context("when no domains exist", func() {
            It("returns a status code of 200", func() {
                server.ServeHTTP(recorder, request)
                Expect(recorder.Code).To(Equal(200))
            })

            It("returns a null body", func() {
                server.ServeHTTP(recorder, request)
                Expect(recorder.Body.String()).To(Equal("[]"))
            })
        })

        Context("when domains exist", func() {

            // Insert two valid signatures into the database
            // before each test in this context.
            BeforeEach(func() {
                collection := session.DB(dbName).C("domains")
                collection.Insert(gory.Build("domain"))
                collection.Insert(gory.Build("domain"))
            })

            It("returns a status code of 200", func() {
                server.ServeHTTP(recorder, request)
                Expect(recorder.Code).To(Equal(200))
            })

            It("returns those domains in the body", func() {
                server.ServeHTTP(recorder, request)

                domainJSON := sliceFromJSON(recorder.Body.Bytes())
                Expect(len(peopleJSON)).To(Equal(2))

                domainJSON := peopleJSON[0].(map[string]interface{})
                Expect(personJSON["domain_name"]).To(Equal("slashthought.com"))
            })
        })
    })

    Describe("POST /domains", func() {

        Context("with invalid JSON", func() {

            // Create a POST request using JSON from our invalid
            // factory object before each test in this context.
            BeforeEach(func() {
                body, _ := json.Marshal(
                    gory.Build("domainTooShort"))
                request, _ = http.NewRequest(
                    "POST", "/domains", bytes.NewReader(body))
            })

            It("returns a status code of 400", func() {
                server.ServeHTTP(recorder, request)
                Expect(recorder.Code).To(Equal(400))
            })
        })

        Context("with valid JSON", func() {

            // Create a POST request with valid JSON from
            // our factory before each test in this context.
            BeforeEach(func() {
                body, _ := json.Marshal(
                    gory.Build("domain"))
                request, _ = http.NewRequest(
                    "POST", "/domains", bytes.NewReader(body))
            })

            It("returns a status code of 201", func() {
                server.ServeHTTP(recorder, request)
                Expect(recorder.Code).To(Equal(201))
            })

            It("returns the inserted domain", func() {
                server.ServeHTTP(recorder, request)

                personJSON := mapFromJSON(recorder.Body.Bytes())
                Expect(personJSON["domain_name"]).To(Equal("slashthought.com"))
            })
        })
    })
})
