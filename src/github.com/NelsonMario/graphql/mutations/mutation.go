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
					"serviceTax": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:resolvers.InsertFlight,
			},
		},
	})

}

