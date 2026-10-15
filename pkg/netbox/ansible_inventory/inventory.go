package ansible_inventory

type Inventory struct {
	groups map[string]map[string]string
}
