package domain

type UserID int64
type TotalPrice int64

type Cart struct {
	userID UserID
	items  map[SKU]Item
}

func (c *Cart) AddItem(sku SKU, count Count) {
	if item, ok := c.items[sku]; !ok {
		item.count += count
	}

	c.items[sku] = Item{
		sku:   sku,
		count: count,
	}
}

func (c *Cart) DeleteItem(sku SKU) {
	delete(c.items, sku)
}
