package mutations

import (
	 "github.com/NelsonMario/graphql/query/type"
	 "github.com/NelsonMario/graphql/resolvers"
	"github.com/graphql-go/graphql"
)

func GetRoot() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createAdmin": &graphql.Field{
				Type: types.GetAdminType(),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.InsertAdmin,
			},
			"updateAdmin": &graphql.Field{
				Type: types.GetAdminType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.UpdateAdmin,
			},
			"removeAdmin": &graphql.Field{
				Type: types.GetAdminType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.RemoveAdmin,
			},
			"insertSlider": &graphql.Field{
				Type: types.GetSliderType(),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.InsertSlider,
			},

			"insertUser": &graphql.Field{
				Type: types.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"firstName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"lastName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"phoneNumber": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.InsertUser,
			},
			"updateUser": &graphql.Field{
				Type: types.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"first_name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"last_name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"phone_number": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"city_name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"address": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"post_code": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"language": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.UpdateUser,
			},
			"insertAirline": &graphql.Field{
				Type: types.GetAirlineType(),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:resolvers.InsertAirline,
			},

			"insertFlight": &graphql.Field{
				Type: types.GetFlightType(),
				Args: graphql.FieldConfigArgument{
					"airlineRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"fromRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"toRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"departure": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"arrival": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"tax": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"serviceCharge": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:resolvers.InsertFlight,
			},

			"insertRailroad": &graphql.Field{
				Type: types.GetRailroadType(),
				Args: graphql.FieldConfigArgument{
					"trainRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"fromRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"toRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"departure": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"arrival": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"tax": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"serviceCharge": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:resolvers.InsertRailroad,
			},
			"updateRailroad": &graphql.Field{
				Type: types.GetRailroadType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"train_refer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"from_refer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"to_refer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"departure": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"arrival": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"tax": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"service_charge": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:resolvers.UpdateRailroad,
			},

			"removeRailroad": &graphql.Field{
				Type: types.GetRailroadType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:resolvers.RemoveRailroad,
			},

			"removeFlight": &graphql.Field{
				Type: types.GetFlightType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:resolvers.RemoveFlight,
			},
			"updateFlight": &graphql.Field{
				Type: types.GetFlightType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"airline_refer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"from_refer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"to_refer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"departure": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"arrival": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"tax": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"service_charge": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:resolvers.UpdateFlight,
			},

			"insertHotel": &graphql.Field{
				Type: types.GetHotelType(),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"rating": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"type": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"locationRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"hotelLat": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
					"hotelLng": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
				},
				Resolve:resolvers.InsertHotel,
			},
			"updateHotel": &graphql.Field{
				Type: types.GetHotelType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"rating": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:resolvers.UpdateHotel,
			},
			"removeHotel": &graphql.Field{
				Type: types.GetHotelType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:resolvers.RemoveHotel,
			},

			"insertBlog": &graphql.Field{
				Type: types.GetBlogType(),
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"value": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"thumbnail": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"viewer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:resolvers.InsertBlog,
			},
			"updateBlog": &graphql.Field{
				Type: types.GetBlogType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"thumbnail": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"value": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:resolvers.UpdateBlog,
			},
			"removeBlog": &graphql.Field{
				Type: types.GetBlogType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:resolvers.RemoveBlog,
			},
			"insertEvent": &graphql.Field{
				Type: types.GetEventType(),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"locationRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"startDate": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"endDate": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"eventLat": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
					"eventLng": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
					"category": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:resolvers.InsertEvent,
			},
			"updateEvent": &graphql.Field{
				Type: types.GetEventType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"locationRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"startDate": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"endDate": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
				},
				Resolve:resolvers.UpdateEvent,
			},
			"removeEvent": &graphql.Field{
				Type: types.GetEventType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:resolvers.RemoveEvent,
			},
			"insertEventDetail": &graphql.Field{
				Type: types.GetEventDetailType(),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"eventRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:resolvers.InsertEventDetail,
			},
			"updateEventDetail": &graphql.Field{
				Type: types.GetEventDetailType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"eventRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:resolvers.UpdateEventDetail,
			},
			"removeEventDetail": &graphql.Field{
				Type: types.GetEventDetailType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:resolvers.RemoveEventDetail,
			},

			"increaseViewer": &graphql.Field{
				Type: types.GetBlogType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"viewer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:resolvers.IncreaseViewer,
			},
		},
	})
}

