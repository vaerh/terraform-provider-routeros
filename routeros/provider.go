package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	ErrorMsgPut    = "An error was encountered while sending a PUT request to the API"
	ErrorMsgGet    = "An error was encountered while sending a GET request to the API"
	ErrorMsgPatch  = "An error was encountered while sending a PATCH request to the API"
	ErrorMsgDelete = "An error was encountered while sending a DELETE request to the API"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"hosturl": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"ROS_HOSTURL", "MIKROTIK_HOST"}, nil),
				Description: `URL of the MikroTik router, default is TLS connection to REST.    
	* API: api[s]://host[:port]
		* api://router.local
		* apis://router.local:8729
	* REST: https://host
		* https://router.local
		* router.local
		* 127.0.0.1  


	export ROS_HOSTURL=router.local or export MIKROTIK_HOST=router.local
`,
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"ROS_USERNAME", "MIKROTIK_USER"}, nil),
				Description: `Username for the MikroTik WEB/Winbox.


	export ROS_USERNAME=admin or export MIKROTIK_USER=admin
`,
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"ROS_PASSWORD", "MIKROTIK_PASSWORD"}, nil),
				Description: "Password for the MikroTik user.",
				Sensitive:   true,
			},
			"ca_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"ROS_CA_CERTIFICATE", "MIKROTIK_CA_CERTIFICATE"}, nil),
				Description: "Path to MikroTik's certificate authority file.",
			},
			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"ROS_INSECURE", "MIKROTIK_INSECURE"}, false),
				Description: "Whether to verify the SSL certificate or not.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"routeros_dhcp_client":          ResourceDhcpClient(),
			"routeros_dhcp_server":          ResourceDhcpServer(),
			"routeros_dhcp_server_lease":    ResourceDhcpServerLease(),
			"routeros_dhcp_server_network":  ResourceDhcpServerNetwork(),
			"routeros_dns":                  ResourceDns(),
			"routeros_dns_record":           ResourceDnsRecord(),
			"routeros_firewall_addr_list":   ResourceIPFirewallAddrList(),
			"routeros_firewall_filter":      ResourceIPFirewallFilter(),
			"routeros_firewall_mangle":      ResourceIPFirewallMangle(),
			"routeros_firewall_nat":         ResourceIPFirewallNat(),
			"routeros_ip_address":           ResourceIPAddress(),
			"routeros_ip_pool":              ResourceIPPool(),
			"routeros_ip_route":             ResourceIPRoute(),
			"routeros_ipv6_address":         ResourceIPv6Address(),
			"routeros_ipv6_firewall_filter": ResourceIPv6FirewallFilter(),
			"routeros_bridge":               ResourceInterfaceBridge(),
			"routeros_bridge_port":          ResourceInterfaceBridgePort(),
			"routeros_bridge_vlan":          ResourceInterfaceBridgeVlan(),
			"routeros_gre":                  ResourceInterfaceGre(),
			"routeros_vlan":                 ResourceInterfaceVlan(),
			"routeros_vrrp":                 ResourceInterfaceVrrp(),
			"routeros_wireguard":            ResourceInterfaceWireguard(),
			"routeros_wireguard_peer":       ResourceInterfaceWireguardPeer(),
			"routeros_identity":             ResourceSystemIdentity(),
			"routeros_scheduler":            ResourceSystemScheduler(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"routeros_interfaces":   DatasourceInterfaces(),
			"routeros_ip_addresses": DatasourceIPAddresses(),
			"routeros_ip_routes":    DatasourceIPRoutes(),
			"routeros_firewall":     DatasourceFirewall(),
		},
		ConfigureContextFunc: NewClient,
	}
}

func NewProvider() *schema.Provider {
	return Provider()
}
