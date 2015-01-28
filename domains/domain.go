package domains

import "gopkg.in/mgo.v2"

/*
Each domain is composed of a first name, last name,
email, age, and short message. When represented in
JSON, ditch TitleCase for snake_case.
*/
type Domain struct {
	DomainName string `json:"domain_name"`
}

/*
I want to make sure all these fields are present. The message
is optional, but if it's present it has to be less than
140 characters--it's a short blurb, not your life story.
*/
func (domain *Domain) valid() bool {
	return len(domain.DomainName) > 0
}

/*
I'll use this method when displaying all domains for
"GET /domains". Consult the mgo docs for more info:
http://godoc.org/labix.org/v2/mgo
*/
func fetchAllDomains(db *mgo.Database) []Domain {
	domains := []Domain{}
	err := db.C("domains").Find(nil).All(&domains)
	if err != nil {
		panic(err)
	}

	return domains
}
