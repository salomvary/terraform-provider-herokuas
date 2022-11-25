package herokuas

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTrigger() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTriggerCreate,
		ReadContext:   resourceTriggerRead,
		UpdateContext: resourceTriggerUpdate,
		DeleteContext: resourceTriggerDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dyno": {
				Type:     schema.TypeString,
				Required: true,
			},
			"frequency_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"schedule": {
				Type:     schema.TypeString,
				Required: true,
			},
			"timezone": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceTriggerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	newTrigger := Trigger{
		Name:          d.Get("name").(string),
		State:         d.Get("state").(string),
		Dyno:          d.Get("dyno").(string),
		FrequencyType: d.Get("frequency_type").(string),
		Schedule:      d.Get("schedule").(string),
		Timezone:      d.Get("timezone").(string),
		Value:         d.Get("value").(string),
		Timeout:       d.Get("timeout").(int),
	}

	trigger, err := apiClient.NewTrigger(&newTrigger)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(trigger.UUID)
	d.Set("timeout", trigger.Timeout)

	return diags
}

func resourceTriggerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	triggerUUID := d.Id()

	trigger, err := apiClient.GetTrigger(triggerUUID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(trigger.UUID)
	d.Set("name", trigger.Name)
	d.Set("state", trigger.State)
	d.Set("dyno", trigger.Dyno)
	d.Set("frequency_type", trigger.FrequencyType)
	d.Set("schedule", trigger.Schedule)
	d.Set("timezone", trigger.Timezone)
	d.Set("value", trigger.Value)
	d.Set("timeout", trigger.Timeout)

	return diags
}

func resourceTriggerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	trigger := Trigger{
		UUID:          d.Id(),
		Name:          d.Get("name").(string),
		State:         d.Get("state").(string),
		Dyno:          d.Get("dyno").(string),
		FrequencyType: d.Get("frequency_type").(string),
		Schedule:      d.Get("schedule").(string),
		Timezone:      d.Get("timezone").(string),
		Value:         d.Get("value").(string),
		Timeout:       d.Get("timeout").(int),
	}

	err := apiClient.UpdateTrigger(&trigger)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("last_updated", time.Now().Format(time.RFC850))

	return diags
}

func resourceTriggerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	triggerUUID := d.Id()

	err := apiClient.DeleteTrigger(triggerUUID)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
