package main

// FakePost is struct
type FakePost struct {
	ID      int
	Content string
	Author  string
}

func (post *FakePost) fetch(id int) (err error) {
	post.ID = id
	return
}

func (post *FakePost) create() (err error) {
	return
}

func (post *FakePost) update() (err error) {
	return
}

func (post *FakePost) delete() (err error) {
	return
}
