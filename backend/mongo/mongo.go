package mongo

type Database struct {
	Collections(string) Collection
	Client() Client
}

type Collection interface {

}

type SingleResult interface {

}

type Cursor interface {

}

type mongoClient struct {

}

type mongoCursor struct {

}

type mongoDatabase struct {

}
type mongoCollection struct {

}

type mongoSingleResult struct {

}

type mongoSession struct {

}

type nullwareDecoder struct {
	
}
