package build_asset

import (
	"github.com/mahdi-cpp/photokit/internal/update"
)

type Comment struct {
	ID    int
	Title string
	Count int
}

type Like struct {
	ID    int
	Title string
	Count int
}

type Person struct {
	ID       int
	Name     string
	Email    string
	Groups   []string
	Comments []Comment
	Likes    []Like
}

type PersonUpdate struct {
	ID    int
	Name  *string
	Email *string

	Groups       *[]string
	AddGroups    []string
	RemoveGroups []string

	Comments       *[]Comment
	AddComments    []Comment
	RemoveComments []Comment

	Likes       *[]Like
	AddLikes    []Like
	RemoveLikes []Like

	CommentUpdates []update.NestedFieldUpdate[Comment]
	LikeUpdates    []update.NestedFieldUpdate[Like]
}

// Key extractors for nested structs
func commentKeyExtractor(c Comment) int { return c.ID }
func likeKeyExtractor(l Like) int       { return l.ID }

// Initialize updater
var personUpdater = update.NewUpdater[Person, PersonUpdate]()

func init() {
	// Scalar fields
	personUpdater.AddScalarUpdater(func(p *Person, u PersonUpdate) {
		if u.Name != nil {
			p.Name = *u.Name
		}
	})

	personUpdater.AddScalarUpdater(func(p *Person, u PersonUpdate) {
		if u.Email != nil {
			p.Email = *u.Email
		}
	})

	// Groups (basic slice)
	personUpdater.AddCollectionUpdater(func(p *Person, u PersonUpdate) {
		op := update.CollectionUpdateOp[string]{
			FullReplace: u.Groups,
			Add:         u.AddGroups,
			Remove:      u.RemoveGroups,
		}
		p.Groups = update.ApplyCollectionUpdate(p.Groups, op)
	})

	// Comments (ID-based updates)
	personUpdater.AddNestedUpdater(func(p *Person, u PersonUpdate) {
		op := update.CollectionUpdateOp[Comment]{
			FullReplace: u.Comments,
			Add:         u.AddComments,
			Remove:      u.RemoveComments,
		}
		p.Comments = update.ApplyCollectionUpdateByID(
			p.Comments,
			op,
			commentKeyExtractor,
		)

		// Apply field-level updates to existing comments
		p.Comments = update.ApplyNestedUpdate(
			p.Comments,
			u.CommentUpdates,
			commentKeyExtractor,
		)
	})

	// Likes (ID-based updates)
	personUpdater.AddNestedUpdater(func(p *Person, u PersonUpdate) {
		op := update.CollectionUpdateOp[Like]{
			FullReplace: u.Likes,
			Add:         u.AddLikes,
			Remove:      u.RemoveLikes,
		}
		p.Likes = update.ApplyCollectionUpdateByID(
			p.Likes,
			op,
			likeKeyExtractor,
		)

		// Apply field-level updates to existing likes
		p.Likes = update.ApplyNestedUpdate(
			p.Likes,
			u.LikeUpdates,
			likeKeyExtractor,
		)
	})
}

func UpdatePerson(p *Person, update PersonUpdate) *Person {
	personUpdater.Apply(p, update)
	return p
}
