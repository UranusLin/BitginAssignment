package cashier

// ChargingMode 定義 ChargingMode 介面，用於計算收費
type ChargingMode interface {
	Calculate(price float64) float64
}

// NormalCharging 正常收費模式
type NormalCharging struct{}

// Calculate 計算正常收費的方法
func (n *NormalCharging) Calculate(price float64) float64 {
	return price
}

// VipCharging VIP 收費模式
type VipCharging struct {
	Discount float64 // 折扣率
}

// Calculate 計算 VIP 收費的方法
func (v *VipCharging) Calculate(price float64) float64 {
	return price * v.Discount
}

// PointCharging 點數折扣收費模式
type PointCharging struct {
	Rate float64 // 點數折扣比例
}

// Calculate 計算點數折扣收費上限的方法
func (p *PointCharging) Calculate(limit float64) float64 {
	// 計算方式為 limit * Rate
	return limit * p.Rate
}

// ChargingSystem 結構，用於管理不同的收費模式
type ChargingSystem struct {
	ChargingMap map[string]ChargingMode // 收費模式名稱與對應的 ChargingMode 介面
}

// AddChargingMode 新增收費模式
func (c *ChargingSystem) AddChargingMode(name string, mode ChargingMode) {
	c.ChargingMap[name] = mode
}

// GetChargingMode 取得收費模式
func (c *ChargingSystem) GetChargingMode(name string) ChargingMode {
	return c.ChargingMap[name]
}

// Calculate 計算收費
func (c *ChargingSystem) Calculate(name string, price float64) float64 {
	mode := c.GetChargingMode(name)
	return mode.Calculate(price)
}
