package kong

import "github.com/hashicorp/terraform/helper/schema"

func readArrayFromResource(d *schema.ResourceData, key string) []string {

	if attr, ok := d.GetOk(key); ok {
		array := []string{}
		items := attr.([]interface{})
		for _, x := range items {
			item := x.(string)
			array = append(array, item)
		}

		return array
	}

	return nil
}

func readStringFromResource(d *schema.ResourceData, key string) string {
	if attr, ok := d.GetOk(key); ok {
		return attr.(string)
	}
	return ""
}

func readBoolFromResource(d *schema.ResourceData, key string) bool {
	if attr, ok := d.GetOk(key); ok {
		return attr.(bool)
	}
	return false
}

func readIntFromResource(d *schema.ResourceData, key string) int {
	if attr, ok := d.GetOk(key); ok {
		return attr.(int)
	}
	return 0
}
