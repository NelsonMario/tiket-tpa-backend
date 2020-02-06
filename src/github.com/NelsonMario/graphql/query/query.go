package query

import (
	"github.com/NelsonMario/graphql/query/type"
	"github.com/NelsonMario/graphql/resolvers"
	"github.com/graphql-go/graphql"
)

func GetRoot() *graphql.Object {

	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{

			"admins": {
				Type:        graphql.NewList(types.GetAdminType()),
				Resolve:     resolvers.GetAllAdmin,
				Description: "Get All Admins",
			},
			"admin": {
				Type: graphql.NewList(types.GetAdminType()),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:     resolvers.GetAdmin,
				Description: "Get Single Users",
			},

			"sliders": {
				Type:        graphql.NewList(types.GetSliderType()),
				Resolve:     resolvers.GetAllSlider,
				Description: "Get All Slider",
			},

			"userById": {
				Type: graphql.NewList(types.GetUserType()),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:     resolvers.GetUserById,
				Description: "Get Single Users by Id",
			},

			"userByEmailOrPhone": {
				Type: graphql.NewList(types.GetUserType()),
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolvers.GetUserByEmailOrPhone,
				Description: "Get Single Users by Email or Phone",
			},

			"users": {
				Type:        graphql.NewList(types.GetUserType()),
				Resolve:     resolvers.GetAllUser,
				Description: "Get All Users",
			},

			"airlines": {
					Type: graphql.NewList(types.GetAirlineType()),
					Resolve: resolvers.GetAllAirline,
			},

			"airlineById": {
				Type: graphql.NewList(types.GetAirlineType()),
				Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.GetAirlineById,
				Description: "Get airline by id",
			},

			"airports": {
				Type:        graphql.NewList(types.GetAirportType()),
				Resolve:     resolvers.GetAllAirport,
				Description: "Get All Airport",
			},

			"distinctAirports": {
				Type:        graphql.NewList(types.GetAirportType()),
				Resolve:     resolvers.GetDistinctAirport,
				Description: "Get Distinct Airport",
			},

			"flights": {
				Type:	graphql.NewList(types.GetFlightType()),
				Resolve: resolvers.GetAllFlight,
			},

			"flightByFromAndTo": {
				Type:	graphql.NewList(types.GetFlightType()),
				Args:	graphql.FieldConfigArgument{
					"from": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"to": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.GetFlightByToAndFrom,
			},

			"flightRoutes": {
				Type:	graphql.NewList(types.GetFlightRoutesType()),
				Resolve: resolvers.GetAllFlightRoutes,
			},

			"flightFacilities": {
				Type:	graphql.NewList(types.GetFlightFacilityType()),
				Resolve: resolvers.GetAllFlightFacility,
			},

			"facilities": {
				Type:	graphql.NewList(types.GetFacilityType()),
				Resolve: resolvers.GetAllFacility,
			},
		},
	})

}

