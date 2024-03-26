package acl
import "acl/acl_lib"
import (
    "fmt"
    "golang.org/x/crypto/ssh"
    "log"
    "testing"
    "os"
    "time"
    "regexp"
    "strings"
)
var testResults = make(map[string]string)

// ------------Configure Interfaces Testcase ------------------
func Configure_Interfaces_testcase(deviceIP string, intf1 string, intf1_ip string, intf2 string, intf2_ip string) error {
    router1Address := deviceIP
    routerUsername := "admin"
    routerPassword := "tcs123"
    interf1 := intf1
    interf1_ip := intf1_ip
    interf2 := intf2
    interf2_ip := intf2_ip

    executeCommand := "show running-config"
    fmt.Printf("Testcase: Executing interface configurations------\n")
    config := &ssh.ClientConfig{
        User: routerUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(routerPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    client, err := ssh.Dial("tcp", router1Address+":22", config)
    if err != nil {
        panic(err)
    }
    defer client.Close()

    session, err := client.NewSession()
    if err != nil {
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
    if err := session.Shell(); err != nil {
        panic(err)
    }

    commands := []string{
            "conf t\n",
            fmt.Sprintf("interface %s\n", interf1),
            fmt.Sprintf("no switchport\n"),
            fmt.Sprintf("no ip address\n"),
            fmt.Sprintf("ip address %s\n", interf1_ip),
            fmt.Sprintf("no shutdown\n"),
            fmt.Sprintf("exit\n"),
            fmt.Sprintf("interface %s\n", interf2),
            fmt.Sprintf("no switchport\n"),
            fmt.Sprintf("no ip address\n"),
            fmt.Sprintf("ip address %s\n", interf2_ip),
            fmt.Sprintf("no shutdown\n"),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }

    //  New Session for command
    router1Config := &ssh.ClientConfig{
        User: routerUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(routerPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    router1Client, err := ssh.Dial("tcp", router1Address+":22", router1Config)
    if err != nil {
        log.Fatalf("Failed to connect to Router 1: %v", err)
    }
    defer router1Client.Close()

    router1Session, err := router1Client.NewSession()
    if err != nil {
        log.Fatalf("Failed to create session on Router 1: %v", err)

    }
    defer router1Session.Close()

    //  Output
    router1Output, err := router1Session.CombinedOutput(executeCommand)
    if err != nil {
        fmt.Printf("Failed to execute command on Router 1: %v", err)
        return fmt.Errorf("failed to execute command on Router 1: %w",err)
    }
    fmt.Printf("Router 1 - Output:\n%s\n", router1Output)

    fmt.Print("Execution ended - Interface configurations testcase------\n\n")

    return nil
}


// Test function
func Test_Calling_configure_interfaces_testcases(t *testing.T) {
    testCaseName := "Configure Interfaces"
    err := Configure_Interfaces_testcase("10.133.35.148","Ethernet 1/1", "192.168.10.1/24","Ethernet 1/5","192.168.5.1/24")
    if err != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName, err)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
}


// Assuming you have more test functions similar to TestConfigureInterfaces...

// TestSummary prints a summary of all test cases execute
}

// ----------- Testcase Ended ------------


// ----------- Testcase 1: ACL Configuration ------------------
func Configure_acl_testcase(deviceIP string, acl_name string, rulee string) error {
    router1Address := deviceIP
    routerUsername := "admin"
    routerPassword := "tcs123"
    aclName := acl_name
    rule := rulee
    executeCommand := "show ip access-lists"
    fmt.Printf("Testcase 1: Executing Basic ACL configuration\n")

    config := &ssh.ClientConfig{
        User: routerUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(routerPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    client, err := ssh.Dial("tcp", router1Address+":22", config)
    if err != nil {
        panic(err)
    }
    defer client.Close()

    session, err := client.NewSession()
    if err != nil {
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
    if err := session.Shell(); err != nil {
        panic(err)
    }

    //  ACL Configuration
    commands := []string{
            "config\n",
            fmt.Sprintf("ip access-list %s\n", aclName),
            fmt.Sprintf("%s\n", rule),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }

    //  New Session for command
    router1Config := &ssh.ClientConfig{
        User: routerUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(routerPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    router1Client, err := ssh.Dial("tcp", router1Address+":22", router1Config)
    if err != nil {
        log.Fatalf("Failed to connect to Router 1: %v", err)
    }
    defer router1Client.Close()

    router1Session, err := router1Client.NewSession()
    if err != nil {
        log.Fatalf("Failed to create session on Router 1: %v", err)

    }
    defer router1Session.Close()

    //  Output
    router1Output, err := router1Session.CombinedOutput(executeCommand)
    if err != nil {
        fmt.Printf("Failed to execute command on Router 1: %v", err)
        return fmt.Errorf("failed to execute command on Router 1: %w",err)
    }
    fmt.Printf("Router 1 - Output:\n%s\n", router1Output)

    //  Validation
    if strings.Contains(string(router1Output), aclName) {
        fmt.Printf("Success: ACL Configured\n")
    }
    fmt.Print("Execution ended - ACL configuration testcase --------\n\n")

    return nil
}


func Test_Calling_configure_acl(t *testing.T) {
    testCaseName := "Configure acl testcase"
    err := Configure_acl_testcase("10.133.35.148", "acltemp", "5 permit icmp any any")
    if err != nil {
        t.Errorf("Failed to unconfigure ACL on device: %v", err)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
}
}
// ----------- Testcase Ended ----------------}

// ----------- Testcase 2: ACL Unconfiguration ------------------
func Unconfigure_acl_testcase(deviceIP string, acl_name string) error {
    router1Address := deviceIP
    routerUsername := "admin"
    routerPassword := "tcs123"
    aclName := acl_name
    executeCommand := "show ip access-lists"

    fmt.Printf("Testcase 2: Executing Basic ACL Unconfiguration-----------\n")

    config := &ssh.ClientConfig{
        User: routerUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(routerPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    client, err := ssh.Dial("tcp", router1Address+":22", config)
    if err != nil {
        panic(err)
    }
    defer client.Close()

    session, err := client.NewSession()
    if err != nil {
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
    if err := session.Shell(); err != nil {
        panic(err)
    }

    //  ACL Configuration
    commands := []string{
            "config\n",
            fmt.Sprintf("no ip access-list %s\n", aclName),
//            fmt.Sprintf("%s\n", rule),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }

    //  New Session for command
    router1Config := &ssh.ClientConfig{
        User: routerUsername,
        Auth: []ssh.AuthMethod{
            ssh.Password(routerPassword),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    router1Client, err := ssh.Dial("tcp", router1Address+":22", router1Config)
    if err != nil {
        log.Fatalf("Failed to connect to Router 1: %v", err)
    }
    defer router1Client.Close()

    router1Session, err := router1Client.NewSession()
    if err != nil {
        log.Fatalf("Failed to create session on Router 1: %v", err)

    }
    defer router1Session.Close()

    //  Output
    router1Output, err := router1Session.CombinedOutput(executeCommand)
    if err != nil {
        fmt.Printf("Failed to execute command on Router 1: %v", err)
        return fmt.Errorf("failed to execute command on Router 1: %w",err)
    }

    fmt.Printf("Router 1 - Output:\n%s\n", router1Output)

    //  Validation
    if strings.Contains(string(router1Output), aclName) {
        fmt.Printf("Failed: ACL NOT Unconfigured\n")
    } else {
	fmt.Printf("Success: ACL unconfigured")
    }
    fmt.Print("Execution ended - ACL configuration testcase----------\n\n")

    return nil
}


func Test_Calling_unconfigure_acl(t *testing.T) {
   testCaseName := "Acl Unconfiguration"
    err := Unconfigure_acl_testcase("10.133.35.148", "acltemp")
    if err != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName, err)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
}
}

// ----------- Testcase Ended ----------------


// ----------- Testcase 3: ACL Permit ------------------
func ACLPermitonDevices(ip string, acl_namee string, rulee string, boundd string, interf string, ping_ip string) error {
    acl_name := acl_namee
    rule := rulee
    bound := boundd
    intf := interf
    serverAddress := ip
    serverUsername := "admin"
    serverPassword := "tcs123"
    server2_address := ping_ip
    executeCommand := "ping "+string(server2_address)

    fmt.Printf("Testcase 3: Executing ACL Permit -------\n")

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

    // configuration
    acl_lib.Configure_acl_on_device(serverAddress, acl_name, rule)
    acl_lib.Configure_acl_on_interface(serverAddress, acl_name, intf, bound)

    fmt.Printf("ACL %s enabled on device %s %s via %s\n", acl_name, serverAddress, intf, bound)

    // new session for ping
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

    output1, err := session.CombinedOutput(string(executeCommand))
    if err != nil {
        log.Fatal(err)
        return fmt.Errorf("failed to execute command on Router 1: %w",err)
    }
    fmt.Printf("\nPing Statistics------\n%s\n", output1)
/*
    // validation
    match1, err := regexp.MatchString(string(acl_name), string(fetching_acls))
    if match1 == true {
        fmt.Printf("acl %s  is configured\n", acl_name)
        fmt.Printf("Failed: acl %s is present and configured on interface %s\n", acl_name, intf)
    } else {
        fmt.Printf("acl %s  is NOT configured\n", acl_name)       //main condition for this testcase
        fmt.Printf("Success: acl %s is NOT present and NOT configured on interface %s\n", acl_name, intf)
    }
*/
    // Cleanup
    acl_lib.Unconfigure_acl_on_device(serverAddress, acl_name)
    acl_lib.Unconfigure_acl_on_interface(serverAddress, acl_name, intf, bound)

    fmt.Print("Execution ended - ACL Permit testcase --------\n\n")

    return nil

}

func TestAclPermitDevices(t *testing.T){
    testCaseName := "Acl Permit On DEvices"
    err := ACLPermitonDevices("10.133.35.150", "acltemp", "10 permit icmp 192.168.10.2/24 any", "in", "ethernet 1/1", "192.168.10.1")
     if err != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName, err)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
}
}


// ----------- Testcase Ended ----------------



// ----------- Testcase 4: ACL Deny Testcase ------------------
func ACLDenyonDevices(ip string, acl_namee string, rulee string, boundd string, interf string, ping_ip string) error {
    acl_name := acl_namee
    rule := rulee
    bound := boundd
    intf := interf
    serverAddress := ip
    serverUsername := "admin"
    serverPassword := "tcs123"
    server2_address := ping_ip
    executeCommand := "ping "+string(server2_address)

    fmt.Printf("\nTestcase 4: Executing ACL Deny ---------\n")

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

    // configuration
    acl_lib.Configure_acl_on_device(serverAddress, acl_name, rule)
    acl_lib.Configure_acl_on_interface(serverAddress, acl_name, intf, bound)

    fmt.Printf("ACL enabled on device %s %s via %s\n", serverAddress, intf, bound)

    // new session for ping
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

    output1, err := session.CombinedOutput(string(executeCommand))
    if err != nil {
        log.Fatal(err)
        return fmt.Errorf("failed to execute command on Router 1: %w",err)
    }
    fmt.Printf("Ping Statistics------\n %s\n", output1)

    // Cleanup
    acl_lib.Unconfigure_acl_on_device(serverAddress, acl_name)
    acl_lib.Unconfigure_acl_on_interface(serverAddress, acl_name, intf, bound)

    fmt.Print("Execution ended - ACL Deny testcase -------------\n\n")

    return nil

}

func TestAclDenyDevices(t *testing.T){
    testCaseName := "Acl deny on devices"
    err1 := ACLDenyonDevices("10.133.35.150", "acltemp", "10 deny icmp 192.168.5.1/24 any", "in", "Ethernet1/5", "192.168.5.2")
     if err1 != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName, err1)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
}
}

    

// ----------- Testcase Ended ----------------


// ----------- ACL with Multiple Rules Testcase ------------------
func ACLWithMultipleRules(ip string, acl_namee string, rulee string, boundd string, interf string, ping_ip string, message string) error {
    acl_name := acl_namee
    rule := rulee
    bound := boundd
    intf := interf
    msg := message
    serverAddress := ip
    serverUsername := "admin"
    serverPassword := "tcs123"
    server2Address := ping_ip
    executeCommand := "ping "+string(server2Address)

    fmt.Printf("\nTestcase 5: Executing ACL with Multiple Rules -------\n")
    fmt.Printf("%s\n", msg)

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
    acl_lib.Configure_acl_on_device(serverAddress, acl_name, rule)
    acl_lib.Configure_acl_on_interface(serverAddress, acl_name, intf, bound)

    fmt.Printf("ACL %s enabled on device %s %s via %s\n", acl_name, serverAddress, intf, bound)

    //New Session for Ping
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

    output1, err := session.CombinedOutput(string(executeCommand))
    if err != nil {
        log.Fatal(err)
        return fmt.Errorf("failed to execute command on Router 1: %w",err)
    }
    fmt.Printf("Ping Statistics ------\n %s\n", output1)

    // Cleanup
    acl_lib.Unconfigure_acl_on_device(serverAddress, acl_name)
    acl_lib.Unconfigure_acl_on_interface(serverAddress, acl_name, intf, bound)

    fmt.Print("Execution ended - ACL with Multiple Rules Testcase\n\n")

    return nil
}

func TestAclMultipleRulesDevices(t *testing.T){
    testCaseName := "Acl with Multiple rules-device1"
  
    err := ACLWithMultipleRules("10.133.35.148", "acltemp", "permit icmp 192.168.10.1/24 any", "in", "Ethernet1/1", "192.168.10.2", "Part 1: Permitting")
     if err != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName, err)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
}

    testCaseName1 := "Acl with Multiple rules-device2"
    err1 := ACLWithMultipleRules("10.133.35.150", "acltemp","deny icmp 192.168.10.2/24 any", "in", "Ethernet1/1", "192.168.10.1", "Part 2: Denying")
     if err1 != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName1, err1)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
}

    testCaseName2 := "Acl with Multiple rules-device3"

    err2 := ACLWithMultipleRules("10.133.35.152", "acltemp", "deny icmp 192.168.5.2/24 any", "in", "Ethernet1/5", "192.168.5.1", "Part 3: Denying")
         if err2 != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName2, err2)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
}
}

// ----------- Testcase Ended ----------------


// ----------- ACL - In & Out Bounds Testcase --------------
func ACLinoutBound(ip string, acl_namee string, rulee string, boundd string, interf string, ping_ip string, message string) error {
    acl_name := acl_namee
    rule := rulee
    bound := boundd
    intf := interf
    msg := message
    serverAddress := ip
    serverUsername := "admin"
    serverPassword := "tcs123"
    server2Address := ping_ip
    executeCommand := "ping "+string(server2Address)

    fmt.Printf("\nTestcase 6: Executing ACL with Multiple Rules -------\n")
    fmt.Printf("%s\n", msg)

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
    acl_lib.Configure_acl_on_device(serverAddress, acl_name, rule)
    acl_lib.Configure_acl_on_interface(serverAddress, acl_name, intf, bound)

    fmt.Printf("ACL %s enabled on device %s %s via %s\n", acl_name, serverAddress, intf, bound)

    //New Session for Ping
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

    output1, err := session.CombinedOutput(string(executeCommand))
    if err != nil {
        log.Fatal(err)
        return fmt.Errorf("failed to execute command on Router 1: %w",err)
    }
    fmt.Printf("Ping Statistics ------\n %s\n", output1)

    // Cleanup
    acl_lib.Unconfigure_acl_on_device(serverAddress, acl_name)
    acl_lib.Unconfigure_acl_on_interface(serverAddress, acl_name, intf, bound)

    fmt.Print("Execution ended - ACL with Multiple Rules Testcase----------\n\n")

    return nil
}

func TestACLinoutBound(t *testing.T){
    testCaseName := "ACL In - Out Bound - device 1"
    err := ACLinoutBound("10.133.35.148", "acltemp", "permit icmp 192.168.10.1/24 any", "in", "Ethernet1/1", "192.168.10.2", "In Bound")
    if err != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName, err)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
}

    testCaseName1 := "ACL In - Out Bound - device 2"
    err1 := ACLinoutBound("10.133.35.150", "acltemp","permit icmp 192.168.10.2/24 any", "out", "Ethernet1/1", "192.168.10.1", "Out Bound")
    if err1 != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName1, err1)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"

}
}
// ----------- Testcase Ended ----------------



// ----------- Checking acl configured on device Testcase ------------------
func Acl_configured_on_router(ip string, acl_namee string, rulee string) error {
    acl_name := acl_namee
    rule := rulee

    serverAddress := ip
    serverUsername := "admin"
    serverPassword := "tcs123"
    executeCommand1 := "show ip access-list"

    fmt.Printf("\nTestcase 7: Checking acl configured on device ------\n")

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

    // Configure
    acl_lib.Configure_acl_on_device(serverAddress, acl_name, rule)

    // New Session
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

    // executing command
    fetching_acls, err := session.CombinedOutput(string(executeCommand1))
    if err != nil {
        log.Fatal(err)
        return fmt.Errorf("failed to execute command on Router 1: %w",err)
    }
    // validating
    match1, err := regexp.MatchString(string(acl_name), string(fetching_acls))
    if match1 == true {
        fmt.Printf("Success: acl %s  is present and configured on router\n", acl_name)
    } else {
        fmt.Printf("Failed: acl %s is NOT present and configured on router\n", acl_name)
    }

    // Cleanup
    acl_lib.Unconfigure_acl_on_device(serverAddress, acl_name)
    fmt.Printf("\nTestcase : Checking acl configured on device - Ended ---------\n\n")

    return nil
}

func Test_Acl_configured_on_router(t *testing.T){
    testCaseName := "ACL Configured on router"
    err := Acl_configured_on_router("10.133.35.148","acltemp", "10 permit icmp any any")
     if err != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName, err)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
}
}

// ----------- Testcase Ended ----------------


// ----------- Checking acl not configured on Device Testcase ------------------
func Acl_not_configured_on_router(ip string, acl_namee string) error {
    acl_name := acl_namee

    serverAddress := ip
    serverUsername := "admin"
    serverPassword := "tcs123"
    executeCommand1 := "show ip access-list"

    fmt.Printf("\nTestcase 8: Checking acl - not configured on device ---------\n")

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

    // New Session
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

    // executing command
    fetching_acls, err := session.CombinedOutput(string(executeCommand1))
    if err != nil {
        log.Fatal(err)
        return fmt.Errorf("failed to execute command on Router 1: %w",err)
    }
    // validating
    match1, err := regexp.MatchString(string(acl_name), string(fetching_acls))
    if match1 == true {
        fmt.Printf("acl %s  is configured\n", acl_name)
        fmt.Printf("Failed: acl is present and configured\n")
    } else {
        fmt.Printf("acl %s  is NOT configured\n", acl_name)       //main condition for this testcase
        fmt.Printf("Success: acl is NOT present and NOT configured\n")
    }

    // Cleanup
    acl_lib.Unconfigure_acl_on_device(serverAddress, acl_name)
 
    fmt.Printf("Testcase - Checking acl - not configured on device - Ended --------\n\n")

    return nil
}

func Test_Acl_not_configured_on_router(t *testing.T){
     testCaseName := "ACL not configured on router"
     err := Acl_not_configured_on_router("10.133.35.148","acltemp")
     if err != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName, err)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
}
}


// ----------- Testcase Ended ----------------


// ----------- Checking acl configured on Interface Testcase ------------------
func Acl_configured_on_interface(ip string, acl_namee string, rulee string, boundd string, interf string) error {
    acl_name := acl_namee
    rule := rulee
    intf := interf
    bound := boundd

    serverAddress := ip
    serverUsername := "admin"
    serverPassword := "tcs123"
    executeCommand := "show running-config interface "+string(intf)

    fmt.Printf("\nTestcase 9: Checking acl configured on interface ------\n\n")

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

    // Configure
    acl_lib.Configure_acl_on_device(serverAddress, acl_name, rule)
    acl_lib.Configure_acl_on_interface(serverAddress, acl_name, intf, bound)

    // New Session
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

    // executing command
    fetching_acls, err := session.CombinedOutput(string(executeCommand))
    if err != nil {
        log.Fatal(err)
        return fmt.Errorf("failed to execute command on Router 1: %w",err)
    }
    // validating
    match1, err := regexp.MatchString(string(acl_name), string(fetching_acls))
    if match1 == true {
        fmt.Printf("Success: acl %s is present and configured on interface %s\n", acl_name, intf)
    } else {
        fmt.Printf("Failed: acl %s is NOT present and configured on interface %s\n", acl_name, intf)
    }

    // Cleanup
    acl_lib.Unconfigure_acl_on_device(serverAddress, acl_name)
    acl_lib.Unconfigure_acl_on_interface(serverAddress, acl_name, intf, bound)

    fmt.Printf("\nTestcase : Checking acl configured on interface - Ended\n")

    return nil
}

func Test_Acl_configured_on_interface(t *testing.T){
    testCaseName := "ACL configured on interface"
    err := Acl_configured_on_interface("10.133.35.148","acltag", "10 permit icmp any any", "out", "ethernet 1/1")
    if err != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName, err)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
}
}

// ----------- Testcase Ended ----------------


// ----------- Checking acl not configured on Interface Testcase ------------------
func Acl_not_configured_on_interface(ip string, acl_namee string, rulee string, interf string) error {
    acl_name := acl_namee
    rule := rulee
    intf := interf

    serverAddress := ip
    serverUsername := "admin"
    serverPassword := "tcs123"
    executeCommand := "show running-config interface "+string(intf)

    fmt.Printf("\nTestcase 10: Checking acl not configured on interface -------\n")

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

    // New Session
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

    // configure
    acl_lib.Configure_acl_on_device(serverAddress, acl_name, rule)

    // executing command
    fetching_acls, err := session.CombinedOutput(string(executeCommand))
    if err != nil {
        log.Fatal(err)
        return fmt.Errorf("failed to execute command on Router 1: %w",err)
    }
    // validating
    match1, err := regexp.MatchString(string(acl_name), string(fetching_acls))
    if match1 == true {
        fmt.Printf("acl %s  is configured\n", acl_name)
        fmt.Printf("Failed: acl %s is present and configured on interface %s\n", acl_name, intf)
    } else {
        fmt.Printf("acl %s  is NOT configured\n", acl_name)       //main condition for this testcase
        fmt.Printf("Success: acl %s is NOT present and NOT configured on interface %s\n", acl_name, intf)
    }

    // Cleanup
    acl_lib.Unconfigure_acl_on_device(serverAddress, acl_name)

    fmt.Printf("\nTestcase : Checking acl configured on interface - Ended --------\n\n")

    return nil
}

func Test_Acl_not_configured_on_interface(t *testing.T){
    testCaseName := "Acl not configured on interface"
    err := Acl_not_configured_on_interface("10.133.35.148","acltemp", "10 permit icmp any any", "ethernet 1/1")
    if err != nil {
        t.Errorf("Test '%s' failed: %v",testCaseName, err)
        testResults[testCaseName] = "FAILED"
    } else {
        testResults[testCaseName] = "PASSED"
}
}

// ----------- Testcase Ended ----------------


/*func TestMain(m *testing.M) {
    exitVal := m.Run()
    currentTime := time.Now()
    fmt.Println("-----------------------------------------------------------------------------------")
    fmt.Printf("%s INFO |                  Task Result Summary                   |\n", currentTime.Format("2006-01-02 15:04:05"))
    fmt.Println("-----------------------------------------------------------------------------------")                         
    for testName, result := range testResults {
        fmt.Printf("%s INFO: %-50s %s\n", currentTime.Format("2006-01-02 15:04:05"), testName+":", result)
    }
    fmt.Println("-----------------------------------------------------------------------------------")
    os.Exit(exitVal)
}*/
func TestMain(m *testing.M) {
    exitVal := m.Run()
 
    // Prepare to summarize test results
    currentTime := time.Now()
   // passedCount, failedCount, erroredCount, blockedCount := 0, 0, 0, 0
    passedCount, failedCount := 0, 0
    // Tally passed, failed, errored, and blocked tests
    for _, result := range testResults {
        switch result {
        case "PASSED":
            passedCount++
        case "FAILED":
            failedCount++
        /*case "ERRORED":
            erroredCount++
        case "BLOCKED":
            blockedCount++
        */}
    }
    TotalTests := passedCount + failedCount
    // Print the summary header
    fmt.Println("--------------------------------------------------------------------------------------")
    fmt.Printf("%s INFO |                     Task Result Summary                   |\n", currentTime.Format("2006-01-02 15:04:05"))
    fmt.Println("--------------------------------------------------------------------------------------")
    // Print counts for passed, failed, errored, and blocked tests with headings
    /*fmt.Printf("%-50s %14s\n", "Name", "No. of Tests")
    fmt.Println("--------------------------------------------------------------------------------------")
    fmt.Printf("%-50s %10d\n", "Passed", passedCount)
    fmt.Printf("%-50s %10d\n", "Failed", failedCount)
    //fmt.Printf("%-50s %10d\n", "Errored", erroredCount)
    //fmt.Printf("%-50s %10d\n", "Blocked", blockedCount)
    fmt.Println("--------------------------------------------------------------------------------------")
    */
    fmt.Printf ("%s INFO: %-30s %8s\n", currentTime.Format("2006-01-02 15:04:05"), "Metric", "Count")
    fmt.Println("--------------------------------------------------------------------------------------")
    fmt.Printf("%s INFO: %-30s %8d\n", currentTime.Format("2006-01-02 15:04:05"), "Passed", passedCount)
    fmt.Printf("%s INFO: %-30s %8d\n", currentTime.Format("2006-01-02 15:04:05"), "Failed", failedCount)
    fmt.Println("--------------------------------------------------------------------------------------")
    fmt.Printf("%s INFO: %-30s %8d\n", currentTime.Format("2006-01-02 15:04:05"), "TOTAL", TotalTests)
    fmt.Println("--------------------------------------------------------------------------------------")

    // Print the task summary
    fmt.Printf("%s INFO |                         Task Summary                      |\n", currentTime.Format("2006-01-02 15:04:05"))
    fmt.Println("--------------------------------------------------------------------------------------")
    for testName, result := range testResults {
        fmt.Printf("%s INFO: %-50s %s\n", currentTime.Format("2006-01-02 15:04:05"), testName+":", result)
    }
    fmt.Println("--------------------------------------------------------------------------------------")
 
    // Exit with the appropriate status code
    os.Exit(exitVal)
}



--------------------------------------------------------------------------------------
2024-03-26 10:57:38 INFO |                     Task Result Summary                   |
--------------------------------------------------------------------------------------
2024-03-26 10:57:38 INFO: Metric                            Count
--------------------------------------------------------------------------------------
2024-03-26 10:57:38 INFO: Passed                               11
2024-03-26 10:57:38 INFO: Failed                                0
--------------------------------------------------------------------------------------
2024-03-26 10:57:38 INFO: TOTAL                                11
--------------------------------------------------------------------------------------
2024-03-26 10:57:38 INFO |                         Task Summary                      |
--------------------------------------------------------------------------------------
2024-03-26 10:57:38 INFO: Acl Permit On DEvices:                             PASSED
2024-03-26 10:57:38 INFO: Acl with Multiple rules-device1:                   PASSED
2024-03-26 10:57:38 INFO: ACL In - Out Bound - device 1:                     PASSED
2024-03-26 10:57:38 INFO: ACL Configured on router:                          PASSED
2024-03-26 10:57:38 INFO: ACL configured on interface:                       PASSED
2024-03-26 10:57:38 INFO: Configure Interfaces:                              PASSED
2024-03-26 10:57:38 INFO: Configure acl testcase:                            PASSED
2024-03-26 10:57:38 INFO: Acl Unconfiguration:                               PASSED
2024-03-26 10:57:38 INFO: Acl not configured on interface:                   PASSED
2024-03-26 10:57:38 INFO: Acl deny on devices:                               PASSED
2024-03-26 10:57:38 INFO: ACL not configured on router:                      PASSED
--------------------------------------------------------------------------------------
ok  	command-line-arguments	192.569s

OUTPUT:

