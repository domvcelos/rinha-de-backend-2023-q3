package people

type PeopleHandler struct {
	Service PeopleServiceInterface
}

func NewHandler(ps PeopleServiceInterface) *PeopleHandler {
	return &PeopleHandler{
		Service: ps,
	}
}
