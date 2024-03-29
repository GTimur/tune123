package tune123

type CatalogTree struct {
	Files []CatalogFile
}

type CatalogFile struct {
	ID       int64 `json:"key"`
	ParentID int64
	IsDir    bool          `json:"folder,omitempty"`
	Name     string        `json:"title"`
	Children []CatalogFile `json:"children"`
}

func (c *CatalogTree) TreeJSON(cat Catalog) error {

	c.Files = c.Files[:0]

	var t CatalogFile

	for _, r := range cat.Files {
		if !r.IsDir {
			t.ID = r.ID
			t.IsDir = r.IsDir
			t.Name = r.Name
			t.Children = []CatalogFile{}
			c.Files = append(c.Files, t)
			continue
		}
		t.ID = r.ID
		t.IsDir = r.IsDir
		t.Name = r.Name
		t.Children = append(t.Children, t)
		c.Files = append(c.Files, t)
		if err := c.TreeJSON(cat); err != nil {
			return err
		}

		c.Files = append(c.Files, t)

	}

	return nil
}

func (c *CatalogTree) TreeChilds(cat Catalog) error {

	c.Files = c.Files[:0]

	var t CatalogFile

	for _, r := range cat.Files {
		if !r.IsDir {
			t.ID = r.ID
			t.IsDir = r.IsDir
			t.Name = r.Name
			t.Children = []CatalogFile{}
			c.Files = append(c.Files, t)
			continue
		}
		t.ID = r.ID
		t.IsDir = r.IsDir
		t.Name = r.Name
		t.Children = append(t.Children, t)
		c.Files = append(c.Files, t)
		if err := c.TreeJSON(cat); err != nil {
			return err
		}

		c.Files = append(c.Files, t)

	}

	return nil
}
