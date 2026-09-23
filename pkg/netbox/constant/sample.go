package constant

const FixtureMacList = `{
  "count": 1,
  "results": [
    {
      "id": 2602,
      "url": "http://example/api/dcim/mac-addresses/2602/",
      "display": "02:00:00:00:00:01",
      "mac_address": "02:00:00:00:00:01"
    }
  ]
}`

const FixtureMacRecord = `{
  "id": 2602,
  "url": "http://example/api/dcim/mac-addresses/2602/",
  "display": "02:00:00:00:00:01",
  "mac_address": "02:00:00:00:00:01",
  "assigned_object_type": "virtualization.vminterface",
  "assigned_object_id": 1
}`

const FixtureVirtualInterfaceList = `{
  "count": 1,
  "results": [
    {
      "id": 1,
      "url": "http://example/api/virtualization/interfaces/1/",
      "display": "ens3",
      "name": "ens3",
      "virtual_machine": {
        "id": 2,
        "url": "http://example/api/virtualization/virtual-machines/2/",
        "display": "alfa",
        "name": "alfa"
      },
      "count_ipaddresses": 0,
      "count_fhrp_groups": 0
    }
  ]
}`

const FixtureVirtualInterface = `{
  "id": 1,
  "url": "http://example/api/virtualization/interfaces/1/",
  "display": "ens3",
  "name": "ens3",
  "virtual_machine": {
    "id": 2,
    "url": "http://example/api/virtualization/virtual-machines/2/",
    "display": "alfa",
    "name": "alfa"
  },
  "count_ipaddresses": 0,
  "count_fhrp_groups": 0
}`
