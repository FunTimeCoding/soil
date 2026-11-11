package charlie

import "example/pkg/alfa"

type Reserve = alfa.Reserve

type Keeper struct {
	Reserve
}

func Make() Keeper {
	return Keeper{Reserve: Reserve{Amount: 2}}
}
