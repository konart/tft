package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/konart/tft/models"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Thing)
// DB Table: Plural (things)
// Resource: Plural (Things)
// Path: Plural (/things)
// View Template Folder: Plural (/templates/things/)

// ThingsResource is the resource for the thing model
type ThingsResource struct {
	buffalo.Resource
}

// List gets all Things. This function is mapped to the path
// GET /things
func (v ThingsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	things := &models.Things{}
	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=2".
	q := tx.PaginateFromParams(c.Params())
	// You can order your list here. Just change
	err := q.All(things)
	// to:
	// err := q.Order("create_at desc").All(things)
	if err != nil {
		return errors.WithStack(err)
	}
	// Make Things available inside the html template
	c.Set("things", things)
	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
	return c.Render(200, r.HTML("things/index.html"))
}

// Show gets the data for one Thing. This function is mapped to
// the path GET /things/{thing_id}
func (v ThingsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty Thing
	thing := &models.Thing{}
	// To find the Thing the parameter thing_id is used.
	err := tx.Find(thing, c.Param("thing_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	// Make thing available inside the html template
	c.Set("thing", thing)
	return c.Render(200, r.HTML("things/show.html"))
}

// New renders the formular for creating a new Thing.
// This function is mapped to the path GET /things/new
func (v ThingsResource) New(c buffalo.Context) error {
	// Make thing available inside the html template
	c.Set("thing", &models.Thing{})
	return c.Render(200, r.HTML("things/new.html"))
}

// Create adds a Thing to the DB. This function is mapped to the
// path POST /things
func (v ThingsResource) Create(c buffalo.Context) error {
	// Allocate an empty Thing
	thing := &models.Thing{}
	// Bind thing to the html form elements
	err := c.Bind(thing)
	if err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(thing)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		// Make thing available inside the html template
		c.Set("thing", thing)
		// Make the errors available inside the html template
		c.Set("errors", verrs)
		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("things/new.html"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "Thing was created successfully")
	// and redirect to the things index page
	return c.Redirect(302, "/things/%s", thing.ID)
}

// Edit renders a edit formular for a thing. This function is
// mapped to the path GET /things/{thing_id}/edit
func (v ThingsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty Thing
	thing := &models.Thing{}
	err := tx.Find(thing, c.Param("thing_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	// Make thing available inside the html template
	c.Set("thing", thing)
	return c.Render(200, r.HTML("things/edit.html"))
}

// Update changes a thing in the DB. This function is mapped to
// the path PUT /things/{thing_id}
func (v ThingsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty Thing
	thing := &models.Thing{}
	err := tx.Find(thing, c.Param("thing_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	// Bind Thing to the html form elements
	err = c.Bind(thing)
	if err != nil {
		return errors.WithStack(err)
	}
	verrs, err := tx.ValidateAndUpdate(thing)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		// Make thing available inside the html template
		c.Set("thing", thing)
		// Make the errors available inside the html template
		c.Set("errors", verrs)
		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("things/edit.html"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "Thing was updated successfully")
	// and redirect to the things index page
	return c.Redirect(302, "/things/%s", thing.ID)
}

// Destroy deletes a thing from the DB. This function is mapped
// to the path DELETE /things/{thing_id}
func (v ThingsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty Thing
	thing := &models.Thing{}
	// To find the Thing the parameter thing_id is used.
	err := tx.Find(thing, c.Param("thing_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	err = tx.Destroy(thing)
	if err != nil {
		return errors.WithStack(err)
	}
	// If there are no errors set a flash message
	c.Flash().Add("success", "Thing was destroyed successfully")
	// Redirect to the things index page
	return c.Redirect(302, "/things")
}
