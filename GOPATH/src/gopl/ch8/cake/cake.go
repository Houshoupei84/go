package cake

import (
	"time"
	"fmt"
	"math/rand"
)

type Shop struct {

	Verbose 	bool
	Cakes   	int           	//要烘烤的蛋糕总数
	BakeTime 	time.Duration   //烘烤一个蛋糕的时间
	BakeStdDev  time.Duration   //烘烤的标准偏差时间
	BakeBuf     int             //buffer slots between baking and icing 有几个空盒子
	NumIcers    int             //number of cooks doing icing
	IceTime     time.Duration   // time to ice one cake
	IceStdDev   time.Duration   // standard deviation of icing time
	IceBuf      int             // buffer slots between icing and inscribing
	InscribeTime time.Duration  // time to inscribe one cake
	InscribeStdDev time.Duration// standard deviation of inscribing time
}

type cake int

//烘焙
func (s *Shop) baker( baked chan <- cake) {

	for i := 0 ; i < s.Cakes ; i++ {

		c := cake(i)

		if s.Verbose{
			fmt.Println("baking", c)
		}
		work(s.BakeTime, s.BakeStdDev)
		baked <- c
	}
	close(baked)
}

//上糖衣
func (s * Shop) icer( iced chan <- cake, baked <-chan cake){

	for c := range baked{
		if s.Verbose{
			fmt.Println("icing", c)
		}
		work(s.IceTime, s.IceStdDev)
		iced <- c
	}

}


func (s * Shop) inscriber (iced <- chan cake){

	for i := 0 ; i < s.Cakes; i++{

		c :=  <- iced

		if s.Verbose{
			fmt.Println("inscribing",  c)
		}

		work(s.InscribeTime, s.InscribeStdDev)
		if s.Verbose {
			fmt.Println("finished", c)
		}
	}
}

func (s* Shop) Work( runs int ) {

	for run := 0 ; run < runs; run++ {

		baked := make(chan cake, s.BakeBuf)
		iced := make(chan cake, s.IceBuf)

		go s.baker(baked)

		for i :=0 ; i < s.NumIcers; i++{
			go s.icer(iced, baked)
		}
		s.inscriber(iced)
	}
}

func work(d ,  stddev time.Duration){

	delay := d + time.Duration(rand.NormFloat64()* float64(stddev))
	time.Sleep(delay)
}