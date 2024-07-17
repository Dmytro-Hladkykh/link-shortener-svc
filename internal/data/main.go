package data

type MasterQ interface {
	New() MasterQ

	ShortLink() ShortLinkQ

	Transaction(fn func(db MasterQ) error) error
}