package bign

type BigN struct {
	Symbol int // 0正，1负
	Data   string
}

func reverse(s string) string {
	result := []byte{}
	nLen := len(s)
	for i := nLen - 1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return string(result)
}

func findMaxIgnoreSymbol(n1, n2 *BigN) *BigN {
	n1Len := len(n1.Data)
	n2Len := len(n2.Data)
	if n1Len > n2Len {
		return n1
	} else if n1Len < n2Len {
		return n2
	} else {
		for j := n1Len - 1; j >= 0; j-- {
			if n1.Data[j] > n2.Data[j] {
				return n1
			}
		}
		return n2
	}
}

func Add(x, y BigN) BigN {
	maxN := findMaxIgnoreSymbol(&x, &y)
	if maxN == &y { // 如果y的数据部分（不包括字符部分）比较大，翻转x, y
		x, y = y, x
	}

	x1 := reverse(x.Data)
	y1 := reverse(y.Data)
	nX := len(x1)
	nY := len(y1)

	// x,y 同正或同负
	if x.Symbol == y.Symbol {
		carry := 0
		data := make([]byte, nX)
		j := 0
		for ; j < nX && j < nY; j++ {
			c1 := int(x1[j] - '0')
			c2 := int(y1[j] - '0')
			data[j] = byte((c1+c2+carry)%10 + '0')
			carry = (c1 + c2 + carry) / 10
		}
		for j < nX {
			data[j] = byte(x1[j])
			j++
		}

		ret := BigN{x.Symbol, reverse(string(data))}
		return ret
	} else {
		// 求 x.data - y.data 即可
		copyX := make([]int, len(x1))
		for i, c := range x1 {
			copyX[i] = int(c)
		}

		data := make([]byte, nX)

		j := 0
		for ; j < nX && j < nY; j++ {
			c1 := int(copyX[j] - '0')
			c2 := int(y1[j] - '0')
			if c1 < c2 {
				c1 = c1 + 10
				copyX[j+1] = copyX[j+1] - 1
			}
			data[j] = byte(c1 - c2 + '0')
		}
		for j < nX {
			data[j] = byte(copyX[j])
			j++
		}
		ret := BigN{x.Symbol, reverse(string(data))}
		return ret
	}
}

func moveNumber(x BigN, n int) BigN {
	y := x
	for j := 0; j < n; j++ {
		y.Data = y.Data + "0"
	}
	return y
}

func SmallMultiply(x, y BigN) BigN {
	x1 := reverse(x.Data)
	y1 := reverse(y.Data)
	data := make([]int, len(x1)+len(y1))
	for i, _ := range x1 {
		for j, _ := range y1 {
			data[i+j] += int(x1[i]-'0') * int(y1[j]-'0')
		}
	}
	for k := 0; k < len(x1)+len(y1)-1; k++ {
		if data[k] > 10 {
			data[k+1] += data[k] / 10
			data[k] = data[k] % 10
		}
	}

	bdata := make([]byte, len(data))
	for i, _ := range data {
		bdata[i] = byte(data[i] + '0')
	}

	symbol := x.Symbol
	if x.Symbol != y.Symbol {
		symbol = 1
	}
	return BigN{symbol, reverse(string(bdata))}
}

func Multiply(x, y BigN) BigN {
	if len(x.Data) < 8 || len(y.Data) < 8 {
		return SmallMultiply(x, y)
	}

	a, b := x, x
	a.Data = a.Data[:len(a.Data)/2]
	b.Data = b.Data[len(b.Data)/2:]

	c, d := y, y
	c.Data = c.Data[:len(c.Data)/2]
	d.Data = d.Data[len(d.Data)/2:]

	ac := Multiply(a, c)
	bd := Multiply(b, d)

	negab := b
	negab.Symbol = ^negab.Symbol

	negac := c
	negac.Symbol = ^negac.Symbol

	step1 := moveNumber(ac, len(b.Data)+len(d.Data)) // move ac with 10^n
	step2 := Add(step1, bd)                          // add bd

	step3 := Add(ac, bd)
	step31 := Add(a, negab)
	step32 := Add(d, negac)
	step33 := Multiply(step31, step32)
	step3 = Add(step3, step33)
	step3 = moveNumber(step3, (len(b.Data)+1)/2+(len(d.Data)+1)/2)

	return Add(step2, step3)
}
