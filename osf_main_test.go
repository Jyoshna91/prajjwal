package ospf
import "ospf/ospf_lib"

import (
    "fmt"
//    "golang.org/x/crypto/ssh"
//    "log"
    "testing"
//    "time"
//    "strings"
//    "bytes"
//    "bufio"

)
/*
func Test_Configure_ospf_routers(t *testing.T){
    fmt.Printf("Enabling ospf on routers\n")
    err := ospf_lib.Enable_ospf("10.133.35.148")
    if err != nil {
        t.Errorf("Failed to enable OSPF on device: %v", err)
    }
    err1 := ospf_lib.Enable_ospf("10.133.35.143")
    if err1 != nil {
        t.Errorf("Failed to enable OSPF on device: %v", err1)
    }
    err2 := ospf_lib.Enable_ospf("10.133.35.150")
    if err2 != nil {
        t.Errorf("Failed to enable OSPF on device: %v", err2)
    }
    fmt.Printf("Configuring ospf on routers\n")
    err3 := ospf_lib.Configure_ospf_router("10.133.35.148","2")
    if err3 != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err3)
    }
    err4 := ospf_lib.Configure_ospf_router("10.133.35.143","2")
    if err4 != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err4)
    }
    err5 := ospf_lib.Configure_ospf_router("10.133.35.150","2")
    if err5 != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err5)
    }
    fmt.Printf("Configuring ospf on interface\n") //Configure_ospf_interface
    err6 := ospf_lib.Configure_ospf_interface("10.133.35.148","Ethernet 1/11","2")
    if err6 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err6)
    }
    err7 := ospf_lib.Configure_ospf_interface("10.133.35.148","Ethernet 1/1","2")
    if err7 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err7)
    }
    err8 := ospf_lib.Configure_ospf_interface("10.133.35.150","Ethernet 1/1","2")
    if err8 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err8)
    }
    err9 := ospf_lib.Configure_ospf_interface("10.133.35.143","Ethernet 1/11","2")
    if err9 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err9)
    }
    ospf_lib.Print_Output("10.133.35.148","show ip ospf neighbors")
    err10 := ospf_lib.Unconfigure_ospf_interface("10.133.35.148","Ethernet 1/11","2")
    if err10 != nil {
        t.Errorf("Failed to unconfigure OSPF on interface: %v", err10)
    }
    err11 := ospf_lib.Unconfigure_ospf_interface("10.133.35.148","Ethernet 1/1","2")
    if err11 != nil {
        t.Errorf("Failed to unconfigure OSPF on interface: %v", err11)
    }
    err12 := ospf_lib.Unconfigure_ospf_interface("10.133.35.150","Ethernet 1/1","2")
    if err12 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err12)
    }
    err13 := ospf_lib.Unconfigure_ospf_interface("10.133.35.143","Ethernet 1/11","2")
    if err13 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err13)
    }
    err14 := ospf_lib.Unconfigure_ospf_router("10.133.35.148","2")
    if err14 != nil {
        t.Errorf("Failed to unconfigure OSPF on device: %v", err14)
    }
    err15 := ospf_lib.Unconfigure_ospf_router("10.133.35.143","2")
    if err15 != nil {
        t.Errorf("Failed to unconfigure OSPF on device: %v", err15)
    }
    err16 := ospf_lib.Unconfigure_ospf_router("10.133.35.150","2")
    if err16 != nil {
        t.Errorf("Failed to unconfigure OSPF on device: %v", err16)
    }
}
*/
func Test_Multiarea(t *testing.T){
    fmt.Printf("Enabling ospf on routers\n")
    err111 := ospf_lib.Enable_ospf("10.133.35.148")
    if err111 != nil {
        t.Errorf("Failed to enable OSPF on device: %v", err111)
    }
    err121 := ospf_lib.Enable_ospf("10.133.35.143")
    if err121 != nil {
        t.Errorf("Failed to enable OSPF on device: %v", err121)
    }
    err131 := ospf_lib.Enable_ospf("10.133.35.150")
    if err131 != nil {
        t.Errorf("Failed to enable OSPF on device: %v", err131)
    }
    fmt.Printf("Configuring ospf on routers\n")
    err := ospf_lib.Configure_ospf_router("10.133.35.148","2")
    if err != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err)
    }
    err1 := ospf_lib.Configure_ospf_router("10.133.35.143","2")
    if err1 != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err1)
    }
    err2 := ospf_lib.Configure_ospf_router("10.133.35.150","2")
    if err2 != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err2)
    }
    fmt.Printf("Configure ospf in interface\n")
    err3 := ospf_lib.Configure_ospf_interface("10.133.35.148","Ethernet 1/11","2")
    if err3 != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err3)
    } 
    err4 := ospf_lib.Configure_ospf_multiarea("10.133.35.148","Ethernet 1/1","2")
    if err4 != nil {
        t.Errorf("Failed to configure OSPF multiarea on device: %v", err4)
    }
    err5 := ospf_lib.Configure_ospf_interface("10.133.35.150","Ethernet 1/1","2")
    if err5 != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err5)
    }
    err6 := ospf_lib.Configure_ospf_interface("10.133.35.150","Ethernet 1/7","2")
    if err6 != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err6)
    }
    err7 := ospf_lib.Configure_ospf_interface("10.133.35.143","Ethernet 1/11","2")
    if err7 != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err7)
    }
    err8 := ospf_lib.Configure_ospf_multiarea("10.133.35.143","Ethernet 1/11","2")
    if err8 != nil {
        t.Errorf("Failed to configure OSPF multiarea on device: %v", err8)
    }
    err9 := ospf_lib.Configure_ospf_loopback("10.133.35.150","loopback 0","2")
    if err9 != nil {
        t.Errorf("Failed to configure OSPF loopback on device: %v", err9)
    }
    ospf_lib.Print_Output("10.133.35.148","show ip ospf route")
    err10 := ospf_lib.Unconfigure_ospf_interface("10.133.35.148","Ethernet 1/11","2")
    if err10 != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err10)
    }
    err11 := ospf_lib.Unconfigure_ospf_multiarea("10.133.35.148","Ethernet 1/1","2")
    if err11 != nil {
        t.Errorf("Failed to configure OSPF multiarea on device: %v", err11)
    }
    err12 := ospf_lib.Unconfigure_ospf_interface("10.133.35.150","Ethernet 1/1","2")
    if err12 != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err12)
    }
    err13 := ospf_lib.Unconfigure_ospf_interface("10.133.35.150","Ethernet 1/7","2")
    if err13 != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err13)
    }
    err14 := ospf_lib.Unconfigure_ospf_interface("10.133.35.143","Ethernet 1/11","2")
    if err14 != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err14)
    }
    err15 := ospf_lib.Unconfigure_ospf_multiarea("10.133.35.143","Ethernet 1/7","2")
    if err15 != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err15)
    }
    err16 := ospf_lib.Unconfigure_ospf_loopback("10.133.35.143","loopback 0","2")
    if err16 != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err16)
    }   
}
/*
func Test_ospf_Authentication(t *testing.T) {
    fmt.Printf("Enabling ospf on routers\n")
    err := ospf_lib.Enable_ospf("10.133.35.148")
    if err != nil {
        t.Errorf("Failed to enable OSPF on device: %v", err)
    }
    err1 := ospf_lib.Enable_ospf("10.133.35.143")
    if err1 != nil {
        t.Errorf("Failed to enable OSPF on device: %v", err1)
    }
    err2 := ospf_lib.Enable_ospf("10.133.35.150")
    if err2 != nil {
        t.Errorf("Failed to enable OSPF on device: %v", err2)
    }
    fmt.Printf("Configuring ospf on routers\n")
    err3 := ospf_lib.Configure_ospf_router("10.133.35.148","2")
    if err3 != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err3)
    }
    err4 := ospf_lib.Configure_ospf_router("10.133.35.143","2")
    if err4 != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err4)
    }
    err5 := ospf_lib.Configure_ospf_router("10.133.35.150","2")
    if err5 != nil {
        t.Errorf("Failed to configure OSPF on device: %v", err5)
    }
    fmt.Printf("Configuring ospf on interface\n") //Configure_ospf_interface
    err6 := ospf_lib.Configure_ospf_interface("10.133.35.148","Ethernet 1/11","2")
    if err6 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err6)
    }
    err7 := ospf_lib.Configure_ospf_interface("10.133.35.148","Ethernet 1/1","2")
    if err7 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err7)
    }
    err8 := ospf_lib.Configure_ospf_interface("10.133.35.150","Ethernet 1/1","2")
    if err8 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err8)
    }
    err9 := ospf_lib.Configure_ospf_interface("10.133.35.150","Ethernet 1/7","2")
    if err9 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err9)
    }
    err10 := ospf_lib.Configure_ospf_interface("10.133.35.143","Ethernet 1/11","2")
    if err10 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err10)
    }
    err11 := ospf_lib.Configure_ospf_interface("10.133.35.143","Ethernet 1/7","2")
    if err11 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err11)
    }
    fmt.Printf("Configure ospf authentication in interface\n")
    err12 := ospf_lib.Configure_ospf_authentication("10.133.35.148","Ethernet 1/11","2")
    if err12 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err12)
    }
    err13 := ospf_lib.Configure_ospf_authentication("10.133.35.148","Ethernet 1/1","2")
    if err13 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err13)
    }
    err14 := ospf_lib.Configure_ospf_authentication("10.133.35.150","Ethernet 1/1","2")
    if err14 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err14)
    }
    err15 := ospf_lib.Configure_ospf_authentication("10.133.35.150","Ethernet 1/7","2")
    if err15 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err15)
    }
    err16 := ospf_lib.Configure_ospf_authentication("10.133.35.143","Ethernet 1/11","2")
    if err16 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err16)
    }
    err17 := ospf_lib.Configure_ospf_authentication("10.133.35.143","Ethernet 1/7","2")
    if err17 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err17)
    }
    ospf_lib.Print_Output("10.133.35.148","show ip ospf 2 interface eth1/11")
    err18 := ospf_lib.Unconfigure_ospf_authentication("10.133.35.148","Ethernet 1/11","2")
    if err18 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err18)
    }
    err19 := ospf_lib.Unconfigure_ospf_authentication("10.133.35.148","Ethernet 1/1","2")
    if err19 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err19)
    }
    err20 := ospf_lib.Unconfigure_ospf_authentication("10.133.35.150","Ethernet 1/1","2")
    if err20 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err20)
    }
    err21 := ospf_lib.Unconfigure_ospf_authentication("10.133.35.150","Ethernet 1/7","2")
    if err21 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err21)
    }
    err22 := ospf_lib.Unconfigure_ospf_authentication("10.133.35.143","Ethernet 1/11","2")
    if err22 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err22)
    }
    err23 := ospf_lib.Unconfigure_ospf_authentication("10.133.35.143","Ethernet 1/7","2")
    if err23 != nil {
        t.Errorf("Failed to configure OSPF on interface: %v", err23)
    }
}
func Test_ospf_timers(t *testing.T){
    fmt.Printf("Configure ospf timers on routers\n")
    err := ospf_lib.Configure_ospf_timers("10.133.35.148","Ethernet 1/11")
    if err != nil {
        t.Errorf("Failed to configure OSPF timers on device: %v", err)
    }
    err1 := ospf_lib.Configure_ospf_timers("10.133.35.143","Ethernet 1/11")
    if err1 != nil {
        t.Errorf("Failed to configure OSPF timers on device: %v", err1)
    }
    ospf_lib.Print_Output("10.133.35.148","show ip ospf interface ethernet 1/11")
    err2 := ospf_lib.Unconfigure_ospf_timers("10.133.35.148","Ethernet 1/11")
    if err2 != nil {
        t.Errorf("Failed to unconfigure OSPF on device: %v", err2)
    }
    err3 := ospf_lib.Unconfigure_ospf_timers("10.133.35.143","Ethernet 1/11")
    if err3 != nil {
        t.Errorf("Failed to unconfigure OSPF timers on device: %v", err3)
    }
}
func Test_ospf_passive_interface(t *testing.T) {
    fmt.Printf("Configure ospf passive interface on device\n")
    err := ospf_lib.Configure_ospf_passive_interface("10.133.35.148","Ethernet 1/11")
    if err != nil {
        t.Errorf("Failed to configure OSPF passive interface on device: %v", err)
    }
    ospf_lib.Print_Output("10.133.35.148","show ip ospf interface ethernet 1/11")
    err1 := ospf_lib.Unconfigure_ospf_passive_interface("10.133.35.148","Ethernet 1/11")
    if err1 != nil {
        t.Errorf("Failed to unconfigure OSPF passive interface on device: %v", err1)
    }
} 

func Test_Ospf_DR_BDR(t *testing.T){
    fmt.Printf("Configure ospf priority on routers\n")
    err := ospf_lib.Configure_ospf_DR_BDR("10.133.35.148","Ethernet 1/11")
    if err != nil {
        t.Errorf("Failed to configure OSPF priority on device: %v", err)
    }
    err1 := ospf_lib.Configure_ospf_clear_process("10.133.35.150")
    if err1 != nil {
        t.Errorf("Failed to configure OSPF priority on device: %v", err1)
    }
    err2 := ospf_lib.Configure_ospf_clear_process("10.133.35.143")
    if err2 != nil {
        t.Errorf("Failed to configure OSPF priority on device: %v", err2)
    }
    ospf_lib.Print_Output("10.133.35.143","show ip ospf route")
    err3 := ospf_lib.Unconfigure_ospf_DR_BDR("10.133.35.148","Ethernet 1/11")
    if err3 != nil {
        t.Errorf("Failed to unconfigure OSPF priority on device: %v", err3)
    }
    err4 := ospf_lib.Configure_ospf_clear_process("10.133.35.150")
    if err4 != nil {
        t.Errorf("Failed to configure OSPF priority on device: %v", err4)
    }
    err5 := ospf_lib.Configure_ospf_clear_process("10.133.35.143")
    if err5 != nil {
        t.Errorf("Failed to configure OSPF priority on device: %v", err5)
    }
}
func Test_Ospf_Graceful_shutdown(t *testing.T){
    fmt.Printf("Configure ospf graceful shutdown\n")
    err := ospf_lib.Configure_ospf_graceful_shutdown("10.133.35.143","2")
    if err != nil {
        t.Errorf("Failed to configure OSPF shutdown on device: %v", err)
    }
    ospf_lib.Print_Output("10.133.35.143","show ip ospf neighbors") 
    err1 := ospf_lib.Unconfigure_ospf_graceful_shutdown("10.133.35.143","2")
    if err1 != nil {
        t.Errorf("Failed to configure OSPF shutdown on device: %v", err1)
    }
}
func Test_Ospf_Virtual_Links(t *testing.T){
    fmt.Printf("Configure ospf area on devices\n")
    err := ospf_lib.Configure_ospf_multiarea("10.133.35.148","Ethernet 1/11","2")
    if err != nil {
        t.Errorf("Failed to configure OSPF multiarea on device: %v", err)
    }
    err1 := ospf_lib.Configure_ospf_multiarea("10.133.35.150","Ethernet 1/1","2")
    if err1 != nil {
        t.Errorf("Failed to configure OSPF multiarea on device: %v", err1)
    }
    err2 := ospf_lib.Configure_ospf_loopback("10.133.35.143","Ethernet 1/7","2")
    if err2 != nil {
        t.Errorf("Failed to configure OSPF loopback on device: %v", err2)
    }
    err3 := ospf_lib.Configure_ospf_loopback("10.133.35.143","Ethernet 1/11","2")
    if err3 != nil {
        t.Errorf("Failed to configure OSPF loopback on device: %v", err3)
    }
    err4 := ospf_lib.Configure_ospf_interface("10.133.35.143","Ethernet 1/7","2")
    if err4 != nil {
        t.Errorf("Failed to configure OSPF interface on device: %v", err4)
    }
    err5 := ospf_lib.Configure_ospf_vlinks("10.133.35.150","2")
    if err5 != nil {
        t.Errorf("Failed to configure OSPF vlinks on device: %v", err5)
    }
    err6 := ospf_lib.Configure_ospf_vlinks("10.133.35.143","2")
    if err6 != nil {
        t.Errorf("Failed to configure OSPF vlinks on device: %v", err6)
    }
    ospf_lib.Print_Output("10.133.35.143","show ip ospf virtual-links")
    err7 := ospf_lib.Unconfigure_ospf_vlinks("10.133.35.150","2")
    if err7 != nil {
        t.Errorf("Failed to unconfigure OSPF vlinks on device: %v", err7)
    }
    err8 := ospf_lib.Unconfigure_ospf_vlinks("10.133.35.143","2")
    if err8 != nil {
        t.Errorf("Failed to configure OSPF vlinks on device: %v", err8)
    }
    err9 := ospf_lib.Unconfigure_ospf_multiarea("10.133.35.148","Ethernet 1/11","2")
    if err9 != nil {
        t.Errorf("Failed to configure OSPF multiarea on device: %v", err9)
    }
    err10 := ospf_lib.Unconfigure_ospf_multiarea("10.133.35.150","Ethernet 1/1","2")
    if err10 != nil {
        t.Errorf("Failed to configure OSPF multiarea on device: %v", err10)
    }
    err11 := ospf_lib.Unconfigure_ospf_loopback("10.133.35.143","Ethernet 1/7","2")
    if err11 != nil {
        t.Errorf("Failed to configure OSPF loopback on device: %v", err11)
    }
    err12 := ospf_lib.Unconfigure_ospf_loopback("10.133.35.143","Ethernet 1/11","2")
    if err12 != nil {
        t.Errorf("Failed to configure OSPF loopback on device: %v", err12)
    }
    err13 := ospf_lib.Unconfigure_ospf_interface("10.133.35.143","Ethernet 1/7","2")
    if err13 != nil {
        t.Errorf("Failed to configure OSPF interface on device: %v", err13)
    }
}*/
