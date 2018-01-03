# govc about -u='agent.test@aildap.prophetstor.com:agent.test@172.31.17.100' -k=true

govc datacenter.info -u='agent.test@aildap.prophetstor.com:agent.test@172.31.17.100' -k=true
# Name:                Elsvent
#   Path:              /Elsvent
#   Hosts:             1
#   Clusters:          0
#   Virtual Machines:  7
#   Networks:          2
#   Datastores:        4
# Name:                DiskProphet
#   Path:              /DiskProphet
#   Hosts:             4
#   Clusters:          1
#   Virtual Machines:  16
#   Networks:          4
#   Datastores:        7

# need to specify a datacenter
govc datastore.info -u='agent.test@aildap.prophetstor.com:agent.test@172.31.17.100' -k=true -dc='DiskProphet'
govc datastore.info -u='agent.test@aildap.prophetstor.com:agent.test@172.31.17.100' -k=true -dc='Elsvent'

# Usage: govc datastore.info [OPTIONS] [PATH]...

# Options:
#   -cert=                    Certificate [GOVC_CERTIFICATE]
#   -dc=                      Datacenter [GOVC_DATACENTER]
#   -debug=false              Store debug logs [GOVC_DEBUG]
#   -dump=false               Enable Go output
#   -json=false               Enable JSON output
#   -k=false                  Skip verification of server certificate [GOVC_INSECURE]
#   -key=                     Private key [GOVC_PRIVATE_KEY]
#   -persist-session=true     Persist session to disk [GOVC_PERSIST_SESSION]
#   -tls-ca-certs=            TLS CA certificates file [GOVC_TLS_CA_CERTS]
#   -tls-known-hosts=         TLS known hosts file [GOVC_TLS_KNOWN_HOSTS]
#   -u=                       ESX or vCenter URL [GOVC_URL]
#   -vim-namespace=urn:vim25  Vim namespace [GOVC_VIM_NAMESPAC



govc ls -u='agent.test@aildap.prophetstor.com:agent.test@172.31.17.100' -k=true -l '*'

govc ls -u='agent.test@aildap.prophetstor.com:agent.test@172.31.17.100' -k=true -l DiskProphet/vm
# Usage: govc ls [OPTIONS] [PATH]...

# List inventory items.

# Examples:
#   govc ls -l '*'
#   govc ls -t ClusterComputeResource host
#   govc ls -t Datastore host/ClusterA/* | grep -v local | xargs -n1 basename | sort | uniq

# Options:
#   -L=false                  Follow managed object references
#   -cert=                    Certificate [GOVC_CERTIFICATE]
#   -dc=                      Datacenter [GOVC_DATACENTER]
#   -debug=false              Store debug logs [GOVC_DEBUG]
#   -dump=false               Enable Go output
#   -i=false                  Print the managed object reference
#   -json=false               Enable JSON output
#   -k=false                  Skip verification of server certificate [GOVC_INSECURE]
#   -key=                     Private key [GOVC_PRIVATE_KEY]
#   -l=false                  Long listing format
#   -persist-session=true     Persist session to disk [GOVC_PERSIST_SESSION]
#   -t=                       Object type
#   -tls-ca-certs=            TLS CA certificates file [GOVC_TLS_CA_CERTS]
#   -tls-known-hosts=         TLS known hosts file [GOVC_TLS_KNOWN_HOSTS]
#   -u=                       ESX or vCenter URL [GOVC_URL]
#   -vim-namespace=urn:vim25  Vim namespace [GOVC_VIM_NAMESPACE]
#   -vim-version=6.5          Vim version [GOVC_VIM_VERSION]

# this command lind would return all the vm under DiksProphet

govc vm.info -u='agent.test@aildap.prophetstor.com:agent.test@172.31.17.100' -k=true -dc='DiskProphet' VCSA_6.5_17_100

govc vm.info -u='agent.test@aildap.prophetstor.com:agent.test@172.31.17.100' -k=true -dc='DiskProphet' DSBB


# query disk under particular datacenter / datastore
govc host.storage.info -u='agent.test@aildap.prophetstor.com:agent.test@172.31.17.100' -k=true -dc='DiskProphet' -host='172.31.17.94'

# query snapshot vm under particular datacenter datastore
govc snapshot.tree -u='agent.test@aildap.prophetstor.com:agent.test@172.31.17.100' -k=true -dc='DiskProphet' -vm 'DiskProphet Agents'
govc snapshot.tree -u='agent.test@aildap.prophetstor.com:agent.test@172.31.17.100' -k=true -dc='DiskProphet' -vm 'hxexsa-fix'
govc snapshot.tree -u='agent.test@aildap.prophetstor.com:agent.test@172.31.17.100' -k=true -dc='DiskProphet' -vm 'JP1'




govc datastore.disk.info -u='agent.test@aildap.prophetstor.com:agent.test@172.31.17.100' -k=true -dc='DiskProphet' -ds='vsanB'


govc datastore.ls -u='matt.wu:password@172.31.17.100' -k=true -dc='DiskProphet' -ds='vsanB'

govc find -u='matt.wu:password@172.31.17.100' -k=true . -type s -summary.type vsan
