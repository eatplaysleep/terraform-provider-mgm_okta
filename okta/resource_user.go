package okta

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"admin_roles": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "User Okta admin roles - ie. ['APP_ADMIN', 'USER_ADMIN']",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"city": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User city",
			},
			"cost_center": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User cost center",
			},
			"country_code": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User country code",
			},
			"department": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User department",
			},
			"display_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User display name, suitable to show end users",
			},
			"division": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User division",
			},
			"email": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "User primary email address. Default = user login",
				ValidateFunc: matchEmailRegexp,
			},
			"employee_number": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User employee number",
			},
			"first_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "User first name",
			},
			"honorific_prefix": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User honorific prefix",
			},
			"honorific_suffix": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User honorific suffix",
			},
			"language": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User preferred language",
			},
			"last_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "User last name",
			},
			"locale": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User default location",
			},
			"login": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				Description:  "User Okta login (must be an email address)",
				ForceNew:     true,
				ValidateFunc: matchEmailRegexp,
			},
			"manager": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Manager of User",
			},
			"manager_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Manager ID of User",
			},
			"middle_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User middle name",
			},
			"mobile_phone": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User mobile phone number",
			},
			"nickname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User nickname",
			},
			"organization": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User organization",
			},
			"postal_address": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User mailing address",
			},
			"primary_phone": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User primary phone number",
			},
			"profile_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User online profile (web page)",
			},
			"second_email": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User secondary email address, used for account recovery",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User state or region",
			},
			"street_address": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User street address",
			},
			"time_zone": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User default timezone",
			},
			"title": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User title",
			},
			"user_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User employee type",
			},
			"zip_code": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User zipcode or postal code",
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

//regex lovingly lifted from: http://www.golangprograms.com/regular-expression-to-validate-email-address.html
func matchEmailRegexp(val interface{}, key string) (warnings []string, errors []error) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if re.MatchString(val.(string)) == false {
		errors = append(errors, fmt.Errorf("%s field not a valid email address", key))
	}
	return warnings, errors
}
