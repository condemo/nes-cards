package types

type Structure interface {
	TakeDMG(uint8)
	AddDefense(uint8)
}

type MeatStructure interface {
	AddAlteredEffect(AlteredEffect)
	CleanAltered()
	ApplyAlteredStack()
}

type AlteredEffect interface {
	Apply() uint8
}

type DamageEffect interface {
	GetDmg() uint8
	AddDmg(uint8)
}
