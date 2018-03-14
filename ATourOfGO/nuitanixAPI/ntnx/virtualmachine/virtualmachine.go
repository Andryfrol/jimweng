package virtualmachine

type NtnxVirtualMachine struct {
	Metadata struct {
		GrandTotalEntities int `json:"grand_total_entities"`
		TotalEntities      int `json:"total_entities"`
		Count              int `json:"count"`
	} `json:"metadata"`
	Entities []struct {
		AllowLiveMigrate bool   `json:"allow_live_migrate"`
		GpusAssigned     bool   `json:"gpus_assigned"`
		Description      string `json:"description,omitempty"`
		HaPriority       int    `json:"ha_priority"`
		HostUUID         string `json:"host_uuid"`
		MemoryMb         int    `json:"memory_mb"`
		Name             string `json:"name"`
		NumCoresPerVcpu  int    `json:"num_cores_per_vcpu"`
		NumVcpus         int    `json:"num_vcpus"`
		PowerState       string `json:"power_state"`
		Timezone         string `json:"timezone"`
		UUID             string `json:"uuid"`
		VMFeatures       struct {
			AGENTVM bool `json:"AGENT_VM"`
		} `json:"vm_features"`
		VMLogicalTimestamp int `json:"vm_logical_timestamp"`
		Boot               struct {
			DiskAddress struct {
				DeviceBus   string `json:"device_bus"`
				DeviceIndex int    `json:"device_index"`
			} `json:"disk_address"`
			BootDeviceType string `json:"boot_device_type"`
		} `json:"boot,omitempty"`
	} `json:"entities"`
}
