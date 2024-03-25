package main

import (
	"github.com/psyduck-std/psyduck/consume"
	"github.com/psyduck-std/psyduck/produce"
	"github.com/psyduck-std/psyduck/transform"
	"github.com/psyduck-std/sdk"
	"github.com/zclconf/go-cty/cty"
)

func Plugin() *sdk.Plugin {
	return &sdk.Plugin{
		Name: "psyduck",
		Resources: []*sdk.Resource{
			{
				Name:            "psyduck-constant",
				Kinds:           sdk.PRODUCER,
				ProvideProducer: produce.Constant,
				Spec: sdk.SpecMap{
					"value": &sdk.Spec{
						Name:        "value",
						Description: "constant value to produce",
						Type:        cty.String,
						Default:     cty.StringVal("0"),
					},
					"stop-after": &sdk.Spec{
						Name:        "stop-after",
						Description: "stop after n iterations",
						Type:        cty.Number,
						Default:     cty.NumberIntVal(0),
					},
				},
			},
			{
				Name:            "psyduck-trash",
				Kinds:           sdk.CONSUMER,
				ProvideConsumer: consume.Trash,
			},
			{
				Name:               "psyduck-inspect",
				Kinds:              sdk.TRANSFORMER,
				ProvideTransformer: transform.Inspect,
				Spec: sdk.SpecMap{
					"be-string": &sdk.Spec{
						Name:        "be-string",
						Description: "should the data bytes should be a string",
						Type:        cty.Bool,
						Default:     cty.BoolVal(true),
					},
				},
			},
			{
				Name:               "psyduck-snippet",
				Kinds:              sdk.TRANSFORMER,
				ProvideTransformer: transform.Snippet,
				Spec: sdk.SpecMap{
					"fields": &sdk.Spec{
						Name:        "fields",
						Description: "fields to take a snippet of",
						Type:        cty.List(cty.String),
						Required:    true,
					},
				},
			},
			{
				Name:               "psyduck-zoom",
				Kinds:              sdk.TRANSFORMER,
				ProvideTransformer: transform.Zoom,
				Spec: sdk.SpecMap{
					"field": &sdk.Spec{
						Name:        "field",
						Description: "field to zoom into",
						Type:        cty.String,
						Required:    true,
					},
				},
			},
			{
				Name:               "psyduck-wait",
				Kinds:              sdk.TRANSFORMER,
				ProvideTransformer: transform.Wait,
				Spec: sdk.SpecMap{
					"milliseconds": &sdk.Spec{
						Name:        "milliseconds",
						Description: "duratin to wait in ms",
						Type:        cty.Number,
						Required:    true,
					},
				},
			},
		},
	}
}
