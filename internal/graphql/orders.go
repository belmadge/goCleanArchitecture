package graphql

import (
    "github.com/graphql-go/graphql"
    "github.com/belmadge/goCleanArchitecture/internal/app/service"
)

func NewSchema(orderService *service.OrderService) (graphql.Schema, error) {
    orderType := graphql.NewObject(graphql.ObjectConfig{
        Name: "Order",
        Fields: graphql.Fields{
            "id": &graphql.Field{
                Type: graphql.Int,
            },
            "customer_name": &graphql.Field{
                Type: graphql.String,
            },
            "total": &graphql.Field{
                Type: graphql.Float,
            },
        },
    })

    queryType := graphql.NewObject(graphql.ObjectConfig{
        Name: "Query",
        Fields: graphql.Fields{
            "orders": &graphql.Field{
                Type: graphql.NewList(orderType),
                Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                    return orderService.ListOrders()
                },
            },
        },
    })

    schemaConfig := graphql.SchemaConfig{Query: queryType}
    return graphql.NewSchema(schemaConfig)
}
