# Nutanix Prism https://10.2.10.119:9440 ID: admin PW:RyfUA8xC3b}7@3[
# work for retriving vms and related information
curl --user USERACCOUNT:USERPASSWORD --insecure -H "Content-Type: application/json" -H "Accept: application/json" https://LOGINURL:9440/PrismGateway/services/rest/v1/vms/

# Retrieve data like VM and their configs, extract performance values, …
# Create VM, VLAN, vdisk, projects….
# Show typical tasks in the new self service portal (SSP)
# Using different kind of Nutanix REST APIs (v0.8, v1, v2 , v3)
# … a lot more

curl --user admin:RyfUA8xC3b}7@3[ --insecure -H "Content-Type: application/json" -H "Accept: application/json" https://10.2.10.119:9440/PrismGateway/services/rest/v1/vms/

curl --user admin:RyfUA8xC3b}7@3[ --insecure -H "Content-Type: application/json" -H "Accept: application/json" https://10.2.10.119:9440/PrismGateway/services/rest/v2.0/vms/

# still find way to work
curl --user admin:RyfUA8xC3b}7@3[ --insecure -H "Content-Type: application/json" -H "Accept: application/json" https://10.2.10.119:9440/api/nutanix/v3/batch

# list clusters with particular uuid
curl --user admin:RyfUA8xC3b}7@3[ --insecure -H "Content-Type: application/json" -H "Accept: application/json" https://10.2.10.119:9440/api/nutanix/v3/clusters/00054fd5-6de4-8422-74c5-782bcb637d0e

# curl --user admin:RyfUA8xC3b}7@3[ --insecure -H "Content-Type: application/json" -H "Accept: application/json" -d '{ "kind" : "cluster" , "offset" : 0 , "length" : 0}' https://10.2.10.119:9440/api/nutanix/v3/clusters/list


curl --user admin:RyfUA8xC3b}7@3[ --insecure -H "Content-Type: application/json" -H "Accept: application/json" https://10.2.10.119:9440/api/nutanix/v3/clusters/list -d '{ "kind" : "cluster" , "offset" : 0 , "length" : 0}' 


curl --user admin:RyfUA8xC3b}7@3[ --insecure -H "Content-Type: application/json" -H "Accept: application/json" https://10.2.10.119:9440/api/nutanix/v3/clusters/list -d '{ "kind" : "cluster" }' 


curl --user admin:RyfUA8xC3b}7@3[ --insecure -H "Content-Type: application/json" -H "Accept: application/json" https://10.2.10.119:9440/api/nutanix/v3/disks/list -d '{ "kind" : "disk" }' 


 
#  
# nutanix curl with v2 restApi
# 

# cluster
curl -X GET --header 'Accept: application/json' 'https://10.2.10.119:9440/api/nutanix/v2.0/cluster/' -u admin:RyfUA8xC3b}7@3[ -k

# disk
curl -X GET --header 'Accept: application/json' 'https://10.2.10.119:9440/api/nutanix/v2.0/disks/' -u admin:RyfUA8xC3b}7@3[ -k

# host
curl -X GET --header 'Accept: application/json' 'https://10.2.10.119:9440/api/nutanix/v2.0/hosts/' -u admin:RyfUA8xC3b}7@3[ -k

# storage_container
curl -X GET --header 'Accept: application/json' 'https://10.2.10.119:9440/api/nutanix/v2.0/storage_containers/' -u admin:RyfUA8xC3b}7@3[ -k

# virtual machine
curl -X GET --header 'Accept: application/json' 'https://10.2.10.119:9440/api/nutanix/v2.0/vms/' -u admin:RyfUA8xC3b}7@3[ -k

# volume groups
curl -X GET --header 'Accept: application/json' 'https://10.2.10.119:9440/api/nutanix/v2.0/volume_groups/' -u admin:RyfUA8xC3b}7@3[ -k

# vstores
curl -X GET --header 'Accept: application/json' 'https://10.2.10.119:9440/api/nutanix/v2.0/vstores/' -u admin:RyfUA8xC3b}7@3[ -k


# 
# nutanix curl with v1 restApi 
# 

 curl -u admin:RyfUA8xC3b}7@3[ -k -H "Content-Type: application/json" -H "Accept: application/json" https://10.2.10.119:9440/PrismGateway/services/rest/v1/hosts/

 curl --user admin:RyfUA8xC3b}7@3[ --insecure -H "Content-Type: application/json" -H "Accept: application/json" https://10.2.10.119:9440/PrismGateway/services/rest/v1/disks/
 

 RyfUA8xC3b}7@3[

