package domains

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

/*
Wrap the Martini server struct.
*/
type Server *martini.ClassicMartini

/*
Create a new *martini.ClassicMartini server.
We'll use a JSON renderer and our MongoDB
database handler. We define two routes:
"GET /domains" and "POST /domains".
*/
func NewServer(session *DatabaseSession) Server {
	// Create the server and set up middleware.
	m := Server(martini.Classic())
	m.Use(render.Renderer(render.Options{
		IndentJSON: true,
	}))
	m.Use(session.Database())

	// Define the "GET /domains" route.
	m.Get("/domains", func(r render.Render, db *mgo.Database) {
		r.JSON(200, fetchAllDomains(db))
	})

	// Define the "POST /domains" route.
	m.Post("/domains", binding.Json(Domain{}),
		func(domain Domain,
			r render.Render,
			db *mgo.Database) {

			if domain.valid() {
				// domain is valid, insert into database
				err := db.C("domains").Insert(domain)
				if err == nil {
					// insert successful, 201 Created
					r.JSON(201, domain)
				} else {
					// insert failed, 400 Bad Request
					r.JSON(400, map[string]string{
						"error": err.Error(),
					})
				}
			} else {
				// domain is invalid, 400 Bad Request
				r.JSON(400, map[string]string{
					"error": "Not a valid domain",
				})
			}
		})

	// Return the server. Call Run() on the server to
	// begin listening for HTTP requests.
	return m
}
