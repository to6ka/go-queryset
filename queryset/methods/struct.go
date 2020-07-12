package methods

// StructModifierMethod represents method, modifying current struct
type StructModifierMethod struct {
	namedMethod
	structMethod
	dbArgMethod
	gormErroredMethod
}

// NewStructModifierMethod create StructModifierMethod method
func NewStructModifierMethod(name, structTypeName string, cfg Config) StructModifierMethod {
	r := StructModifierMethod{
		namedMethod:       newNamedMethod(name),
		dbArgMethod:       newDbArgMethod(cfg),
		structMethod:      newStructMethod("o", "*"+structTypeName),
		gormErroredMethod: newGormErroredMethod(name, "o", "db", cfg),
	}
	return r
}
