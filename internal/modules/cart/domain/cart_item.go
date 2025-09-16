package domain

type SKU int32
type Count uint16

type Item struct {
	Sku   SKU
	Count Count
}
