package rip
import "rip/rip_lib"
import (
    "fmt"
    "github.com/tealeg/xlsx"
    "golang.org/x/crypto/ssh"
    "log"
    "testing"
    "time"
    "os"
)
var testResults = make(map[string]string)
func EnableRiponDevices(deviceIP string, interf string, intf_ip string, ripName string) error {
    serverAddress := deviceIP
    serverUsername := "admin"
    serverPassword := "tcs123"
    intf := interf
    rip_name := ripName

    fmt.Printf("\nTestcase Started: Enable Rip -----\n")

    config := &ssh.ClientConfig{
        User: serverUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(serverPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    client, err := ssh.Dial("tcp", serverAddress+":22", config)
    if err != nil {
        panic(err)
    }
    defer client.Close()

    // configure
    rip_lib.Configure_interface_ip(serverAddress, intf, intf_ip)
    rip_lib.Enable_rip(serverAddress)
    rip_lib.Configure_rip_on_device(serverAddress, rip_name)
    rip_lib.Configure_rip_on_interface(serverAddress, intf, rip_name)
    
    fmt.Printf("RIP enabled on device %s %s\n", serverAddress,interf)

    fmt.Printf("Testcase Ended -----\n\n")
    return nil
}

func TestEnableRip(t *testing.T) {
    testCaseName := "Test Enable Rip"
//    err2 := EnableRiponDevices("10.133.35.148","Ethernet1/11", "192.168.3.1/24", "riptemp")
    err3 := EnableRiponDevices("10.133.35.143","Ethernet1/11", "192.168.3.1/24", "riptemp")
    
//    if err2 != nil {
//        t.Errorf("Failed to enable RIP on device: %v", err2)
//    }
    if err3 != nil {
        t.Errorf("Failed to enable RIP on device: %v", err3)
        testResults[testCaseName] = "FAILED"
    } else {
       testResults[testCaseName] = "PASSED"
    }
}



func RipPassive(deviceIP string,interf string, interf_ip string, rip_name string) error {
    serverAddress := deviceIP
    serverUsername := "admin"
    serverPassword := "tcs123"
    OutputCommand := "show ip rip interface "+string(interf)

    fmt.Printf("\nTestcase Started: Rip Passive -----\n")

    config := &ssh.ClientConfig{
        User: serverUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(serverPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    client, err := ssh.Dial("tcp", serverAddress+":22", config)
    if err != nil {
        panic(err)
    }
    defer client.Close()

    session, err := client.NewSession()
    if err !=nil {
        panic(err)
    }
    defer session.Close()
    modes := ssh.TerminalModes{
            ssh.ECHO: 0,
            ssh.TTY_OP_ISPEED: 14400,
            ssh.TTY_OP_OSPEED: 14400,
    }

    if err := session.RequestPty("vt100", 0, 0, modes); err != nil {
        panic(err)
    }

    sshShell, err := session.StdinPipe()
    if err != nil {
        panic(err)
    }
    if err :=session.Shell(); err !=nil {
        panic(err)
    }

    //enabling
    EnableRiponDevices(deviceIP, interf, interf_ip, rip_name)
    fmt.Printf("RIP %s enabled on device %s %s\n", rip_name, serverAddress,interf)

    commands := []string{
            "conf t\n",
            fmt.Sprintf("interface %s\n",interf),
            fmt.Sprintf("ip rip passive-interface\n"),
            fmt.Sprintf("no shutdown\n"),
	    "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }

    // New Session
    router1Config := &ssh.ClientConfig{
        User: serverUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(serverPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    router1Client, err := ssh.Dial("tcp", serverAddress+":22", router1Config)
    if err !=nil {
        log.Fatalf("Failed to connect to Router 1: %v", err)
    }
    defer router1Client.Close()

    router1Session, err := router1Client.NewSession()
    if err != nil {
        log.Fatalf("Failed to create session on Router 1: %v", err)

    }
    defer router1Session.Close()

    router1Output, err := router1Session.CombinedOutput(OutputCommand)
    if err != nil {
        fmt.Printf("Failed to execute ping command on Router 1: %v", err)
        return fmt.Errorf("failed to execute command on Router 1: %w",err)
    }
    fmt.Printf("Router %s Output:\n%s\n", serverAddress, router1Output)
    rip_lib.Validation("passive",string(router1Output))
    fmt.Printf("Testcase Ended -----\n\n")

    return nil
}

func TestRipPassive(t *testing.T){
    testCaseName := "TestRipPassive"
    err := RipPassive("10.133.35.143","Ethernet1/11", "192.168.3.2/24", "riptemp")
    if err != nil {
        t.Errorf("Failed to enable RIP on device: %v", err)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
    }
}



// Testcase RipTimers -------------
func RipTimers(deviceIP string, interf string, rip_name string) error {
    serverAddress := deviceIP
    serverUsername := "admin"
    serverPassword := "tcs123"
    OutputCommand := "show ip rip"

    fmt.Printf("\nTestcase Started: Rip Timers -----\n")

    config := &ssh.ClientConfig{
        User: serverUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(serverPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    client, err := ssh.Dial("tcp", serverAddress+":22", config)
    if err != nil {
        panic(err)
    }
    defer client.Close()

    session, err := client.NewSession()
    if err !=nil {
        panic(err)
    }
    defer session.Close()

    modes := ssh.TerminalModes{
            ssh.ECHO: 0,
            ssh.TTY_OP_ISPEED: 14400,
            ssh.TTY_OP_OSPEED: 14400,
    }

    if err := session.RequestPty("vt100", 0, 0, modes); err != nil {
        panic(err)
    }

    sshShell, err := session.StdinPipe()
    if err != nil {
        panic(err)
    }
    if err :=session.Shell(); err !=nil {
        panic(err)
    }

    commands := []string{
            "conf t\n",
            fmt.Sprintf("router rip %s\n", rip_name),
            fmt.Sprintf("address-family ipv4 unicast\n"),
            fmt.Sprintf("timers basic 10 60 60 40\n"),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }

    router1Config := &ssh.ClientConfig{
        User: serverUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(serverPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    router1Client, err := ssh.Dial("tcp", serverAddress+":22", router1Config)
    if err !=nil {
        log.Fatalf("Failed to connect to Router 1: %v", err)
    }
    defer router1Client.Close()

    router1Session, err := router1Client.NewSession()
    if err != nil {
        log.Fatalf("Failed to create session on Router 1: %v", err)

    }
    defer router1Session.Close()

    router1Output, err := router1Session.CombinedOutput(OutputCommand)
    if err != nil {
        fmt.Printf("Failed to execute ping command on Router 1: %v", err)
        return fmt.Errorf("failed to execute command on Router 1: %w",err)
    }
    fmt.Printf("Router %s Output:\n%s\n", serverAddress, router1Output)
    rip_lib.Validation("Updates every 10 sec, expire in 60 sec",string(router1Output))
    fmt.Printf("Testcase Ended -----\n\n")
    return nil
    
}

func TestRipTimers(t *testing.T){
    testCaseName :=  "Test Rip Timers"
    err := RipTimers("10.133.35.143","Ethernet1/11", "riptemp")
    if err != nil {
        t.Errorf("Failed to enable RipTimers on device: %v", err)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
    }
}



func RipDistance(deviceIP string, interf string, rip_name string) error {
    serverAddress := deviceIP
    serverUsername := "admin"
    serverPassword := "tcs123"
    OutputCommand := "show ip rip"

    fmt.Printf("\nTestcase Started: Rip Distance -----\n")

    config := &ssh.ClientConfig{
        User: serverUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(serverPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    client, err := ssh.Dial("tcp", serverAddress+":22", config)
    if err != nil {
        panic(err)
    }
    defer client.Close()

    session, err := client.NewSession()
    if err !=nil {
        panic(err)
    }
    defer session.Close()

    modes := ssh.TerminalModes{
            ssh.ECHO: 0,
            ssh.TTY_OP_ISPEED: 14400,
            ssh.TTY_OP_OSPEED: 14400,
    }

    if err := session.RequestPty("vt100", 0, 0, modes); err != nil {
        panic(err)
    }

    sshShell, err := session.StdinPipe()
    if err != nil {
        panic(err)
    }
    if err :=session.Shell(); err !=nil {
        panic(err)
    }

    commands := []string{
            "conf t\n",
            fmt.Sprintf("router rip %s\n", rip_name),
            fmt.Sprintf("address-family ipv4 unicast\n"),
            fmt.Sprintf("distance 100\n"),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }

    router1Config := &ssh.ClientConfig{
        User: serverUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(serverPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    router1Client, err := ssh.Dial("tcp", serverAddress+":22", router1Config)
    if err !=nil {
        log.Fatalf("Failed to connect to Router 1: %v", err)
    }
    defer router1Client.Close()

    router1Session, err := router1Client.NewSession()
    if err != nil {
        log.Fatalf("Failed to create session on Router 1: %v", err)

    }
    defer router1Session.Close()

    router1Output, err := router1Session.CombinedOutput(OutputCommand)
    if err != nil {
        fmt.Printf("Failed to execute ping command on Router 1: %v", err)
        return fmt.Errorf("failed to execute command on Router 1: %w",err)
    }
    fmt.Printf("Router %s Output:\n%s\n", serverAddress, router1Output)
    rip_lib.Validation("Admin-distance: 100",string(router1Output))
    fmt.Printf("Testcase Ended -----\n\n")
    return nil
}

func TestRipDistance(t *testing.T){
    testCaseName := "Test Rip Distance"
    err := RipDistance("10.133.35.143","Ethernet1/11", "riptemp")
    if err != nil {
        t.Errorf("Failed to enable RipDistance on device: %v", err)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
    }
}


func RipAuthentication(deviceIP string, interf string)error {
    serverAddress := deviceIP
    serverUsername := "admin"
    serverPassword := "tcs123"
    OutputCommand := "show ip rip interface "+string(interf)

    fmt.Printf("\nTestcase Started: Rip Authentication -----\n")

    config := &ssh.ClientConfig{
        User: serverUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(serverPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    client, err := ssh.Dial("tcp", serverAddress+":22", config)
    if err != nil {
        panic(err)
    }
    defer client.Close()

    session, err := client.NewSession()
    if err !=nil {
        panic(err)
    }
    defer session.Close()

    modes := ssh.TerminalModes{
            ssh.ECHO: 0,
            ssh.TTY_OP_ISPEED: 14400,
            ssh.TTY_OP_OSPEED: 14400,
    }
    if err := session.RequestPty("vt100", 0, 0, modes); err != nil {
        panic(err)
        panic(err)
    }

    sshShell, err := session.StdinPipe()
    if err != nil {
        panic(err)
    }

    if err :=session.Shell(); err !=nil {
        panic(err)
    }

//    EnableRiponDevices(deviceIP,interf)
    commands := []string{
            "conf t\n",
            fmt.Sprintf("key chain rip\n"),
            fmt.Sprintf("key 1\n"),
            fmt.Sprintf("key-string CISCO\n"),
            fmt.Sprintf("exit\n"),
            fmt.Sprintf("int %s\n",interf),
            fmt.Sprintf("ip rip authentication key-chain rip\n"),
            fmt.Sprintf("ip rip authentication mode md5\n"),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    fmt.Printf("Rip authentication successfully verified on %s via %s\n",deviceIP,interf)

    router1Config := &ssh.ClientConfig{
        User: serverUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(serverPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    router1Client, err := ssh.Dial("tcp", serverAddress+":22", router1Config)
    if err !=nil {
        log.Fatalf("Failed to connect to Router 1: %v", err)
    }
    defer router1Client.Close()

    router1Session, err := router1Client.NewSession()
    if err != nil {
        log.Fatalf("Failed to create session on Router 1: %v", err)

    }
    defer router1Session.Close()

    router1Output, err := router1Session.CombinedOutput(OutputCommand)
    if err != nil {
        fmt.Printf("Failed to execute ping command on Router 1: %v", err)
        return fmt.Errorf("failed to execute command on Router 1: %w",err)
    }
    fmt.Printf("Router %s Output:\n%s\n", serverAddress, router1Output)
    rip_lib.Validation("Authentication Mode: md5  Keychain: rip",string(router1Output))
    fmt.Printf("Testcase Ended -----\n\n")
    return nil
}

func TestRipAuthentication(t *testing.T){
    testCaseName := "Test Rip Authentication"
    err := RipAuthentication("10.133.35.143","Ethernet1/11")
    if err != nil {
        t.Errorf("Failed to enable RipAuthentication on device: %v", err)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
    }
}


func RipMaxPaths(deviceIP string, interf string)error {
    serverAddress := deviceIP
    serverUsername := "admin"
    serverPassword := "tcs123"
    OutputCommand := "show ip rip"

    fmt.Printf("\nTestcase Started: RIP Max Paths -----\n")

    config := &ssh.ClientConfig{
        User: serverUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(serverPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    client, err := ssh.Dial("tcp", serverAddress+":22", config)
    if err != nil {
        panic(err)
    }
    defer client.Close()

    session, err := client.NewSession()
    if err !=nil {
        panic(err)
    }
    defer session.Close()

    modes := ssh.TerminalModes{
            ssh.ECHO: 0,
            ssh.TTY_OP_ISPEED: 14400,
            ssh.TTY_OP_OSPEED: 14400,
    }
     if err := session.RequestPty("vt100", 0, 0, modes); err != nil {
        panic(err)
        panic(err)
    }

    sshShell, err := session.StdinPipe()
    if err != nil {
        panic(err)
    }

    if err :=session.Shell(); err !=nil {
        panic(err)
    }

    commands := []string{
            "conf t\n",
            fmt.Sprintf("router rip 2\n"),
            fmt.Sprintf("address-family ipv4 unicast\n"),
            fmt.Sprintf("maximum-paths 5\n"),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    fmt.Printf("Rip authentication successfully verified on %s via %s\n",deviceIP,interf)

    router1Config := &ssh.ClientConfig{
        User: serverUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(serverPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    router1Client, err := ssh.Dial("tcp", serverAddress+":22", router1Config)
    if err !=nil {
        log.Fatalf("Failed to connect to Router 1: %v", err)
    }
    defer router1Client.Close()

    router1Session, err := router1Client.NewSession()
    if err != nil {
        log.Fatalf("Failed to create session on Router 1: %v", err)

    }
    defer router1Session.Close()

    router1Output, err := router1Session.CombinedOutput(OutputCommand)
    if err != nil {
        fmt.Printf("Failed to execute ping command on Router 1: %v", err)
        return fmt.Errorf("failed to execute command on Router 1: %w",err)
    }
    fmt.Printf("Router %s Output:\n%s\n", serverAddress, router1Output)
    rip_lib.Validation("Max-paths: 5",string(router1Output))
    fmt.Printf("Testcase Ended -----\n\n")
    return nil
}

func TestRipMaxPaths(t *testing.T){
    testCaseName := "Test Rip Max Paths"
    err := RipMaxPaths("10.133.35.143","Ethernet1/11")
    if err != nil {
        t.Errorf("Failed to enable RipMaxPaths on device: %v", err)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
}
}


func RipRedistribution(deviceIP string, interf string, rip_name string) error {
    serverAddress := deviceIP
    serverUsername := "admin"
    serverPassword := "tcs123"
    OutputCommand := "show ip rip"

    fmt.Printf("\nTestcase Started: Rip Redistribution -----\n")

    config := &ssh.ClientConfig{
        User: serverUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(serverPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    client, err := ssh.Dial("tcp", serverAddress+":22", config)
    if err != nil {
        panic(err)
    }
    defer client.Close()

    session, err := client.NewSession()
    if err !=nil {
        panic(err)
    }
    defer session.Close()

    modes := ssh.TerminalModes{
            ssh.ECHO: 0,
            ssh.TTY_OP_ISPEED: 14400,
            ssh.TTY_OP_OSPEED: 14400,
    }
     if err := session.RequestPty("vt100", 0, 0, modes); err != nil {
        panic(err)
        panic(err)
    }

    sshShell, err := session.StdinPipe()
    if err != nil {
        panic(err)
    }
    if err :=session.Shell(); err !=nil {
        panic(err)
    }

    commands := []string{
            "conf t\n",
            fmt.Sprintf("router rip %s\n", rip_name),
            fmt.Sprintf("address-family ipv4 unicast\n"),
            fmt.Sprintf("redistribute ospf 1 route-map rip\n"),
            fmt.Sprintf("route-map rip permit\n"),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    fmt.Printf("RipRedistribution successfully verified on %s via %s\n",deviceIP,interf)

    // New Session
    router1Config := &ssh.ClientConfig{
        User: serverUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(serverPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    router1Client, err := ssh.Dial("tcp", serverAddress+":22", router1Config)
    if err !=nil {
        log.Fatalf("Failed to connect to Router 1: %v", err)
    }
    defer router1Client.Close()

    router1Session, err := router1Client.NewSession()
    if err != nil {
        log.Fatalf("Failed to create session on Router 1: %v", err)

    }
    defer router1Session.Close()

    router1Output, err := router1Session.CombinedOutput(OutputCommand)
    if err != nil {
        fmt.Printf("Failed to execute ping command on Router 1: %v", err)
        return fmt.Errorf("failed to execute command on Router 1: %w",err)
    }
    fmt.Printf("Router %s Output:\n%s\n", serverAddress, router1Output)
    rip_lib.Validation("ospf-1          policy rip",string(router1Output))
    fmt.Printf("Testcase Ended -----\n\n")
    return nil
}

func TestRipRedistribution(t *testing.T) {
    testCaseName := "Rip Redistribution"
    err := RipRedistribution("10.133.35.143","Ethernet1/11", "riptemp")
    if err != nil {
        t.Errorf("Failed to enable RipRedistribution on device: %v", err)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"

    }

}



// ----------- CLEANUP SECTION -----------
// Function to clean RIP Passive testcase
func CleaningRIPpassive(deviceIP string, interf string) error {
    serverAddress := deviceIP
    serverUsername := "admin"
    serverPassword := "tcs123"
    config := &ssh.ClientConfig{
        User: serverUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(serverPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    client, err := ssh.Dial("tcp", serverAddress+":22", config)
    if err != nil {
        panic(err)
    }
    defer client.Close()
    session, err := client.NewSession()
    if err !=nil {
        panic(err)
    }
    defer session.Close()
    modes := ssh.TerminalModes{
            ssh.ECHO: 0,
            ssh.TTY_OP_ISPEED: 14400,
            ssh.TTY_OP_OSPEED: 14400,
    }
    if err := session.RequestPty("vt100", 0, 0, modes); err != nil {
        panic(err)
    }
    sshShell, err := session.StdinPipe()
    if err != nil {
        panic(err)
    }
    if err :=session.Shell(); err !=nil {
        panic(err)
    }
    commands := []string{
            "conf t\n",
            fmt.Sprintf("int %s\n",interf),
            fmt.Sprintf("no ip rip passive-interface\n"),
            "end\n",
    }
    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil
}

// Function to clean RIP Timers, Distance, Max Paths, Redistribution testcases
func CleaningRIPDistanceTimers(deviceIP string, interf string, rip_name string) error {
    serverAddress := deviceIP
    serverUsername := "admin"
    serverPassword := "tcs123"
    config := &ssh.ClientConfig{
        User: serverUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(serverPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    client, err := ssh.Dial("tcp", serverAddress+":22", config)
    if err != nil {
        panic(err)
    }
    defer client.Close()
    session, err := client.NewSession()
    if err !=nil {
        panic(err)
    }
    defer session.Close()
    modes := ssh.TerminalModes{
            ssh.ECHO: 0,
            ssh.TTY_OP_ISPEED: 14400,
            ssh.TTY_OP_OSPEED: 14400,
    }
    if err := session.RequestPty("vt100", 0, 0, modes); err != nil {
        panic(err)
    }
    sshShell, err := session.StdinPipe()
    if err != nil {
        panic(err)
    }
    if err :=session.Shell(); err !=nil {
        panic(err)
    }
    commands := []string{
            "conf t\n",
            fmt.Sprintf("router rip %s\n",rip_name),
            fmt.Sprintf("no address-family ipv4 unicast\n"),
            "end\n",
    }
    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil
}

// Function to clean RIP Authentication
func CleaningRIPAuthenticationRedisMaxpaths(deviceIP string, interf string) error {
    serverAddress := deviceIP
    serverUsername := "admin"
    serverPassword := "tcs123"
    config := &ssh.ClientConfig{
        User: serverUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(serverPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    client, err := ssh.Dial("tcp", serverAddress+":22", config)
    if err != nil {
        panic(err)
    }
    defer client.Close()
    session, err := client.NewSession()
    if err !=nil {
        panic(err)
    }
    defer session.Close()
    modes := ssh.TerminalModes{
            ssh.ECHO: 0,
            ssh.TTY_OP_ISPEED: 14400,
            ssh.TTY_OP_OSPEED: 14400,
    }
    if err := session.RequestPty("vt100", 0, 0, modes); err != nil {
        panic(err)
    }
    sshShell, err := session.StdinPipe()
    if err != nil {
        panic(err)
    }
    if err :=session.Shell(); err !=nil {
        panic(err)
    }
    commands := []string{
            "conf t\n",
            fmt.Sprintf("no key chain rip\n"),
            fmt.Sprintf("int %s\n",interf),
            fmt.Sprintf("no ip rip authentication key-chain rip\n"),
            fmt.Sprintf("no ip rip authentication mode md5\n"),
            "end\n",
    }
    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil
}

// ------------------------------------------

// Testcase to Call all CLEANUP Functions

func TestCleanup(t *testing.T) {

    err := CleaningRIPpassive("10.133.35.143","Ethernet1/11")
    if err != nil {
        t.Errorf("Failed to enable Rip Passice on device: %v", err)
    }

    err1 := CleaningRIPDistanceTimers("10.133.35.143","Ethernet1/11", "riptemp")
    if err1 != nil {
        t.Errorf("Failed to clean Rip on device: %v", err1)
    }

    err2 := CleaningRIPAuthenticationRedisMaxpaths("10.133.35.143","Ethernet1/11")
    if err2 != nil {
        t.Errorf("Failed to enable Rip Authentication on device: %v", err2)
    }

}


func DisableRiponDevices(deviceIP string, interf string, ripName string) error {
    serverAddress := deviceIP
    serverUsername := "admin"
    serverPassword := "tcs123"
    intf := interf
    rip_name := ripName

    fmt.Printf("\nTestcase Started: Disable Rip -----\n")

    config := &ssh.ClientConfig{
        User: serverUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(serverPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    client, err := ssh.Dial("tcp", serverAddress+":22", config)
    if err != nil {
        panic(err)
    }
    defer client.Close()

    // unconfigure
    rip_lib.Unconfigure_rip_on_device(serverAddress, rip_name)
    rip_lib.Unconfigure_rip_on_interface(serverAddress, intf, rip_name)

    rip_lib.Disable_rip(serverAddress)

    fmt.Printf("RIP disabled on device %s %s\n", serverAddress,intf)

    fmt.Printf("Testcase Ended -----\n\n")
    return nil
}

func TestDisableRip(t *testing.T) {
//    err2 := DisableRiponDevices("10.133.35.148","Ethernet1/11", "riptemp")
    err3 := DisableRiponDevices("10.133.35.143","Ethernet1/11", "riptemp")

//    if err2 != nil {
//        t.Errorf("Failed to disable RIP on device: %v", err2)
//    }
    if err3 != nil {
        t.Errorf("Failed to disable RIP on device: %v", err3)
    }
}


//-------------------------------------------------------------------------


func TestMain(m *testing.M) {
	
outputFilePath:"/home/tcs/sample/ondatra/debug/rip/test_output.txt"

File, erros. Create(outputFilePath)

if err != nil {

fmt.Printf("Failed to create file: %s: %v\n", outputFilePath, err) 
os.Exit(1)

defer File.Close()

origStdout: os. Stdout

os. Stdout File

defer func() (os. Stdout origStdout)()
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
excelFilePath := "/home/tcs/sample/ondatra/debug/rip/excel_summary.xlsx"
testcaseFilePath := "/home/tcs/sample/ondatra/debug/rip/valid_test.go"

// Save the Excel file at the specified path
if err := file.Save(excelFilePath); err != nil {
fmt.Printf("Failed to save Excel file at %s: %v\n", excelFilePath, err)
os.Exit(1)
}


// Print the paths for the testcase file and the Excel sheet file
fmt.Printf("Path for testcase file: %s\n", testcaseFilePath)
fmt.Printf("Path for Excel: %s\n", excelFilePath)

// Exit with the appropriate status code
os.Exit(exitVal)
}
