# BitGin Assignment
## 敘述
針對平台產品銷售有不同的收費機制，該系統中有平台點數及平台幣（平台幣是主要扣款使用的幣別），依照不同商品及不同的用戶等級會有不同的收費模式。

主要收費模式會分成下列三種。

Ａ、正常平台幣收費

Ｂ、VIP會員有平台幣優惠價格 (例如: VIP1: 95折，VIP2: 9折，VIP3: 85折，各個等級的折扣會依照活動做調整。)

Ｃ、扣平台點數折抵平台幣優惠（例如：設定平台點數1:1折抵，則1000元商品可以使用100點扣抵扣在另外支付900元購買，折抵比例依照活動做調整。）



平台會不定時舉辦各式的優惠活動，活動時會彈性調整方案B及C的優惠內容。依照不同的活動內容可以排程調整活動。

1. 請使用Golang實作一個收銀系統（請考慮Clean Code及Design Pattern）。

2. 承上題，平台後來新增了另一個收費模式，如果有VIP身份扣100點以上折抵，另外享再九折優惠。（考慮SOLID）

## How to use
```go
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
```
output
```go
商品金額: 1000.00 元, 一般會員購買商品價格: 1000.00 可使用最多 200.00 點數折抵 100.00 元  
商品金額: 1000.00 元 VIP折數: 0.95 VIP會員購買商品價格: 950.00 可使用最多 200.00 點數折抵 100.00 元
```