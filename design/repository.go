package design

import (
	d "github.com/goadesign/goa/design"
	a "github.com/goadesign/goa/design/apidsl"
)

var _ = a.Resource("repositories", func() {
	a.BasePath("/repositories")

	a.Action("show", func() {
		a.Routing(
			a.GET("/:id"),
		)
		a.Description("Retrieve repository from the given ID.")
		a.Params(func() {
			a.Param("id", d.String, "id")
		})
		a.Response(d.InternalServerError, JSONAPIErrors)
		a.Response(d.OK)
	})

	a.Action("create", func() {
		a.Security("jwt")
		a.Routing(
			a.POST(""),
		)
		a.Description("create the repository")
		a.Payload(createRepositorySingle)
		a.Response(d.Created)
		a.Response(d.InternalServerError, JSONAPIErrors)
		a.Response(d.BadRequest, JSONAPIErrors)
		a.Response(d.Unauthorized, JSONAPIErrors)
		a.Response(d.Forbidden, JSONAPIErrors)
		a.Response(d.Conflict, JSONAPIErrors)
	})

})

var createRepositorySingle = JSONSingle(
	"CreateRepository", "Holds a single CreateRepository object",
	createRepository,
	nil)

var createRepository = a.Type("CreateRepository", func() {
	a.Description(`JSONAPI store for the data of a CreateRepository object. See also http://jsonapi.org/format/#document-resource-object`)
	a.Attribute("type", d.String, func() {
		a.Enum("createRepositories")
	})
	a.Attribute("attributes", createRepositoryAttributes)
	a.Required("type", "attributes")
})

var createRepositoryAttributes = a.Type("CreateRepositoryDataAttributes", func() {
	a.Attribute("name", d.String, "The name of the repository")
	a.Attribute("description", d.String, "A short description of the repository")
	a.Attribute("homepage", d.String, "A URL with more information about the repository")
	a.Attribute("private", d.Boolean, "Either true to create a private repository or false to create a public one. Creating private repositories requires a paid GitHub account. Default: false")
	a.Attribute("has_issues", d.Boolean, "Either true to enable issues for this repository or false to disable them. Default: true")
	a.Attribute("has_projects", d.Boolean, "Either true to enable projects for this repository or false to disable them. Default: true. Note: If you're creating a repository in an organization that has disabled repository projects, the default is false, and if you pass true, the API returns an error")
	a.Attribute("has_wiki", d.Boolean, "Either true to enable the wiki for this repository or false to disable it. Default: true")
	a.Attribute("team_id", d.Integer, "The id of the team that will be granted access to this repository. This is only valid when creating a repository in an organization")
	a.Attribute("auto_init", d.Boolean, "Pass true to create an initial commit with empty README. Default: false")
	a.Attribute("gitignore_template", d.String, "Desired language or platform .gitignore template to apply. Use the name of the template without the extension")
	a.Attribute("license_template", d.String, "Choose an open source license template that best suits your needs, and then use the license keyword as the license_template string")
	a.Attribute("allow_squash_merge", d.Boolean, "Either true to allow squash-merging pull requests, or false to prevent squash-merging")
	a.Attribute("allow_merge_commit", d.Boolean, "Either true to allow merging pull requests with a merge commit, or false to prevent merging pull requests with merge commits")
	a.Attribute("allow_rebase_merge", d.Boolean, "Either true to allow rebase-merging pull requests, or false to prevent rebase-merging")
	a.Required("name")
})
