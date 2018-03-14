package hosts

type NtnxHosts struct {
	Metadata struct {
		GrandTotalEntities int    `json:"grand_total_entities"`
		TotalEntities      int    `json:"total_entities"`
		FilterCriteria     string `json:"filter_criteria"`
		SortCriteria       string `json:"sort_criteria"`
		Page               int    `json:"page"`
		Count              int    `json:"count"`
		StartIndex         int    `json:"start_index"`
		EndIndex           int    `json:"end_index"`
	} `json:"metadata"`
	Entities []struct {
		ServiceVmid         string `json:"service_vmid"`
		UUID                string `json:"uuid"`
		DiskHardwareConfigs struct {
			Num1 struct {
				SerialNumber           string      `json:"serial_number"`
				DiskID                 string      `json:"disk_id"`
				DiskUUID               string      `json:"disk_uuid"`
				Location               int         `json:"location"`
				Bad                    bool        `json:"bad"`
				Mounted                bool        `json:"mounted"`
				MountPath              string      `json:"mount_path"`
				Model                  string      `json:"model"`
				Vendor                 string      `json:"vendor"`
				BootDisk               bool        `json:"boot_disk"`
				OnlyBootDisk           bool        `json:"only_boot_disk"`
				UnderDiagnosis         bool        `json:"under_diagnosis"`
				BackgroundOperation    interface{} `json:"background_operation"`
				CurrentFirmwareVersion string      `json:"current_firmware_version"`
				TargetFirmwareVersion  string      `json:"target_firmware_version"`
				CanAddAsNewDisk        bool        `json:"can_add_as_new_disk"`
				CanAddAsOldDisk        bool        `json:"can_add_as_old_disk"`
			} `json:"1"`
			Num2 struct {
				SerialNumber           string      `json:"serial_number"`
				DiskID                 string      `json:"disk_id"`
				DiskUUID               string      `json:"disk_uuid"`
				Location               int         `json:"location"`
				Bad                    bool        `json:"bad"`
				Mounted                bool        `json:"mounted"`
				MountPath              string      `json:"mount_path"`
				Model                  string      `json:"model"`
				Vendor                 string      `json:"vendor"`
				BootDisk               bool        `json:"boot_disk"`
				OnlyBootDisk           bool        `json:"only_boot_disk"`
				UnderDiagnosis         bool        `json:"under_diagnosis"`
				BackgroundOperation    interface{} `json:"background_operation"`
				CurrentFirmwareVersion string      `json:"current_firmware_version"`
				TargetFirmwareVersion  string      `json:"target_firmware_version"`
				CanAddAsNewDisk        bool        `json:"can_add_as_new_disk"`
				CanAddAsOldDisk        bool        `json:"can_add_as_old_disk"`
			} `json:"2"`
			Num3 struct {
				SerialNumber           string      `json:"serial_number"`
				DiskID                 interface{} `json:"disk_id"`
				DiskUUID               interface{} `json:"disk_uuid"`
				Location               int         `json:"location"`
				Bad                    bool        `json:"bad"`
				Mounted                bool        `json:"mounted"`
				MountPath              string      `json:"mount_path"`
				Model                  string      `json:"model"`
				Vendor                 string      `json:"vendor"`
				BootDisk               bool        `json:"boot_disk"`
				OnlyBootDisk           bool        `json:"only_boot_disk"`
				UnderDiagnosis         bool        `json:"under_diagnosis"`
				BackgroundOperation    interface{} `json:"background_operation"`
				CurrentFirmwareVersion string      `json:"current_firmware_version"`
				TargetFirmwareVersion  string      `json:"target_firmware_version"`
				CanAddAsNewDisk        bool        `json:"can_add_as_new_disk"`
				CanAddAsOldDisk        bool        `json:"can_add_as_old_disk"`
			} `json:"3"`
			Num4 struct {
				SerialNumber           string      `json:"serial_number"`
				DiskID                 string      `json:"disk_id"`
				DiskUUID               string      `json:"disk_uuid"`
				Location               int         `json:"location"`
				Bad                    bool        `json:"bad"`
				Mounted                bool        `json:"mounted"`
				MountPath              string      `json:"mount_path"`
				Model                  string      `json:"model"`
				Vendor                 string      `json:"vendor"`
				BootDisk               bool        `json:"boot_disk"`
				OnlyBootDisk           bool        `json:"only_boot_disk"`
				UnderDiagnosis         bool        `json:"under_diagnosis"`
				BackgroundOperation    interface{} `json:"background_operation"`
				CurrentFirmwareVersion string      `json:"current_firmware_version"`
				TargetFirmwareVersion  string      `json:"target_firmware_version"`
				CanAddAsNewDisk        bool        `json:"can_add_as_new_disk"`
				CanAddAsOldDisk        bool        `json:"can_add_as_old_disk"`
			} `json:"4"`
			Num5 struct {
				SerialNumber           string      `json:"serial_number"`
				DiskID                 string      `json:"disk_id"`
				DiskUUID               string      `json:"disk_uuid"`
				Location               int         `json:"location"`
				Bad                    bool        `json:"bad"`
				Mounted                bool        `json:"mounted"`
				MountPath              string      `json:"mount_path"`
				Model                  string      `json:"model"`
				Vendor                 string      `json:"vendor"`
				BootDisk               bool        `json:"boot_disk"`
				OnlyBootDisk           bool        `json:"only_boot_disk"`
				UnderDiagnosis         bool        `json:"under_diagnosis"`
				BackgroundOperation    interface{} `json:"background_operation"`
				CurrentFirmwareVersion string      `json:"current_firmware_version"`
				TargetFirmwareVersion  string      `json:"target_firmware_version"`
				CanAddAsNewDisk        bool        `json:"can_add_as_new_disk"`
				CanAddAsOldDisk        bool        `json:"can_add_as_old_disk"`
			} `json:"5"`
			Num6 struct {
				SerialNumber           string      `json:"serial_number"`
				DiskID                 string      `json:"disk_id"`
				DiskUUID               string      `json:"disk_uuid"`
				Location               int         `json:"location"`
				Bad                    bool        `json:"bad"`
				Mounted                bool        `json:"mounted"`
				MountPath              string      `json:"mount_path"`
				Model                  string      `json:"model"`
				Vendor                 string      `json:"vendor"`
				BootDisk               bool        `json:"boot_disk"`
				OnlyBootDisk           bool        `json:"only_boot_disk"`
				UnderDiagnosis         bool        `json:"under_diagnosis"`
				BackgroundOperation    interface{} `json:"background_operation"`
				CurrentFirmwareVersion string      `json:"current_firmware_version"`
				TargetFirmwareVersion  string      `json:"target_firmware_version"`
				CanAddAsNewDisk        bool        `json:"can_add_as_new_disk"`
				CanAddAsOldDisk        bool        `json:"can_add_as_old_disk"`
			} `json:"6"`
			Num7 interface{} `json:"7"`
		} `json:"disk_hardware_configs"`
		Name                 string      `json:"name"`
		ServiceVmexternalIP  string      `json:"service_vmexternal_ip"`
		OplogDiskPct         float64     `json:"oplog_disk_pct"`
		OplogDiskSize        int64       `json:"oplog_disk_size"`
		HypervisorKey        string      `json:"hypervisor_key"`
		HypervisorAddress    string      `json:"hypervisor_address"`
		HypervisorUsername   string      `json:"hypervisor_username"`
		HypervisorPassword   interface{} `json:"hypervisor_password"`
		ManagementServerName string      `json:"management_server_name"`
		IpmiAddress          string      `json:"ipmi_address"`
		IpmiUsername         string      `json:"ipmi_username"`
		IpmiPassword         interface{} `json:"ipmi_password"`
		Monitored            bool        `json:"monitored"`
		Position             struct {
			Ordinal          int         `json:"ordinal"`
			Name             string      `json:"name"`
			PhysicalPosition interface{} `json:"physical_position"`
		} `json:"position"`
		Serial                         string      `json:"serial"`
		BlockSerial                    string      `json:"block_serial"`
		BlockModel                     string      `json:"block_model"`
		BlockModelName                 string      `json:"block_model_name"`
		BlockLocation                  interface{} `json:"block_location"`
		HostMaintenanceModeReason      interface{} `json:"host_maintenance_mode_reason"`
		HypervisorState                string      `json:"hypervisor_state"`
		MetadataStoreStatus            string      `json:"metadata_store_status"`
		MetadataStoreStatusMessage     string      `json:"metadata_store_status_message"`
		State                          string      `json:"state"`
		DynamicRingChangingNode        interface{} `json:"dynamic_ring_changing_node"`
		RemovalStatus                  []string    `json:"removal_status"`
		VzoneName                      string      `json:"vzone_name"`
		CPUModel                       string      `json:"cpu_model"`
		NumCPUCores                    int         `json:"num_cpu_cores"`
		NumCPUThreads                  int         `json:"num_cpu_threads"`
		NumCPUSockets                  int         `json:"num_cpu_sockets"`
		CPUFrequencyInHz               int64       `json:"cpu_frequency_in_hz"`
		CPUCapacityInHz                int64       `json:"cpu_capacity_in_hz"`
		MemoryCapacityInBytes          int64       `json:"memory_capacity_in_bytes"`
		HypervisorFullName             string      `json:"hypervisor_full_name"`
		HypervisorType                 string      `json:"hypervisor_type"`
		NumVms                         int         `json:"num_vms"`
		BootTimeInUsecs                int64       `json:"boot_time_in_usecs"`
		IsDegraded                     bool        `json:"is_degraded"`
		FailoverClusterFqdn            interface{} `json:"failover_cluster_fqdn"`
		FailoverClusterNodeState       interface{} `json:"failover_cluster_node_state"`
		RebootPending                  bool        `json:"reboot_pending"`
		DefaultVMLocation              interface{} `json:"default_vm_location"`
		DefaultVMStorageContainerID    interface{} `json:"default_vm_storage_container_id"`
		DefaultVMStorageContainerUUID  interface{} `json:"default_vm_storage_container_uuid"`
		DefaultVhdLocation             interface{} `json:"default_vhd_location"`
		DefaultVhdStorageContainerID   interface{} `json:"default_vhd_storage_container_id"`
		DefaultVhdStorageContainerUUID interface{} `json:"default_vhd_storage_container_uuid"`
		BiosVersion                    interface{} `json:"bios_version"`
		BiosModel                      interface{} `json:"bios_model"`
		BmcVersion                     interface{} `json:"bmc_version"`
		BmcModel                       interface{} `json:"bmc_model"`
		HbaFirmwaresList               interface{} `json:"hba_firmwares_list"`
		ClusterUUID                    string      `json:"cluster_uuid"`
		Stats                          struct {
			HypervisorAvgIoLatencyUsecs          string `json:"hypervisor_avg_io_latency_usecs"`
			NumReadIops                          string `json:"num_read_iops"`
			HypervisorWriteIoBandwidthKBps       string `json:"hypervisor_write_io_bandwidth_kBps"`
			TimespanUsecs                        string `json:"timespan_usecs"`
			ControllerNumReadIops                string `json:"controller_num_read_iops"`
			ReadIoPpm                            string `json:"read_io_ppm"`
			ControllerNumIops                    string `json:"controller_num_iops"`
			TotalReadIoTimeUsecs                 string `json:"total_read_io_time_usecs"`
			ControllerTotalReadIoTimeUsecs       string `json:"controller_total_read_io_time_usecs"`
			HypervisorNumIo                      string `json:"hypervisor_num_io"`
			ControllerTotalTransformedUsageBytes string `json:"controller_total_transformed_usage_bytes"`
			HypervisorCPUUsagePpm                string `json:"hypervisor_cpu_usage_ppm"`
			ControllerNumWriteIo                 string `json:"controller_num_write_io"`
			AvgReadIoLatencyUsecs                string `json:"avg_read_io_latency_usecs"`
			ContentCacheLogicalSsdUsageBytes     string `json:"content_cache_logical_ssd_usage_bytes"`
			ControllerTotalIoTimeUsecs           string `json:"controller_total_io_time_usecs"`
			ControllerTotalReadIoSizeKbytes      string `json:"controller_total_read_io_size_kbytes"`
			ControllerNumSeqIo                   string `json:"controller_num_seq_io"`
			ControllerReadIoPpm                  string `json:"controller_read_io_ppm"`
			ContentCacheNumLookups               string `json:"content_cache_num_lookups"`
			ControllerTotalIoSizeKbytes          string `json:"controller_total_io_size_kbytes"`
			ContentCacheHitPpm                   string `json:"content_cache_hit_ppm"`
			ControllerNumIo                      string `json:"controller_num_io"`
			HypervisorAvgReadIoLatencyUsecs      string `json:"hypervisor_avg_read_io_latency_usecs"`
			ContentCacheNumDedupRefCountPph      string `json:"content_cache_num_dedup_ref_count_pph"`
			NumWriteIops                         string `json:"num_write_iops"`
			ControllerNumRandomIo                string `json:"controller_num_random_io"`
			NumIops                              string `json:"num_iops"`
			HypervisorNumReadIo                  string `json:"hypervisor_num_read_io"`
			HypervisorTotalReadIoTimeUsecs       string `json:"hypervisor_total_read_io_time_usecs"`
			ControllerAvgIoLatencyUsecs          string `json:"controller_avg_io_latency_usecs"`
			NumIo                                string `json:"num_io"`
			ControllerNumReadIo                  string `json:"controller_num_read_io"`
			HypervisorNumWriteIo                 string `json:"hypervisor_num_write_io"`
			ControllerSeqIoPpm                   string `json:"controller_seq_io_ppm"`
			ControllerReadIoBandwidthKBps        string `json:"controller_read_io_bandwidth_kBps"`
			ControllerIoBandwidthKBps            string `json:"controller_io_bandwidth_kBps"`
			HypervisorNumReceivedBytes           string `json:"hypervisor_num_received_bytes"`
			HypervisorTimespanUsecs              string `json:"hypervisor_timespan_usecs"`
			HypervisorNumWriteIops               string `json:"hypervisor_num_write_iops"`
			TotalReadIoSizeKbytes                string `json:"total_read_io_size_kbytes"`
			HypervisorTotalIoSizeKbytes          string `json:"hypervisor_total_io_size_kbytes"`
			AvgIoLatencyUsecs                    string `json:"avg_io_latency_usecs"`
			HypervisorNumReadIops                string `json:"hypervisor_num_read_iops"`
			ContentCacheSavedSsdUsageBytes       string `json:"content_cache_saved_ssd_usage_bytes"`
			ControllerWriteIoBandwidthKBps       string `json:"controller_write_io_bandwidth_kBps"`
			ControllerWriteIoPpm                 string `json:"controller_write_io_ppm"`
			HypervisorAvgWriteIoLatencyUsecs     string `json:"hypervisor_avg_write_io_latency_usecs"`
			HypervisorNumTransmittedBytes        string `json:"hypervisor_num_transmitted_bytes"`
			HypervisorTotalReadIoSizeKbytes      string `json:"hypervisor_total_read_io_size_kbytes"`
			ReadIoBandwidthKBps                  string `json:"read_io_bandwidth_kBps"`
			HypervisorMemoryUsagePpm             string `json:"hypervisor_memory_usage_ppm"`
			HypervisorNumIops                    string `json:"hypervisor_num_iops"`
			HypervisorIoBandwidthKBps            string `json:"hypervisor_io_bandwidth_kBps"`
			ControllerNumWriteIops               string `json:"controller_num_write_iops"`
			TotalIoTimeUsecs                     string `json:"total_io_time_usecs"`
			ContentCachePhysicalSsdUsageBytes    string `json:"content_cache_physical_ssd_usage_bytes"`
			ControllerRandomIoPpm                string `json:"controller_random_io_ppm"`
			ControllerAvgReadIoSizeKbytes        string `json:"controller_avg_read_io_size_kbytes"`
			TotalTransformedUsageBytes           string `json:"total_transformed_usage_bytes"`
			AvgWriteIoLatencyUsecs               string `json:"avg_write_io_latency_usecs"`
			NumReadIo                            string `json:"num_read_io"`
			WriteIoBandwidthKBps                 string `json:"write_io_bandwidth_kBps"`
			HypervisorReadIoBandwidthKBps        string `json:"hypervisor_read_io_bandwidth_kBps"`
			RandomIoPpm                          string `json:"random_io_ppm"`
			TotalUntransformedUsageBytes         string `json:"total_untransformed_usage_bytes"`
			HypervisorTotalIoTimeUsecs           string `json:"hypervisor_total_io_time_usecs"`
			NumRandomIo                          string `json:"num_random_io"`
			ControllerAvgWriteIoSizeKbytes       string `json:"controller_avg_write_io_size_kbytes"`
			ControllerAvgReadIoLatencyUsecs      string `json:"controller_avg_read_io_latency_usecs"`
			NumWriteIo                           string `json:"num_write_io"`
			TotalIoSizeKbytes                    string `json:"total_io_size_kbytes"`
			IoBandwidthKBps                      string `json:"io_bandwidth_kBps"`
			ContentCachePhysicalMemoryUsageBytes string `json:"content_cache_physical_memory_usage_bytes"`
			ControllerTimespanUsecs              string `json:"controller_timespan_usecs"`
			NumSeqIo                             string `json:"num_seq_io"`
			ContentCacheSavedMemoryUsageBytes    string `json:"content_cache_saved_memory_usage_bytes"`
			SeqIoPpm                             string `json:"seq_io_ppm"`
			WriteIoPpm                           string `json:"write_io_ppm"`
			ControllerAvgWriteIoLatencyUsecs     string `json:"controller_avg_write_io_latency_usecs"`
			ContentCacheLogicalMemoryUsageBytes  string `json:"content_cache_logical_memory_usage_bytes"`
		} `json:"stats"`
		UsageStats struct {
			StorageTierDasSataUsageBytes    string `json:"storage_tier.das-sata.usage_bytes"`
			StorageCapacityBytes            string `json:"storage.capacity_bytes"`
			StorageLogicalUsageBytes        string `json:"storage.logical_usage_bytes"`
			StorageTierDasSataCapacityBytes string `json:"storage_tier.das-sata.capacity_bytes"`
			StorageFreeBytes                string `json:"storage.free_bytes"`
			StorageTierSsdUsageBytes        string `json:"storage_tier.ssd.usage_bytes"`
			StorageTierSsdCapacityBytes     string `json:"storage_tier.ssd.capacity_bytes"`
			StorageTierDasSataFreeBytes     string `json:"storage_tier.das-sata.free_bytes"`
			StorageUsageBytes               string `json:"storage.usage_bytes"`
			StorageTierSsdFreeBytes         string `json:"storage_tier.ssd.free_bytes"`
		} `json:"usage_stats"`
		HasCsr                                 bool          `json:"has_csr"`
		HostNicIds                             []interface{} `json:"host_nic_ids"`
		HostGpus                               interface{}   `json:"host_gpus"`
		KeyManagementDeviceToCertificateStatus struct {
		} `json:"key_management_device_to_certificate_status"`
		HostInMaintenanceMode interface{} `json:"host_in_maintenance_mode"`
	} `json:"entities"`
}
