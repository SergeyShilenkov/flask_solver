package validator

import "strconv"

func ValidateData(data *[][4]int) error {
	if len(*data) < 6 {
		return &AmountFlasks{Amount: len(*data)}
	}

	amountColorsMap := make(map[int]int)

	for _, f := range *data {
		for _, c := range f {
			amountColorsMap[c]++
		}
	}

	if amountColorsMap[0] != 8 {
		return &AmountEmptySlots{Amount: amountColorsMap[0]}
	}

	for i := 2; i < len(amountColorsMap); i++ {
		if amountColorsMap[i] > 4 {
			return &AmountSpecificColorError{Color: strconv.Itoa(i), Amount: amountColorsMap[i], maxAmount: 4}
		}
	}

	return nil
}
