package def

type Niu struct {
	Number int
	Index  NiuIndex
	Score  int
}

type NiuIndex struct {
	Niu    []byte
	Remain []byte
}
