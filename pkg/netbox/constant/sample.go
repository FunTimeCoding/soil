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

const FixtureConfigurationContext = `{
  "id": 1,
  "url": "http://example/api/extras/config-contexts/1/",
  "display": "alfa",
  "name": "alfa",
  "data_path": "",
  "data_file": null,
  "weight": 1000,
  "is_active": true,
  "data": {"bravo": "charlie"}
}`

const FixtureConfigurationContextList = `{
  "count": 1,
  "results": [
    {
      "id": 1,
      "url": "http://example/api/extras/config-contexts/1/",
      "display": "alfa",
      "name": "alfa",
      "data_path": "",
      "data_file": null,
      "weight": 1000,
      "is_active": true,
      "data": {"bravo": "charlie"}
    }
  ]
}`

const FixtureChoiceSet = `{
  "id": 1,
  "name": "alfa",
  "description": "",
  "extra_choices": [["bravo", "bravo"], ["charlie", "charlie"]]
}`

const FixtureChoiceSetList = `{
  "count": 1,
  "results": [
    {
      "id": 1,
      "name": "alfa",
      "description": "",
      "extra_choices": [["bravo", "bravo"], ["charlie", "charlie"]]
    }
  ]
}`

const FixtureCustomField = `{
  "id": 2,
  "name": "delta",
  "label": "",
  "type": {"value": "multiselect", "label": "Multiple selection"},
  "object_types": ["dcim.device", "virtualization.virtualmachine"],
  "choice_set": {"id": 1, "name": "alfa"}
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

const FixtureDevice = `{
  "id": 3,
  "url": "http://example/api/dcim/devices/3/",
  "display": "bravo",
  "name": "bravo",
  "device_type": {
    "id": 1,
    "url": "http://example/api/dcim/device-types/1/",
    "display": "charlie",
    "manufacturer": {
      "id": 1,
      "url": "http://example/api/dcim/manufacturers/1/",
      "display": "delta",
      "name": "delta",
      "slug": "delta"
    },
    "model": "charlie",
    "slug": "charlie"
  },
  "role": {
    "id": 1,
    "url": "http://example/api/dcim/device-roles/1/",
    "display": "echo",
    "name": "echo",
    "slug": "echo",
    "_depth": 0
  },
  "site": {
    "id": 1,
    "url": "http://example/api/dcim/sites/1/",
    "display": "foxtrot",
    "name": "foxtrot",
    "slug": "foxtrot"
  },
  "local_context_data": {"golf": "hotel"},
  "console_port_count": 0,
  "console_server_port_count": 0,
  "power_port_count": 0,
  "power_outlet_count": 0,
  "front_port_count": 0,
  "rear_port_count": 0,
  "device_bay_count": 0,
  "module_bay_count": 0,
  "inventory_item_count": 0
}`

const FixtureDeviceList = `{
  "count": 1,
  "results": [
    {
      "id": 3,
      "url": "http://example/api/dcim/devices/3/",
      "display": "bravo",
      "name": "bravo",
      "device_type": {
        "id": 1,
        "url": "http://example/api/dcim/device-types/1/",
        "display": "charlie",
        "manufacturer": {
          "id": 1,
          "url": "http://example/api/dcim/manufacturers/1/",
          "display": "delta",
          "name": "delta",
          "slug": "delta"
        },
        "model": "charlie",
        "slug": "charlie"
      },
      "role": {
        "id": 1,
        "url": "http://example/api/dcim/device-roles/1/",
        "display": "echo",
        "name": "echo",
        "slug": "echo",
        "_depth": 0
      },
      "site": {
        "id": 1,
        "url": "http://example/api/dcim/sites/1/",
        "display": "foxtrot",
        "name": "foxtrot",
        "slug": "foxtrot"
      },
      "local_context_data": {"golf": "hotel"},
      "console_port_count": 0,
      "console_server_port_count": 0,
      "power_port_count": 0,
      "power_outlet_count": 0,
      "front_port_count": 0,
      "rear_port_count": 0,
      "device_bay_count": 0,
      "module_bay_count": 0,
      "inventory_item_count": 0
    }
  ]
}`

const FixtureVirtualMachine = `{
  "id": 2,
  "url": "http://example/api/virtualization/virtual-machines/2/",
  "display": "alfa",
  "name": "alfa",
  "local_context_data": {"golf": "hotel"},
  "virtual_disk_count": 0
}`

const FixturePowerOutlet = `{
  "id": 5,
  "url": "http://example/api/dcim/power-outlets/5/",
  "display": "india",
  "device": {
    "id": 3,
    "url": "http://example/api/dcim/devices/3/",
    "display": "bravo",
    "name": "bravo"
  },
  "name": "india",
  "link_peers": [],
  "connected_endpoints_reachable": false,
  "_occupied": false
}`

const FixturePowerOutletList = `{
  "count": 1,
  "results": [
    {
      "id": 5,
      "url": "http://example/api/dcim/power-outlets/5/",
      "display": "india",
      "device": {
        "id": 3,
        "url": "http://example/api/dcim/devices/3/",
        "display": "bravo",
        "name": "bravo"
      },
      "name": "india",
      "link_peers": [
        {
          "id": 6,
          "url": "http://example/api/dcim/power-ports/6/",
          "display": "juliett",
          "device": {
            "id": 4,
            "url": "http://example/api/dcim/devices/4/",
            "display": "kilo",
            "name": "kilo"
          },
          "name": "juliett"
        }
      ],
      "link_peers_type": "dcim.powerport",
      "connected_endpoints_reachable": true,
      "_occupied": true
    }
  ]
}`

const FixturePowerPort = `{
  "id": 6,
  "url": "http://example/api/dcim/power-ports/6/",
  "display": "juliett",
  "device": {
    "id": 4,
    "url": "http://example/api/dcim/devices/4/",
    "display": "kilo",
    "name": "kilo"
  },
  "name": "juliett",
  "link_peers": [],
  "connected_endpoints_reachable": false,
  "_occupied": false
}`

const FixtureCable = `{
  "id": 7,
  "url": "http://example/api/dcim/cables/7/",
  "display": "#7"
}`

const FixtureVirtualMachineList = `{
  "count": 1,
  "results": [
    {
      "id": 2,
      "url": "http://example/api/virtualization/virtual-machines/2/",
      "display": "alfa",
      "name": "alfa",
      "local_context_data": {"golf": "hotel"},
      "virtual_disk_count": 0
    }
  ]
}`
