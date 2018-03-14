package volumegroups

type NtnxVolumeGroups struct {
	Metadata struct {
		GrandTotalEntities int `json:"grand_total_entities"`
		TotalEntities      int `json:"total_entities"`
		Count              int `json:"count"`
	} `json:"metadata"`
	Entities []struct {
		UUID             string `json:"uuid"`
		Name             string `json:"name"`
		Description      string `json:"description"`
		LogicalTimestamp int    `json:"logical_timestamp"`
		DiskList         []struct {
			Index                int    `json:"index"`
			VmdiskUUID           string `json:"vmdisk_uuid"`
			ContainerID          int    `json:"container_id"`
			StorageContainerID   int    `json:"storage_container_id"`
			ContainerUUID        string `json:"container_uuid"`
			StorageContainerUUID string `json:"storage_container_uuid"`
			VmdiskSizeMb         int    `json:"vmdisk_size_mb"`
			FlashModeEnabled     bool   `json:"flash_mode_enabled"`
		} `json:"disk_list"`
		IscsiTarget    string `json:"iscsi_target"`
		AttachmentList []struct {
			VMUUID string `json:"vm_uuid"`
		} `json:"attachment_list"`
		FlashModeEnabled bool `json:"flash_mode_enabled"`
		IsShared         bool `json:"is_shared"`
	} `json:"entities"`
}
