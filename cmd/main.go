package main

import (
	"BitginAssignment/cashier"
	"fmt"
)

func main() {
	system := &cashier.ChargingSystem{
		ChargingMap: make(map[string]cashier.ChargingMode),
	}
	goodPrice := float64(1000)
	pointLimit := float64(100)
	vipDiscountRate := 0.95
	pointDiscountRate := float64(2)
	system.AddChargingMode("Normal", &cashier.NormalCharging{})
	system.AddChargingMode("Vip", &cashier.VipCharging{Discount: vipDiscountRate})
	system.AddChargingMode("Point", &cashier.PointCharging{Rate: pointDiscountRate})

	price := system.Calculate("Normal", goodPrice)
	pointDiscount := system.GetChargingMode("Point").Calculate(pointLimit)
	fmt.Printf("商品金額: %.2f 元, 一般會員購買商品價格: %.2f 可使用最多 %.2f 點數折抵 %.2f 元  \n", goodPrice, price, pointDiscount, pointLimit)
	price = system.Calculate("Vip", goodPrice)
	fmt.Printf("商品金額: %.2f 元 VIP折數: %.2f VIP會員購買商品價格: %.2f 可使用最多 %.2f 點數折抵 %.2f 元  \n", goodPrice, vipDiscountRate, price, pointDiscount, pointLimit)
}
