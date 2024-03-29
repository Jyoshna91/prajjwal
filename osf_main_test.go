package ospf
import "ospf/ospf_lib"

import (
    "fmt"
//    "golang.org/x/crypto/ssh"
//    "log"
    "testing"
    "time"
    "github.com/tealeg/xlsx"
//    "strings"
//    "bytes"
//    "bufio"
      "os"
)
var testResults = make(map[string]string)

func Test_Configure_ospf_routers(t *testing.T){
    testCaseName := "Configure ospf routers"
    
    fmt.Printf("Enabling ospf on routers\n")
    err := ospf_lib.Enable_ospf("10.133.35.148")
     if err != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName, err)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"

    err1 := ospf_lib.Enable_ospf("10.133.35.143")
    if err1 != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName, err1)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"

    err2 := ospf_lib.Enable_ospf("10.133.35.150")
    if err2 != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName, err2)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"

    fmt.Printf("Configuring ospf on routers\n")
    err3 := ospf_lib.Configure_ospf_router("10.133.35.148","2")
    if err3 != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName, err3)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"

    err4 := ospf_lib.Configure_ospf_router("10.133.35.143","2")
    if err4 != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName, err4)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"

    err5 := ospf_lib.Configure_ospf_router("10.133.35.150","2")
    if err5 != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName, err5)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
    
    fmt.Printf("Configuring ospf on interface\n") //Configure_ospf_interface
   
    err6 := ospf_lib.Configure_ospf_interface("10.133.35.148","Ethernet 1/11","2")
    if err6 != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName, err6)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"

    err7 := ospf_lib.Configure_ospf_interface("10.133.35.148","Ethernet 1/1","2")
    if err7 != nil {
         t.Errorf("Test '%s' failed: %v",testCaseName, err7)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"

    err8 := ospf_lib.Configure_ospf_interface("10.133.35.150","Ethernet 1/1","2")
    if err8 != nil {
         t.Errorf("Test '%s' failed: %v",testCaseName, err8)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"

    err9 := ospf_lib.Configure_ospf_interface("10.133.35.143","Ethernet 1/11","2")
    if err9 != nil {
         t.Errorf("Test '%s' failed: %v",testCaseName, err9)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"

    for i:=0;i<2;i++ {
   	 ospf_lib.Print_Output("10.133.35.148","show ip ospf neighbors")
    }
    ospf_lib.Validation("10.133.35.148","show ip ospf neighbors","FULL")
    err10 := ospf_lib.Unconfigure_ospf_interface("10.133.35.148","Ethernet 1/11","2")
    if err10 != nil {
         t.Errorf("Test '%s' failed: %v",testCaseName, err10)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"

    err11 := ospf_lib.Unconfigure_ospf_interface("10.133.35.148","Ethernet 1/1","2")
    if err11 != nil {
         t.Errorf("Test '%s' failed: %v",testCaseName, err11)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"

    err12 := ospf_lib.Unconfigure_ospf_interface("10.133.35.150","Ethernet 1/1","2")
    if err12 != nil {
         t.Errorf("Test '%s' failed: %v",testCaseName, err12)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"

    err13 := ospf_lib.Unconfigure_ospf_interface("10.133.35.143","Ethernet 1/11","2")
    if err13 != nil {
         t.Errorf("Test '%s' failed: %v",testCaseName, err13)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"

    err14 := ospf_lib.Unconfigure_ospf_router("10.133.35.148","2")
    if err14 != nil {
         t.Errorf("Test '%s' failed: %v",testCaseName, err14)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"

    err15 := ospf_lib.Unconfigure_ospf_router("10.133.35.143","2")
    if err15 != nil {
         t.Errorf("Test '%s' failed: %v",testCaseName, err15)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"

    err16 := ospf_lib.Unconfigure_ospf_router("10.133.35.150","2")
    if err16 != nil {
         t.Errorf("Test '%s' failed: %v",testCaseName, err16)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
}

/*
func Test_Multiarea(t *testing.T) {
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
    err8 := ospf_lib.Configure_ospf_multiarea("10.133.35.143","Ethernet 1/7","2")
    if err8 != nil {
        t.Errorf("Failed to configure OSPF multiarea on device: %v", err8)
    }
    err9 := ospf_lib.Configure_ospf_loopback("10.133.35.150","loopback 0","2")
    if err9 != nil {
        t.Errorf("Failed to configure OSPF loopback on device: %v", err9)
    }
    ospf_lib.Print_Output("10.133.35.150","show ip ospf route")
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
    ospf_lib.Validation("10.133.35.148","show ip ospf 2 interface eth1/11","Message-digest authentication")
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
    ospf_lib.Validation("10.133.35.148","show ip ospf interface ethernet 1/11","Timer intervals: Hello 30, Dead 100")
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
    ospf_lib.Validation("10.133.35.148","show ip ospf interface ethernet 1/11","Passive interface")
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
    ospf_lib.Validation("10.133.35.143","show ip ospf route","intra")
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
    ospf_lib.Validation("10.133.35.143","show ip ospf neighbors","")
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
    ospf_lib.Print_Output("10.133.35.150","show ip ospf virtual-links")
    ospf_lib.Validation("10.133.35.150","show ip ospf virtual-links","Virtual link VL1")
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
}
func Test_ospf_stub_totally(t *testing.T){
    fmt.Printf("Configure ospf areas on device1\n")
    err := ospf_lib.Configure_ospf_multiarea("10.133.35.148","Ethernet 1/11","2")
    if err != nil {
        t.Errorf("Failed to configure OSPF multiarea on device: %v", err)
    }
    err1 := ospf_lib.Configure_ospf_interface("10.133.35.148","Ethernet 1/1","2")
    if err1 != nil {
        t.Errorf("Failed to configure OSPF interface on device: %v", err1)
    }
    fmt.Printf("Configure ospf areas on device2\n")
    err2 := ospf_lib.Configure_ospf_multiarea("10.133.35.150","Ethernet 1/11","2")
    if err2 != nil {
        t.Errorf("Failed to configure OSPF multiarea on device: %v", err2)
    }
    err3 := ospf_lib.Configure_ospf_interface("10.133.35.150","Ethernet 1/7","2")
    if err3 != nil {
        t.Errorf("Failed to configure OSPF interface on device: %v", err3)
    }
    fmt.Printf("Configure ospf areas on device3\n")
    err4 := ospf_lib.Configure_ospf_interface("10.133.35.143","Ethernet 1/11","2")
    if err4 != nil {
        t.Errorf("Failed to configure OSPF multiarea on device: %v", err4)
    }
    err5 := ospf_lib.Configure_ospf_interface("10.133.35.143","Ethernet 1/7","2")
    if err5 != nil {
        t.Errorf("Failed to configure OSPF interface on device: %v", err5)
    }
    err6 := ospf_lib.Configure_ospf_interface("10.133.35.143","loopback0","2")
    if err6 != nil {
        t.Errorf("Failed to configure OSPF interface on device: %v", err6)
    }
    err7 := ospf_lib.Configure_ospf_stub_totally("10.133.35.143","2")
    if err7 != nil {
        t.Errorf("Failed to configure OSPF stub totally on device: %v", err7)
    }
    err8 := ospf_lib.Configure_ospf_stub_totally("10.133.35.150","2")
    if err8 != nil {
        t.Errorf("Failed to configure OSPF stub totally on device: %v", err8)
    }
    ospf_lib.Print_Output("10.133.35.150","show ip ospf route")
    ospf_lib.Validation("10.133.35.150","show ip ospf route","intra")
    err9 := ospf_lib.Unconfigure_ospf_multiarea("10.133.35.148","Ethernet 1/11","2")
    if err9 != nil {
        t.Errorf("Failed to configure OSPF multiarea on device: %v", err9)
    }
    err10 := ospf_lib.Unconfigure_ospf_interface("10.133.35.148","Ethernet 1/1","2")
    if err10 != nil {
        t.Errorf("Failed to configure OSPF interface on device: %v", err10)
    }
    //fmt.Printf("UnConfigure ospf areas on device2\n")
    err11 := ospf_lib.Unconfigure_ospf_multiarea("10.133.35.150","Ethernet 1/11","2")
    if err11 != nil {
        t.Errorf("Failed to configure OSPF multiarea on device: %v", err11)
    }
    err12 := ospf_lib.Unconfigure_ospf_interface("10.133.35.150","Ethernet 1/7","2")
    if err12 != nil {
        t.Errorf("Failed to configure OSPF interface on device: %v", err12)
    }
    //fmt.Printf("Configure ospf areas on device3\n")
    err13 := ospf_lib.Unconfigure_ospf_interface("10.133.35.143","Ethernet 1/11","2")
    if err13 != nil {
        t.Errorf("Failed to configure OSPF multiarea on device: %v", err13)
    }
    err14 := ospf_lib.Unconfigure_ospf_interface("10.133.35.143","Ethernet 1/7","2")
    if err14 != nil {
        t.Errorf("Failed to configure OSPF interface on device: %v", err14)
    }
    err15 := ospf_lib.Unconfigure_ospf_interface("10.133.35.143","loopback0","2")
    if err15 != nil {
        t.Errorf("Failed to configure OSPF interface on device: %v", err15)
    }
    err16 := ospf_lib.Unconfigure_ospf_stub_totally("10.133.35.143","2")
    if err16 != nil {
        t.Errorf("Failed to configure OSPF stub totally on device: %v", err16)
    }
    err17 := ospf_lib.Unconfigure_ospf_stub_totally("10.133.35.150","2")
    if err17 != nil {
        t.Errorf("Failed to configure OSPF stub totally on device: %v", err17)
    }
}*/

func TestMain(m *testing.M) {
exitVal := m.Run()

// Prepare to summarize test results
currentTime := time.Now()
passedCount, failedCount := 0, 0

// Tally passed and failed tests
for _, result := range testResults {
switch result {
case "PASSED":
passedCount++
case "FAILED":
failedCount++
}
}
totalTests := passedCount + failedCount
successRate := float64(passedCount) / float64(totalTests) * 100.0

// Print the summary header
fmt.Println("--------------------------------------------------------------------------------------")
fmt.Printf("%s INFO |                     Task Result Summary                   |\n", currentTime.Format("2006-01-02 15:04:05"))
fmt.Println("--------------------------------------------------------------------------------------")

// Print counts for passed, failed, and total tests with headings
fmt.Printf("%s INFO: %-30s %8s\n", currentTime.Format("2006-01-02 15:04:05"), "Metric", "Count")
fmt.Println("--------------------------------------------------------------------------------------")
fmt.Printf("%s INFO: %-30s %8d\n", currentTime.Format("2006-01-02 15:04:05"), "Passed", passedCount)
fmt.Printf("%s INFO: %-30s %8d\n", currentTime.Format("2006-01-02 15:04:05"), "Failed", failedCount)
fmt.Printf("%s INFO: %-30s %8d\n", currentTime.Format("2006-01-02 15:04:05"), "TOTAL", totalTests)
fmt.Printf("%s INFO: %-30s %.2f%%\n", currentTime.Format("2006-01-02 15:04:05"), "Success Rate", successRate)
fmt.Println("--------------------------------------------------------------------------------------")

// Print the task summary
fmt.Printf("%s INFO |                         Task Summary                      |\n", currentTime.Format("2006-01-02 15:04:05"))
fmt.Println("--------------------------------------------------------------------------------------")
counter := 1
for testName, result := range testResults {
fmt.Printf("%s INFO: %-20s %-40s %s\n", currentTime.Format("2006-01-02 15:04:05"), fmt.Sprintf("Testcase%d", counter), testName+":", result)
counter++
}
fmt.Println("--------------------------------------------------------------------------------------")

// Write the results to an Excel file
file := xlsx.NewFile()
sheet, err := file.AddSheet("Test Results")
if err != nil {
fmt.Printf("Failed to add sheet: %v\n", err)
os.Exit(1)
}

// Define cell styles for passed and failed tests
passedStyle := xlsx.NewStyle()
passedStyle.Fill = *xlsx.NewFill("solid", "00FF00", "00FF00")
passedStyle.ApplyFill = true
failedStyle := xlsx.NewStyle()
failedStyle.Fill = *xlsx.NewFill("solid", "FF0000", "FF0000")
failedStyle.ApplyFill = true

// Write test result headers
row := sheet.AddRow()
row.AddCell().SetValue("Metric")
row.AddCell().SetValue("Count")

// Write test result counts
row = sheet.AddRow()
row.AddCell().SetValue("Passed")
passedCell := row.AddCell()
passedCell.SetValue(passedCount)
passedCell.SetStyle(passedStyle)

row = sheet.AddRow()
row.AddCell().SetValue("Failed")
failedCell := row.AddCell()
failedCell.SetValue(failedCount)
failedCell.SetStyle(failedStyle)

row = sheet.AddRow()
row.AddCell().SetValue("TOTAL")
row.AddCell().SetValue(totalTests)

row = sheet.AddRow()
row.AddCell().SetValue("Success Rate")
row.AddCell().SetValue(fmt.Sprintf("%.2f%%", successRate))

// Add a blank row for spacing
sheet.AddRow()

// Write individual test results
counter = 1
for testName, result := range testResults {
row = sheet.AddRow()
row.AddCell().SetValue(fmt.Sprintf("Testcase%d", counter))
row.AddCell().SetValue(testName)
cell := row.AddCell()
cell.SetValue(result)
if result == "PASSED" {
cell.SetStyle(passedStyle)
} else if result == "FAILED" {
cell.SetStyle(failedStyle)
}
counter++
}

// Specify the full paths where you want to save the Excel file and the testcase file
excelFilePath := "/home/tcs/sample/ondatra/debug/ospf/test_summary.xlsx"
testcaseFilePath := "/home/tcs/sample/ondatra/debug/ospf/newcode_test.go"

// Save the Excel file at the specified path
if err := file.Save(excelFilePath); err != nil {
fmt.Printf("Failed to save Excel file at %s: %v\n", excelFilePath, err)
os.Exit(1)
}

// Write the paths for the testcase file and the Excel sheet file in the Excel sheet
//sheet.SetCellHyperLink(1, 3, testcaseFilePath, "external")
//sheet.SetCellHyperLink(2, 3, excelFilePath, "external")

// Print the paths for the testcase file and the Excel sheet file
fmt.Printf("Path for testcase file: %s\n", testcaseFilePath)
fmt.Printf("Path for Excel: %s\n", excelFilePath)

// Exit with the appropriate status code
os.Exit(exitVal)
}

