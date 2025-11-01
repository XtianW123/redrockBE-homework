package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Stock int // 库存数量
}

func (p Product) TotalValue() float64 {
	return p.Price * float64(p.Stock)
	//return p.Price * p.Stock
	//无效运算: p.Price * p.Stock(类型 float64 和 int 不匹配)
	//他不会自动转换类型啊//规定的
	//为什么必须要return而不能print// 设计原则：方法应该专注于计算，调用者负责显示
}
func (p Product) IsInStock() bool {
	if p.Stock > 0 {
		return true
	}
	return false
}
func (p Product) Info() string {
	return fmt.Sprintf("Name: %s, Price: %.1f, Stock: %d", p.Name, p.Price, p.Stock)
	//为啥还得返回个输出函数   (ai填充的)
}
func (p *Product) Restock(amount int) {
	//	amount=0
	//fmt.Scan(&amount)应该在调用时传入参数
	p.Stock += amount
}
func (p *Product) Sell(amount int) (success bool, message string) {
	if p.Stock == 0 || amount > p.Stock {
		return false, "库存不足"
	}
	p.Stock -= amount
	return true, "售卖成功"

}
func main() {
	ccc := Product{
		Name:  "Go编程书",
		Price: 89.5,
		Stock: 10,
	}
	a, b := ccc.Sell(5)
	fmt.Println(a, b)
	ccc.Restock(20)
	fmt.Println(ccc.Stock)
	ccc.Sell(30)
	a, b = ccc.Sell(30)
	fmt.Println(a, b)
	fmt.Println(ccc.Info())
	fmt.Println(ccc.TotalValue())
}
