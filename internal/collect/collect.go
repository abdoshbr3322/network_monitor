package collect

import (
	"os"
	"strconv"
	"strings"
)

func CollectNetworkStats() (int, int, error) {
	var rx_bytes, tx_bytes int
	file_content, err := os.ReadFile("/proc/net/dev")
	if err != nil {
		return 0, 0, err
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
		rx, err := strconv.Atoi(fields[1])
		if err != nil {
			return 0, 0, err
		}

		tx, err := strconv.Atoi(fields[9])
		if err != nil {
			return 0, 0, err
		}

		rx_bytes += rx
		tx_bytes += tx
	}
	return rx_bytes, tx_bytes, nil
}
