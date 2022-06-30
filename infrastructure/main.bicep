param location string = resourceGroup().location
param shortloc string = 'eus'
param domain string = 'contoso'
param project string = 'pj1'
param env string = 'poc'

var vnetName = 'vnet-${domain}-${project}-${env}-${shortloc}'
var subnets = [
  {
    SubnetName: 'Default'
    SubnetAddressSpace: '10.50.0.0/24'
  }
  {
    SubnetName: 'AzureBastionSubnet'
    SubnetAddressSpace: '10.50.1.0/24'
  }
  {
    SubnetName: 'acaSubnet'
    SubnetAddressSpace: '10.50.2.0/24'
  }
  {
    SubnetName: 'vmSubnet'
    SubnetAddressSpace: '10.50.3.0/24'
  }
]

resource VNet 'Microsoft.Network/virtualNetworks@2021-08-01' = {
  location: location
  name: vnetName
  properties: {
    addressSpace: {
      addressPrefixes: [
        '10.50.0.0/16'
      ]
    }
    subnets: [for subnet in subnets: {
      name: subnet.SubnetName
      properties: {
        addressPrefix: subnet.SubnetAddressSpace
      }
    }]
  }
}
