package domain

type SKU int32
type Count uint16
type Name uint16
type Price uint16

type Item struct {
	Sku   SKU
	Count Count
}
