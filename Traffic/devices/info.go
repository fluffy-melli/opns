package devices

import (
	"strconv"
	"time"
)

type Traffic struct {
	Tx     float64
	Txc    string
	Txu    string
	TxBps  float64
	TxcBps string
	TxuBps string
	Rx     float64
	Rxc    string
	Rxu    string
	RxBps  float64
	RxcBps string
	RxuBps string
}

var units = []string{"(Byte)", "(KB)", "(MB)", "(GB)", "(TB)", "(PB)", "(EB)"}
var bpsunits = []string{"(Bps)", "(Kbps)", "(Mbps)", "(Gbps)", "(Tbps)", "(Pbps)", "(Ebps)"}

func (D *Device) ReadTraffic(d time.Duration, f func(Traffic)) {
	start_tx := D.Io.BytesSent
	start_rx := D.Io.BytesRecv
	start_time := time.Now()
	time.Sleep(d)
	var r Device
	for _, dv := range Devices() {
		if dv.Name == D.Name {
			r = dv
		}
	}
	end_tx := r.Io.BytesSent
	end_rx := r.Io.BytesRecv
	end_time := time.Now()
	elapsed := end_time.Sub(start_time).Seconds()
	txp, txu, _ := B_to_E(end_tx - start_tx)
	rxp, rxu, _ := B_to_E(end_rx - start_rx)
	txbp, _, txbu := B_to_E(uint64(float64(end_tx-start_tx) * 8 / elapsed))
	rxbp, _, rxbu := B_to_E(uint64(float64(end_rx-start_rx) * 8 / elapsed))
	f(Traffic{
		Tx:     txp,
		Txc:    strconv.FormatFloat(txp, 'f', -1, 64) + txu,
		Txu:    txu,
		TxBps:  txbp,
		TxcBps: strconv.FormatFloat(txbp, 'f', -1, 64) + bpsunits[txbu],
		TxuBps: bpsunits[txbu],
		Rx:     rxp,
		Rxc:    strconv.FormatFloat(rxp, 'f', -1, 64) + rxu,
		Rxu:    rxu,
		RxBps:  rxbp,
		RxcBps: strconv.FormatFloat(rxbp, 'f', -1, 64) + bpsunits[rxbu],
		RxuBps: bpsunits[rxbu],
	})
}

func (D *Device) Tx() (float64, string, int) {
	return B_to_E(D.Io.BytesSent)
}

func (D *Device) Rx() (float64, string, int) {
	return B_to_E(D.Io.BytesRecv)
}

func B_to_E(b uint64) (float64, string, int) {
	if b < 1024 {
		return float64(b), units[0], 0
	}
	div := float64(b)
	exp := 0

	for div >= 1024 && exp < len(units)-1 {
		div /= 1024
		exp++
	}

	return div, units[exp], exp
}
