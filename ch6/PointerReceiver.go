type Point struct {X, Y float64}

/*
if we want to avoid copying, we use pointer
*/

func (p *Point) scaleBy(factor float64){
	p.X *= factor
	p.Y *= factor 
}



//ways to access it 
r := &Point{1,2}
r.scaleBy(r)
fmt.Println(*r)

//or 
p:= Point{1,2}
pptr := &p 
pptr.scaleBy(2)
fmt.Println(p) 

//there is a way to obtain
//value from adress 
//so the following two are equivalent 

pptr.Distance(q) 
(*pptr).Distance(q)

//or 
p:= Point{1,2}
(&p).scaleBy(2)
fmt.Println(p)