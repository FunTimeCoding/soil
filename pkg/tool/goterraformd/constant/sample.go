package constant

const (
	FixtureLockDetail = `{"ID":"11111111-2222-4333-8444-555555555555","Operation":"OperationTypePlan","Info":"","Who":"user@host.example","Version":"1.0.0","Created":"2020-01-02T03:04:05.123456789Z","Path":""}`

	FixtureLockTemplate = `{"ID":%q,"Operation":"OperationTypeApply","Info":"","Who":%q,"Version":"1.16.5","Created":%q,"Path":""}`

	FixtureRecorder = `#!/bin/sh
echo "$*" >> %q
[ "$1" = force-unlock ] && exit %d
[ "$1" = show ] && cat %q
exit 0
`

	FixtureEmptyPlan = `{}`

	FixturePlan = `{
  "resource_changes": [
    {"address": "module.storage.example_disk.charlie", "type": "example_disk",
     "change": {"actions": ["create"], "before": null, "after": {"size": 8}}},
    {"address": "example_network.delta", "type": "example_network",
     "change": {"actions": ["update"], "before": {}, "after": {}}},
    {"address": "module.compute.proxmox_download_file.image", "type": "proxmox_download_file",
     "change": {"actions": ["update"], "before": {}, "after": {}}},
    {"address": "example_role.echo", "type": "example_role",
     "change": {"actions": ["create"], "before": null, "after": {"name": "echo"}}},
    {"address": "module.compute.proxmox_virtual_environment_vm.golf", "type": "proxmox_virtual_environment_vm",
     "change": {"actions": ["delete", "create"], "before": {"name": "golf"}, "after": {"name": "golf"}}},
    {"address": "module.compute.example_disk.juliett", "type": "example_disk",
     "change": {"actions": ["update"], "before": {}, "after": {}}},
    {"address": "module.compute.example_disk.kilo", "type": "example_disk",
     "change": {"actions": ["update"], "before": {}, "after": {}}},
    {"address": "module.compute.proxmox_virtual_environment_vm.foxtrot", "type": "proxmox_virtual_environment_vm",
     "change": {"actions": ["delete"], "before": {"name": "foxtrot"}, "after": null}},
    {"address": "data.example_lookup.india", "type": "example_lookup",
     "change": {"actions": ["read"], "before": null, "after": {}}},
    {"address": "module.compute.proxmox_virtual_environment_vm.mike", "type": "proxmox_virtual_environment_vm",
     "change": {"actions": ["no-op"], "before": {"name": "mike"}, "after": {"name": "mike"}}}
  ],
  "planned_values": {"root_module": {
    "resources": [
      {"address": "example_network.delta", "type": "example_network", "values": {}},
      {"address": "example_role.echo", "type": "example_role", "values": {"name": "echo"}}
    ],
    "child_modules": [
      {"resources": [
        {"address": "module.compute.proxmox_virtual_environment_vm.alfa", "type": "proxmox_virtual_environment_vm", "values": {"name": "alfa"}},
        {"address": "module.compute.proxmox_virtual_environment_vm.bravo", "type": "proxmox_virtual_environment_vm", "values": {"name": "bravo"}},
        {"address": "module.compute.proxmox_virtual_environment_vm.golf", "type": "proxmox_virtual_environment_vm", "values": {"name": "golf"}},
        {"address": "module.compute.proxmox_virtual_environment_vm.hotel", "type": "proxmox_virtual_environment_vm", "values": {"name": "hotel"}},
        {"address": "module.compute.proxmox_virtual_environment_vm.lima", "type": "proxmox_virtual_environment_vm", "values": {"name": "lima"}},
        {"address": "module.compute.proxmox_virtual_environment_vm.mike", "type": "proxmox_virtual_environment_vm", "values": {"name": "mike"}}
      ]},
      {"resources": [
        {"address": "module.storage.example_disk.charlie", "type": "example_disk", "values": {"size": 8}}
      ]}
    ]
  }},
  "prior_state": {"values": {"root_module": {"child_modules": [
    {"resources": [
      {"address": "module.compute.proxmox_virtual_environment_vm.foxtrot", "type": "proxmox_virtual_environment_vm", "values": {"name": "foxtrot"}}
    ]}
  ]}}},
  "configuration": {"root_module": {
    "resources": [
      {"address": "example_network.delta", "type": "example_network",
       "expressions": {"members": {"references": ["module.compute.hotel_id", "module.compute", "module.compute.lima_id"]}}},
      {"address": "example_role.echo", "type": "example_role", "expressions": {"name": {"constant_value": "echo"}}},
      {"address": "data.example_lookup.india", "type": "example_lookup", "expressions": {}}
    ],
    "module_calls": {
      "compute": {"expressions": {}, "module": {
        "resources": [
          {"address": "proxmox_download_file.image", "type": "proxmox_download_file", "expressions": {}},
          {"address": "proxmox_virtual_environment_vm.alfa", "type": "proxmox_virtual_environment_vm",
           "expressions": {"name": {"constant_value": "alfa"}, "disk": [{"file_id": {"references": ["proxmox_download_file.image.id", "proxmox_download_file.image"]}}]}},
          {"address": "proxmox_virtual_environment_vm.bravo", "type": "proxmox_virtual_environment_vm", "expressions": {}},
          {"address": "proxmox_virtual_environment_vm.golf", "type": "proxmox_virtual_environment_vm", "expressions": {}},
          {"address": "proxmox_virtual_environment_vm.hotel", "type": "proxmox_virtual_environment_vm", "expressions": {}},
          {"address": "proxmox_virtual_environment_vm.lima", "type": "proxmox_virtual_environment_vm", "expressions": {}},
          {"address": "proxmox_virtual_environment_vm.mike", "type": "proxmox_virtual_environment_vm", "expressions": {}},
          {"address": "example_disk.juliett", "type": "example_disk",
           "expressions": {"vm_id": {"references": ["proxmox_virtual_environment_vm.golf.id", "proxmox_virtual_environment_vm.golf"]}}},
          {"address": "example_disk.kilo", "type": "example_disk",
           "expressions": {"vm_id": {"references": ["proxmox_virtual_environment_vm.bravo.id", "proxmox_virtual_environment_vm.bravo"]}}}
        ],
        "outputs": {
          "alfa_id": {"expression": {"references": ["proxmox_virtual_environment_vm.alfa.id", "proxmox_virtual_environment_vm.alfa"]}},
          "hotel_id": {"expression": {"references": ["proxmox_virtual_environment_vm.hotel.id", "proxmox_virtual_environment_vm.hotel"]}},
          "lima_id": {"expression": {"references": ["proxmox_virtual_environment_vm.lima.id", "proxmox_virtual_environment_vm.lima"]}}
        }
      }},
      "storage": {"expressions": {"vm_id": {"references": ["module.compute.alfa_id", "module.compute"]}}, "module": {
        "resources": [
          {"address": "example_disk.charlie", "type": "example_disk",
           "expressions": {"vm_id": {"references": ["var.vm_id"]}}}
        ]
      }}
    }
  }}
}`
)
