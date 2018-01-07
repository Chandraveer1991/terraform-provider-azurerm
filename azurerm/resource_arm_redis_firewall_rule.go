package azurerm

import (
	"fmt"
	"log"

	"regexp"

	"github.com/Azure/azure-sdk-for-go/arm/redis"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceArmRedisFirewallRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmRedisFirewallRuleCreateUpdate,
		Read:   resourceArmRedisFirewallRuleRead,
		Update: resourceArmRedisFirewallRuleCreateUpdate,
		Delete: resourceArmRedisFirewallRuleDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRedisFirewallRuleName,
			},

			"redis_cache_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"resource_group_name": resourceGroupNameSchema(),

			"start_ip": {
				Type:     schema.TypeString,
				Required: true,
			},

			"end_ip": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceArmRedisFirewallRuleCreateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).redisFirewallClient
	log.Printf("[INFO] preparing arguments for AzureRM Redis Firewall Rule creation.")

	name := d.Get("name").(string)
	cacheName := d.Get("redis_cache_name").(string)
	resourceGroup := d.Get("resource_group_name").(string)
	startIP := d.Get("start_ip").(string)
	endIP := d.Get("end_ip").(string)

	parameters := redis.FirewallRule{
		Name: &name,
		FirewallRuleProperties: &redis.FirewallRuleProperties{
			StartIP: utils.String(startIP),
			EndIP:   utils.String(endIP),
		},
	}

	_, err := client.CreateOrUpdate(resourceGroup, cacheName, name, parameters)
	if err != nil {
		return err
	}

	read, err := client.Get(resourceGroup, cacheName, name)
	if err != nil {
		return err
	}
	if read.ID == nil {
		return fmt.Errorf("Cannot read Redis Firewall Rule %q (cache %q / resource group %q) ID", name, cacheName, resourceGroup)
	}

	d.SetId(*read.ID)

	return resourceArmRedisFirewallRuleRead(d, meta)
}

func resourceArmRedisFirewallRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).redisFirewallClient

	id, err := parseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resourceGroup := id.ResourceGroup
	cacheName := id.Path["Redis"]
	name := id.Path["firewallRules"]

	resp, err := client.Get(resourceGroup, cacheName, name)

	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[DEBUG] Redis Firewall Rule %q was not found in Cache %q / Resource Group %q - removing from state", name, cacheName, resourceGroup)
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error making Read request on Azure Redis Firewall Rule %q: %+v", name, err)
	}

	d.Set("name", name)
	d.Set("redis_cache_name", cacheName)
	d.Set("resource_group_name", resourceGroup)
	if props := resp.FirewallRuleProperties; props != nil {
		d.Set("start_ip", props.StartIP)
		d.Set("end_ip", props.EndIP)
	}

	return nil
}

func resourceArmRedisFirewallRuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).redisFirewallClient

	id, err := parseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resourceGroup := id.ResourceGroup
	cacheName := id.Path["Redis"]
	name := id.Path["firewallRules"]

	resp, err := client.Delete(resourceGroup, cacheName, name)

	if err != nil {
		if !utils.ResponseWasNotFound(resp) {
			return fmt.Errorf("Error issuing AzureRM delete request of Redis Firewall Rule %q (cache %q / resource group %q): %+v", name, cacheName, resourceGroup, err)
		}
	}

	return nil
}

func validateRedisFirewallRuleName(v interface{}, k string) (ws []string, es []error) {
	value := v.(string)

	if matched := regexp.MustCompile(`^[0-9a-zA-Z]+$`).Match([]byte(value)); !matched {
		es = append(es, fmt.Errorf("%q may only contain alphanumeric characters", k))
	}

	return
}
