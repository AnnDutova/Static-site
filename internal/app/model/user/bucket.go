package model

type Bucket struct {
	Bucket []BucketElem
}

type BucketElem struct {
	Author string
	Song   string
	Count  int
	Price  int
	Salon  string
}
