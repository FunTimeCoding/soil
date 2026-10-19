package server

type Error struct {
	Error string
}

type ErrorResponse struct {
	Error           string
	EventIdentifier string
}

type GetItem400JSONResponse Error

type GetItem404JSONResponse Error

type GetItem503JSONResponse Error

type GetOther503JSONResponse ErrorResponse
