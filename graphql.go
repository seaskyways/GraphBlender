package GraphBlender

import (
	"github.com/graphql-go/graphql"
	"github.com/rocketlaunchr/dataframe-go"
	"strconv"
	"strings"
)

type GraphBlender struct {
	collection DataFrameCollection
}

func New(collection DataFrameCollection) *GraphBlender {
	return &GraphBlender{collection: collection}
}

func (gb GraphBlender) Blend() (graphql.Schema, error) {
	fields := graphql.Fields{}
	for _, entry := range gb.collection {
		fields[entry.Name] = gb.blendDataframe(entry)
	}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "RootQuery",
			Interfaces:  nil,
			Fields:      fields,
			IsTypeOf:    nil,
			Description: "",
		}),
	})

	return schema, err
}

func (gb GraphBlender) blendDataframe(dfe DataFrameEntry) *graphql.Field {
	subFields := graphql.Fields{}
	cols := dfe.DataFrame.Names()
	for _, col := range cols {
		col := col
		if col == "" {
			continue
		}
		fixed := fixGraphQLName(col)
		subFields[fixed] = &graphql.Field{
			Name: fixed,
			Type: graphql.String,
			Args: nil,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				m := p.Source.(map[interface{}]interface{})
				return m[col], nil
			},
			DeprecationReason: "",
			Description:       "",
		}
	}

	field := graphql.Field{
		Name: dfe.Name,
		Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
			Name:        dfe.Name,
			Interfaces:  nil,
			Fields:      subFields,
			IsTypeOf:    nil,
			Description: "Dataset: " + dfe.Name,
		})),
		Args: graphql.FieldConfigArgument{
			"limit": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: 100,
			},
			"skip": &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 0},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			limit := p.Args["limit"].(int)
			skip := p.Args["skip"].(int)

			var out []map[interface{}]interface{}
			iterator := dfe.DataFrame.ValuesIterator(dataframe.ValuesOptions{
				InitialRow: skip,
				Step:       1,
			})

			for i := 0; i < limit; i++ {
				_, row, _ := iterator(dataframe.SeriesName)
				if row == nil {
					break
				}
				out = append(out, row)
			}

			return out, nil
		},
		DeprecationReason: "",
		Description:       "",
	}

	return &field
}

func fixGraphQLName(in string) (out string) {
	out = in
	if _, err := strconv.Atoi(out); err == nil {
		out = "_" + out
	}
	out = strings.ReplaceAll(out, " ", "")
	out = strings.TrimSpace(out)
	return out
}
