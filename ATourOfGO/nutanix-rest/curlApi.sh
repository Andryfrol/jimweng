# Nutanix Prism https://10.2.10.119:9440 ID: admin PW:RyfUA8xC3b}7@3[
# work for retriving vms and related information
curl --user USERACCOUNT:USERPASSWORD --insecure -H "Content-Type: application/json" -H "Accept: application/json" https://LOGINURL:9440/PrismGateway/services/rest/v1/vms/

# Retrieve data like VM and their configs, extract performance values, …
# Create VM, VLAN, vdisk, projects….
# Show typical tasks in the new self service portal (SSP)
# Using different kind of Nutanix REST APIs (v0.8, v1, v2 , v3)
# … a lot more

https://10.2.10.119:9440

admin

RyfUA8xC3b}7@3[


curl --user admin:RyfUA8xC3b}7@3[ --insecure -H "Content-Type: application/json" -H "Accept: application/json" https://10.2.10.119:9440/PrismGateway/services/rest/v1/vms/

curl --user admin:RyfUA8xC3b}7@3[ --insecure -H "Content-Type: application/json" -H "Accept: application/json" https://10.2.10.119:9440/PrismGateway/services/rest/v2.0/vms/

# still find way to work
curl --user admin:RyfUA8xC3b}7@3[ --insecure -H "Content-Type: application/json" -H "Accept: application/json" https://10.2.10.119:9440/api/nutanix/v3/batch

# list clusters with particular uuid
curl --user admin:RyfUA8xC3b}7@3[ --insecure -H "Content-Type: application/json" -H "Accept: application/json" https://10.2.10.119:9440/api/nutanix/v3/clusters/00054fd5-6de4-8422-74c5-782bcb637d0e


