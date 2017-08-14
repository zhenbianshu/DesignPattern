package factory

type ShoesFactory interface {
	Product(material string)
}

type LiNingFactory struct {
}

func (factory LiNingFactory) Product(material string) Shoe {
	var shoe Shoe
	if material == "plastic" {
		shoe = liNingPlasticShoe{"lining plastic shoe", 120}
	} else if material == "cotton" {
		shoe = liNingCottonShoe{"linig cotton shoe", 110}
	} else {
		shoe = liNingPlasticShoe{"linig sports shoe", 100}
	}

	return shoe
}

type AdidasFactory struct {
}

func (factory AdidasFactory) Product(material string) Shoe {
	var shoe Shoe
	if material == "plastic" {
		shoe = adidasPlasticShoe{"adidas plastic shoe", 90}
	} else if material == "cotton" {
		shoe = adidasCottonShoe{"adidas cotton shoe", 80}
	} else {
		shoe = adidasSportsShoe{"adidas sports shoe", 70}
	}

	return shoe
}
