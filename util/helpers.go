package util

var itemsPerPage = 10

func PageLength(page int) (n, n2 int) {
	start := (page - 1) * itemsPerPage
	stop := start + itemsPerPage
	start = start + 1

	return start, stop
}

func SqlOrder(order string) (string, from string) {
	switch order {
	case ``:
		break
	default:
		from = `DESC`
		break
	}
	return order, from
}
