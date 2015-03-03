package fbp

import (
	"fmt"
)

type BepaalDiagnose struct {
	*Component
}

// BepaalDiagnose berekent de diagnose a.h.v. het aantal domeinen en de totale normtijd
// In : fields: aantal_domeinen, totale_normtijd
// Out: fields: diagnose
func NewBepaalDiagnose() *BepaalDiagnose {
	bd := new(BepaalDiagnose)
	bd.Component = NewComponent()
	bd.Owner = bd
	return bd
}

func (b *BepaalDiagnose) HandleData(port *Port) {
	fmt.Println("Data uit Counter", port.name)
	//	data := port.GetData()
	// Hier komt de business logica

	// Onderstaande regel zorgt ervoor dat de data naar uitgaande poorten gepusht wordt.
	b.Component.HandleData(port)
}
