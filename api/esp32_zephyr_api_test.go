package api

import (
	"testing"
)

const testIPv4 = "192.168.0.12"
const testPort = 4242

func TestNewEsp32Client(t *testing.T) {
	tests := []struct {
		name      string
		transport string
		ipv4      string
		destPort  uint16
		wantErr   bool
	}{
		{
			name:      "valid tcp",
			transport: "tcp",
			ipv4:      "192.168.0.1",
			destPort:  testPort,
			wantErr:   false,
		},
		{
			name:      "valid udp",
			transport: "udp",
			ipv4:      "10.0.0.1",
			destPort:  5000,
			wantErr:   false,
		},
		{
			name:      "invalid transport",
			transport: "http",
			ipv4:      "192.168.0.1",
			destPort:  testPort,
			wantErr:   true,
		},
		{
			name:      "invalid ipv4",
			transport: "tcp",
			ipv4:      "invalid.ip",
			destPort:  testPort,
			wantErr:   true,
		},
		{
			name:      "ipv6 address",
			transport: "tcp",
			ipv4:      "::1",
			destPort:  testPort,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewEsp32Client(tt.transport, tt.ipv4, tt.destPort)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEsp32Client() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && client == nil {
				t.Error("NewEsp32Client() returned nil client")
			}
		})
	}
}

func TestEsp32ClientPrintInfo(t *testing.T) {
	client := &Esp32Client{
		Transport: "tcp",
		Ipv4:      "192.168.0.1",
		DestPort:  testPort,
	}
	// Test that PrintInfo doesn't panic
	client.PrintInfo()
}

func TestTransportTypes(t *testing.T) {
	validTransports := []string{"tcp", "udp"}
	for _, transport := range validTransports {
		_, err := NewEsp32Client(transport, "192.168.0.1", testPort)
		if err != nil {
			t.Errorf("valid transport %s should not error: %v", transport, err)
		}
	}

	invalidTransports := []string{"", "http", "websocket", "xyz"}
	for _, transport := range invalidTransports {
		_, err := NewEsp32Client(transport, "192.168.0.1", testPort)
		if err == nil {
			t.Errorf("invalid transport %s should error", transport)
		}
	}
}

func TestVersionGet(t *testing.T) {
	tests := []struct {
		name    string
		client  *Esp32Client
		wantErr bool
	}{
		{
			name: "VersionGet",
			client: &Esp32Client{
				Transport: "tcp",
				Ipv4:      testIPv4,
				DestPort:  testPort,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.client.VersionGet()
			if (err != nil) != tt.wantErr {
				t.Errorf("VersionGet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdcChsGet(t *testing.T) {
	tests := []struct {
		name    string
		client  *Esp32Client
		channels uint32
		wantErr  bool
	}{
		{
			name: "AdcChsGet",
			client: &Esp32Client{
				Transport: "tcp",
				Ipv4:      testIPv4,
				DestPort:  testPort,
			},
			channels: 2,
			wantErr: true, 
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.client.AdcChsGet()
			if (err != nil) != tt.wantErr {
				t.Errorf("AdcChsGet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdcChRead(t *testing.T) {
	tests := []struct {
		name    string
		client  *Esp32Client
		ch      uint32
		wantErr bool
	}{
		{
			name: "AdcChRead channel 0",
			client: &Esp32Client{
				Transport: "tcp",
				Ipv4:      testIPv4,
				DestPort:  testPort,
			},
			ch:      0,
			wantErr: true,
		},
		{
			name: "AdcChRead channel 1",
			client: &Esp32Client{
				Transport: "tcp",
				Ipv4:      testIPv4,
				DestPort:  testPort,
			},
			ch:      1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.client.AdcChRead(tt.ch)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdcChRead() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPwmChsGet(t *testing.T) {
	tests := []struct {
		name    string
		client  *Esp32Client
		wantErr bool
	}{
		{
			name: "PwmChsGet",
			client: &Esp32Client{
				Transport: "tcp",
				Ipv4:      testIPv4,
				DestPort:  testPort,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.client.PwmChsGet()
			if (err != nil) != tt.wantErr {
				t.Errorf("PwmChsGet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPwmChsSet(t *testing.T) {
	tests := []struct {
		name    string
		client  *Esp32Client
		ch 		uint32
		period  uint32
		pulse 	uint32
		wantErr bool
	}{
		{
			name: "PwmChSet",
			client: &Esp32Client{
				Transport: "tcp",
				Ipv4:      testIPv4,
				DestPort:  testPort,
			},
			ch:     0,
			period: 1000,
			pulse:  500,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.client.PwmChSet(tt.ch, tt.period, tt.pulse)
			if (err != nil) != tt.wantErr {
				t.Errorf("PwmChSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}


func TestPwmChGet(t *testing.T) {
	tests := []struct {
		name    string
		client  *Esp32Client
		ch 		uint32
		wantErr bool
	}{
		{
			name: "PwmChGet",
			client: &Esp32Client{
				Transport: "tcp",
				Ipv4:      testIPv4,
				DestPort:  testPort,
			},
			ch:     0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.client.PwmChGet(tt.ch)
			if (err != nil) != tt.wantErr {
				t.Errorf("PwmChGet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPwmPeriodIntervalGet(t *testing.T) {
	tests := []struct {
		name    string
		client  *Esp32Client
		wantErr bool
	}{
		{
			name: "PwmPeriodIntervalGet",
			client: &Esp32Client{
				Transport: "tcp",
				Ipv4:      testIPv4,
				DestPort:  testPort,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.client.PwmPeriodIntervalGet()
			if (err != nil) != tt.wantErr {
				t.Errorf("PwmPeriodIntervalGet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
