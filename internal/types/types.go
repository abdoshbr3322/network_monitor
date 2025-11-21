package types

type Stats struct {
	RX_bytes int64
	TX_bytes int64
}

func (st Stats) Add(st_2 Stats) Stats {
	return Stats{st.RX_bytes + st_2.RX_bytes, st.TX_bytes + st_2.RX_bytes}
}
