package network

type NtnxNetwork struct {
	Metadata struct {
		GrandTotalEntities int `json:"grand_total_entities"`
		TotalEntities      int `json:"total_entities"`
	} `json:"metadata"`
	Entities []struct {
		LogicalTimestamp int `json:"logical_timestamp"`
		VlanID           int `json:"vlan_id"`
		IPConfig         struct {
			NetworkAddress string `json:"network_address"`
			PrefixLength   int    `json:"prefix_length"`
			DefaultGateway string `json:"default_gateway"`
			DhcpOptions    struct {
				DomainName        string `json:"domain_name"`
				DomainNameServers string `json:"domain_name_servers"`
				DomainSearch      string `json:"domain_search"`
			} `json:"dhcp_options"`
			Pool              []interface{} `json:"pool"`
			DhcpServerAddress string        `json:"dhcp_server_address"`
		} `json:"ip_config"`
		UUID string `json:"uuid"`
		Name string `json:"name"`
	} `json:"entities"`
}
