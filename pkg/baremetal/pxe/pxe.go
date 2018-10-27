package pxe

import (
	"fmt"
	"net"
	"time"

	"github.com/pin/tftp"
	"go.universe.tf/netboot/dhcp4"
)

const (
	portDHCP = 67
	portTFTP = 69
)

// Architecture describes a kind of CPU architecture
type Architecture int

// Architecture types that pxe knows how to boot
// These architectures are self-reported by the booting machine. The
// machine may support additional execution mode. For example, legacy
// PC BIOS reports itself as an ArchIA32, but may also support ArchX64
// execution
const (
	// ArchIA32 is a 32-bit x86 machine. It may also support X64
	// execution, but pxe has no way of kowning.
	ArchIA32 Architecture = iota
	// ArchX64 is a 64-bit x86 machine (aka amd64 aka x64)
	ArchX64
	ArchUnknown
)

func (a Architecture) String() string {
	switch a {
	case ArchIA32:
		return "IA32"
	case ArchX64:
		return "X64"
	default:
		return "Unknown architecture"
	}
}

// A Machine describes a machine that is attempting to boot
type Machine struct {
	MAC  net.HardwareAddr
	Arch Architecture
}

// Firmware describes a kind of firmware attempting to boot.
// This should only be used for selecting the right bootloader within
// pxe, kernel selection should key off the more generic
// Architecture
type Firmware int

// The bootloaders that pxe knows how to handle
const (
	FirmwareX86PC   Firmware = iota // "Classic" x86 BIOS with PXE/UNDI support
	FirmwareEFI32                   // 32-bit x86 processor running EFI
	FirmwareEFI64                   // 64-bit x86 processor running EFI
	FirmwareEFIBC                   // 64-bit x86 processor running EFI
	FirmwareX86Ipxe                 // "Classic" x86 BIOS running iPXE (no UNDI support)
	FirmwareUnknown
)

type Server struct {
	// Address to listen on, or empty for all interfaces
	Address     string
	DHCPPort    int
	TFTPPort    int
	TFTPRootDir string
	errs        chan error
}

func (s *Server) Serve() error {
	if s.Address == "" {
		s.Address = "0.0.0.0"
	}
	if s.DHCPPort == 0 {
		s.DHCPPort = portDHCP
	}
	if s.TFTPPort == 0 {
		s.TFTPPort = portTFTP
	}
	tftpHandler, err := NewTFTPHandler(s.TFTPRootDir)
	if err != nil {
		return err
	}
	tftpSrv := tftp.NewServer(tftpHandler.ReadHandler, nil)
	tftpSrv.SetTimeout(5 * time.Second)

	newDHCP := dhcp4.NewConn
	dhcp, err := newDHCP(fmt.Sprintf("%s:%d", s.Address, s.DHCPPort))
	if err != nil {
		return fmt.Errorf("New DHCP error: %v", err)
	}

	s.errs = make(chan error)

	go func() { s.errs <- s.serveDHCP(dhcp) }()
	go func() { s.errs <- s.serveTFTP(tftpSrv) }()

	err = <-s.errs
	dhcp.Close()
	tftpSrv.Shutdown()
	return err
}
