package collect

import (
	"os"
	"strconv"
	"strings"

	"github.com/abdoshbr3322/network_monitor/internal/types"
)

func CollectNetworkStats() (types.Stats, error) {
	var st types.Stats
	file_content, err := os.ReadFile("/proc/net/dev")
	if err != nil {
		return types.Stats{RX_bytes: 0, TX_bytes: 0}, err
	}

	rows := strings.Split(string(file_content), "\n")
	for _, row := range rows[2:] {
		fields := strings.Fields(string(row))
		if len(fields) == 0 {
			continue
		}
		iface := fields[0]

		// skip local interface
		if iface == "lo:" {
			continue
		}
		rx, err := strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			return types.Stats{RX_bytes: 0, TX_bytes: 0}, err
		}

		tx, err := strconv.ParseInt(fields[9], 10, 64)
		if err != nil {
			return types.Stats{RX_bytes: 0, TX_bytes: 0}, err
		}

		st.RX_bytes += rx
		st.TX_bytes += tx
	}
	return st, nil
}
