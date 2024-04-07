package model

type Restaurant struct {
	// ID           primitive.ObjectID `bson:"_id"` // cannot inser when having this fiel
	Name         string
	RestaurantId string `bson:"restaurant_id"`
	Cuisine      string
	Address      interface{}
	Borough      string
	Grades       []interface{}
}

func (r *Restaurant) DbName() string {
	return "sample_restaurants"
}

func (r *Restaurant) CollName() string {
	return "restaurants"
}
