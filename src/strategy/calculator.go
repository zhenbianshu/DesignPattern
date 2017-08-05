package strategy

// 计算器类，有算法属性，且其算法随时可被替换，以使用其他的算法
type Calculator struct {
	alg IAlg
}

// 计算器构造方法，默认使用二分法计算
func NewCalculator() Calculator {
	alg := NewBinarySearch()
	calculator := Calculator{alg}

	return calculator
}

// 改变算法
func (calculator *Calculator) ChangeAlg(alg IAlg) {
	calculator.alg = alg
}

func (calculator Calculator) Cal() {
	calculator.alg.Used()
}
