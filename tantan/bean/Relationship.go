package bean

type Relationship struct {
	ID uint64 `pg:"id" json:"id"` 
	UIdF uint64 `pg:"u_id_f" json:"u_id_f"`
	UIdT uint64 `pg:"u_id_t" json:"u_id_t"`
	StateF uint8 `pg:"state_f" json:"state_f"`
	StateT uint8 `pg:"state_t" json:"state_t"`
	State uint8 `pg:"-" json:"state"`
	Type uint8 `pg:"-" json:"type""`
}

func NewRelationship(uIdF uint64, uIdT uint64, stateF uint8, stateT uint8) (*Relationship){
	return &Relationship{
		UIdF : uIdF,
		UIdT : uIdT,
		StateF : stateF,
		StateT : stateT,
		Type: uint8(1),
	}
}

func (rs *Relationship) setStateF(stateF uint8)  {
	rs.StateF=stateF
}

func (rs *Relationship) setStateT(stateT uint8)  {
	rs.StateT=stateT
}