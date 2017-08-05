package decorator

type pack struct {
	desc string
	gift Goods
}

func NewPack(gift Goods, pack_desc string) pack {
	pack := pack{desc: pack_desc, gift: gift}
	return pack
}

func (pack pack) GetDesc() string {
	return pack.gift.GetDesc() + ", with " + pack.desc + " packed"
}
