package api

import (
    "net"
    "fmt"
    "errors"

    "google.golang.org/protobuf/proto"

    "github.com/ESP32-Zephyr/esp32_zephyr_goapi/cmds"
)

type Esp32Client struct {
    Transport   string
    Ipv4        string
    DestPort    uint16
}

func NewEsp32Client (transport string, ipv4 string, dest_port uint16) (*Esp32Client, error) {
    if transport != "tcp" && transport != "udp" {
        return nil, errors.New("Invalid transport type")
    }

    ip := net.ParseIP(ipv4)

    if ip == nil || ip.To4() == nil {
        return nil, errors.New("Invalid IPv4 address")
    }

    return &Esp32Client{
        Transport:     transport,
        Ipv4:          ipv4,
        DestPort:      dest_port,
    }, nil
}

// SendCmd sends a protobuf request and returns the response
func (c *Esp32Client) SendCmd(req *cmds.Request) (*cmds.Response, error) {
    // Serialize request
    reqBytes, err := proto.Marshal(req)
    if err != nil {
        return nil, fmt.Errorf("marshal request: %w", err)
    }

    addr := fmt.Sprintf("%s:%d", c.Ipv4, c.DestPort)
    var conn net.Conn
    if c.Transport == "udp" {
        conn, err = net.Dial("udp", addr)
    } else {
        conn, err = net.Dial("tcp", addr)
    }
    if err != nil {
        return nil, fmt.Errorf("dial: %w", err)
    }
    defer conn.Close()

    // Send request
    if _, err := conn.Write(reqBytes); err != nil {
        return nil, fmt.Errorf("write: %w", err)
    }

    // Read response
    buf := make([]byte, 1024)
    n, err := conn.Read(buf)
    if err != nil {
        return nil, fmt.Errorf("read: %w", err)
    }

    // Unmarshal response
    var resp cmds.Response
    if err := proto.Unmarshal(buf[:n], &resp); err != nil {
        return nil, fmt.Errorf("unmarshal response: %w", err)
    }

    if resp.Hdr.GetRet() != cmds.RetCode_OK {
        return nil, fmt.Errorf("Command failed! (ret: %d) %s", resp.Hdr.GetRet(), resp.Hdr.GetErrMsg())
    }

    return &resp, nil
}

func (c *Esp32Client) PrintInfo() {
    fmt.Printf("Transport: %s, IPv4: %s, Port: %d\n", c.Transport, c.Ipv4, c.DestPort)
}

func (c *Esp32Client) VersionGet() (*cmds.VersionGetRes, error) {
    cmd_id := cmds.CommandId_VERSION_GET
    req := &cmds.Request{
        Hdr: &cmds.ReqHdr{
            Id: &cmd_id,
        },
        Pl: &cmds.Request_VersionGet{
            VersionGet: &cmds.VersionGetReq{},
        },
    }

    resp, err := c.SendCmd(req)
    if err != nil {
        return nil, err
    }

    version := resp.GetVersionGet()
    if version == nil {
        return nil, errors.New("Invalid response: Failed to retrieve version info")
    }

    return version, nil
}

func (c *Esp32Client) AdcChsGet() (*cmds.AdcChsGetRes, error) {
    cmd_id := cmds.CommandId_ADC_CHS_GET
    req := &cmds.Request{
        Hdr: &cmds.ReqHdr{
            Id: &cmd_id,
        },
        Pl: &cmds.Request_AdcChsGet{
            AdcChsGet: &cmds.AdcChsGetReq{},
        },
    }

    resp, err := c.SendCmd(req)
    if err != nil {
        return nil, err
    }

    result := resp.GetAdcChsGet()
    if result == nil {
        return nil, errors.New("Invalid response: Failed to retrieve ADC channels info")
    }
    return result, nil
}

func (c *Esp32Client) AdcChRead(ch uint32) (*cmds.AdcChReadRes, error) {
    cmd_id := cmds.CommandId_ADC_CH_READ
    req := &cmds.Request{
        Hdr: &cmds.ReqHdr{
            Id: &cmd_id,
        },
        Pl: &cmds.Request_AdcChRead{
            AdcChRead: &cmds.AdcChReadReq{
                Ch: &ch,
            },
        },
    }

    resp, err := c.SendCmd(req)
    if err != nil {
        return nil, err
    }

    result := resp.GetAdcChRead()
    if result == nil {
        return nil, errors.New("Invalid response: Failed to read ADC channel")
    }
    return result, nil
}

func (c *Esp32Client) PwmChsGet() (*cmds.PwmChsGetRes, error) {
    cmd_id := cmds.CommandId_PWM_CHS_GET
    req := &cmds.Request{
        Hdr: &cmds.ReqHdr{
            Id: &cmd_id,
        },
        Pl: &cmds.Request_PwmChsGet{
            PwmChsGet: &cmds.PwmChsGetReq{},
        },
    }

    resp, err := c.SendCmd(req)
    if err != nil {
        return nil, err
    }

    result := resp.GetPwmChsGet()
    if result == nil {
        return nil, errors.New("Invalid response: Failed to retrieve PWM channels info")
    }
    return result, nil
}

func (c *Esp32Client) PwmChSet(ch uint32, period uint32, pulse uint32) (*cmds.PwmChSetRes, error) {
    cmd_id := cmds.CommandId_PWM_CH_SET
    req := &cmds.Request{
        Hdr: &cmds.ReqHdr{
            Id: &cmd_id,
        },
        Pl: &cmds.Request_PwmChSet{
            PwmChSet: &cmds.PwmChSetReq{
                Ch:     &ch,
                Period: &period,
                Pulse:  &pulse,
            },
        },
    }

    resp, err := c.SendCmd(req)
    if err != nil {
        return nil, err
    }

    result := resp.GetPwmChSet()
    if result == nil {
        return nil, errors.New("Invalid response: Failed to set PWM channel")
    }
    return result, nil
}

func (c *Esp32Client) PwmChGet(ch uint32) (*cmds.PwmChGetRes, error) {
    cmd_id := cmds.CommandId_PWM_CH_GET
    req := &cmds.Request{
        Hdr: &cmds.ReqHdr{
            Id: &cmd_id,
        },
        Pl: &cmds.Request_PwmChGet{
            PwmChGet: &cmds.PwmChGetReq{
                Ch: &ch,
            },
        },
    }

    resp, err := c.SendCmd(req)
    if err != nil {
        return nil, err
    }

    result := resp.GetPwmChGet()
    if result == nil {
        return nil, errors.New("Invalid response: Failed to get PWM channel")
    }
    return result, nil
}

func (c *Esp32Client) PwmPeriodIntervalGet() (*cmds.PwmPeriodsGetRes, error) {
    cmd_id := cmds.CommandId_PWM_PERIOD_INTERVAL_GET
    req := &cmds.Request{
        Hdr: &cmds.ReqHdr{
            Id: &cmd_id,
        },
        Pl: &cmds.Request_PwmPeriodsGet{
            PwmPeriodsGet: &cmds.PwmPeriodsGetReq{},
        },
    }

    resp, err := c.SendCmd(req)
    if err != nil {
        return nil, err
    }

    result := resp.GetPwmPeriodsGet()
    if result == nil {
        return nil, errors.New("Invalid response: Failed to get PWM period interval")
    }
    return result, nil
}
