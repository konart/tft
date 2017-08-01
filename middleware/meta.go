package middleware

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gobuffalo/buffalo"
	"github.com/konart/tft/models"
	"github.com/pkg/errors"
)

func GetMetaInfo(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if c.Request().URL.Path == "/things" && c.Request().Method == "POST" {
			thing := &models.Thing{}

			err := c.Bind(thing)
			if err != nil {
				return errors.WithStack(err)
			}

			doc, err := goquery.NewDocument(thing.Content)
			if err == nil {
				var title string
				doc.Find("head title").Each(func(i int, s *goquery.Selection) {
					title = s.Text()
				})
				c.Request().Form.Set("Title", title)
			}
		}
		err := next(c)
		return err
	}
}
