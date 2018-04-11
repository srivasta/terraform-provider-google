package google

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceStackdriverAlertPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceStackdriverAlertPolicyCreate,
		Read:   resourceStackdriverAlertPolicyRead,
		Update: resourceStackdriverAlertPolicyUpdate,
		Delete: resourceStackdriverAlertPolicyDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceStackdriverAlertPolicyCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceStackdriverAlertPolicyRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceStackdriverAlertPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceStackdriverAlertPolicyDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
