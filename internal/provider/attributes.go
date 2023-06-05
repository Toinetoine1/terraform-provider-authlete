package provider

import (
	authlete "github.com/authlete/openapi-for-go/v2"
	authlete3 "github.com/authlete/openapi-for-go/v3"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func createAttributeSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"key": {Type: schema.TypeString,
					Required: true},
				"value": {Type: schema.TypeString,
					Required: true},
			},
		},
	}
}

func mapAttributesToDTO(entry []interface{}) []authlete.Pair {
	var entries = make([]authlete.Pair, 0)

	if entry != nil {
		for _, v := range entry {
			var keypair = v.(map[string]interface{})
			newPair := authlete.NewPair()
			newPair.SetKey(keypair["key"].(string))
			newPair.SetValue(keypair["value"].(string))
			entries = append(entries, *newPair)
		}
	}
	return entries
}

func mapAttributesToDTOV3(entry []interface{}) []authlete3.Pair {
	var entries = make([]authlete3.Pair, 0)

	if entry != nil {
		for _, v := range entry {
			var keypair = v.(map[string]interface{})
			newPair := authlete3.NewPair()
			newPair.SetKey(keypair["key"].(string))
			newPair.SetValue(keypair["value"].(string))
			entries = append(entries, *newPair)
		}
	}
	return entries
}

func mapAttributesFromDTO(pairs []authlete.Pair) []interface{} {

	if pairs != nil {
		entries := make([]interface{}, len(pairs))

		for i, v := range pairs {
			newEntry := make(map[string]interface{})
			newEntry["key"] = v.Key
			newEntry["value"] = v.Value
			entries[i] = newEntry
		}
		return entries
	}
	return make([]interface{}, 0)
}

func mapAttributesFromDTOV3(pairs []authlete3.Pair) []interface{} {

	if pairs != nil {
		entries := make([]interface{}, len(pairs))

		for i, v := range pairs {
			newEntry := make(map[string]interface{})
			newEntry["key"] = v.Key
			newEntry["value"] = v.Value
			entries[i] = newEntry
		}
		return entries
	}
	return make([]interface{}, 0)
}
