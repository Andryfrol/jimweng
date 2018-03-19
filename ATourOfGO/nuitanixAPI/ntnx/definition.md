<!-- all the infroamtion refer to nutanix bible -->
#### Storage Pool
- Key Role: Group of physical devices
> Description: A storage pool is a group of physical storage devices including PCIe SSD, SSD, and HDD devices for the cluster.  The storage pool can span multiple Nutanix nodes and is expanded as the cluster scales.  In most configurations, only a single storage pool is leveraged.

#### Container
- Key Role: Group of VMs/files
> Description: A container is a logical segmentation of the Storage Pool and contains a group of VM or files (vDisks).  Some configuration options (e.g., RF) are configured at the container level, however are applied at the individual VM/file level.  Containers typically have a 1 to 1 mapping with a datastore (in the case of NFS/SMB).

#### vDisk
- Key Role: vDisk
> Description: A vDisk is any file over 512KB on DSF including .vmdks and VM hard disks.  vDisks are composed of extents which are grouped and stored on disk as an extent group.

## Block Services Constructs
The following entities compose Acropolis Block Services:

#### Data Services IP:
> Cluster wide IP address used for iSCSI login requests (Introduced in 4.7)
#### Volume Group:
> iSCSI target and group of disk devices allowing for centralized management, snapshotting, and policy application
#### Disk(s):
> Storage devices in the Volume Group (seen as LUNs for the iSCSI target)
#### Attachment:
> Allowing a specified initiator IQN access to the volume group
#### Secret(s):
> Secret used for CHAP/Mutual CHAP authentication
NOTE: On the backend, a VG’s disk is just a vDisk on DSF.

To use Block Services, the first thing we'll do is create a 'Volume Group' which is the iSCSI target. Once we've specified the details and added disk(s) we'll attach the Volume Group to a VM or Initiator IQN. This will allow the VM to access the iSCSI target (requests from an unknown initiator are rejected):