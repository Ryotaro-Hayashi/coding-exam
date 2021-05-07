package leetcode

type ParkingSystem struct {
	cap map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	// キーをAddCarの入力と合わせる
	return ParkingSystem{
		cap: map[int]int{1: big, 2: medium, 3: small},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.cap[carType] <= 0 {
		return false
	} else {
		this.cap[carType]--
		return true
	}
}

//type ParkingSystem struct {
//	Big int
//	Medium int
//	Small int
//}
//
//func Constructor(big int, medium int, small int) ParkingSystem {
//	return ParkingSystem{
//		Big: big,
//		Medium: medium,
//		Small: small,
//	}
//}
//
//func (this *ParkingSystem) AddCar(carType int) bool {
//	switch carType {
//	case 1:
//		if this.Big <= 0 {
//			return false
//		} else {
//			this.Big--
//		}
//	case 2:
//		if this.Medium <= 0 {
//			return false
//		} else {
//			this.Medium--
//		}
//	case 3:
//		if this.Small <= 0 {
//			return false
//		} else {
//			this.Small--
//		}
//	}
//
//	return true
//}
