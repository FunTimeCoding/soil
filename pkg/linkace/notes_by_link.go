package linkace

import "github.com/funtimecoding/soil/pkg/linkace/note"

func (c *Client) NotesByLink(linkIdentifier int) ([]*note.Note, error) {
	var result []*note.Note
	p := 1

	for {
		r, e := c.NotesByLinkPage(linkIdentifier, p)

		if e != nil {
			return nil, e
		}

		result = append(result, r.Items...)

		if p >= r.LastPage {
			break
		}

		p++
	}

	return result, nil
}
