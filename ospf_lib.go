package ospf_lib
import (
    "fmt"
    "golang.org/x/crypto/ssh"
    "log"
//    "testing"
    "time"
    "regexp"
)

func Enable_ospf(deviceIP string) error {
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
            fmt.Sprintf("no feature ospf\n"),
            fmt.Sprintf("feature ospf\n"),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }

    return nil
}


// Disable ospf
func Disable_ospf(deviceIP string) error {
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
            fmt.Sprintf("no feature ospf\n"),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }

    return nil
}



// Configure ospf on router
func Configure_ospf_router(deviceIP string, ospf_name string) error {
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
            fmt.Sprintf("router ospf %s\n", ospf_name),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }

    return nil
}


// Unconfigure OSPF on Router
func Unconfigure_ospf_router(deviceIP string, ospf_name string) error {
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
            fmt.Sprintf("no router ospf %s\n", ospf_name),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }

    return nil
}


//  Configure_ospf_interface
func Configure_ospf_interface(deviceIP string, interf string, ospf_name string) error {
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
            fmt.Sprintf("interface %s\n", interf),
            fmt.Sprintf("ip router ospf %s area 0\n", ospf_name),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }

    return nil
}


func Unconfigure_ospf_interface(deviceIP string, interf string, ospf_name string) error {
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
            fmt.Sprintf("int %s\n", interf),
            fmt.Sprintf("no ip router ospf %s area 0\n", ospf_name),
	    fmt.Sprintf("exit\n"),
	    fmt.Sprintf("no feature ospf\n"),
            "end\n",
    }
    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil
}
func Print_Output(deviceIP string, OutputCommand string) error{
    serverUsername := "admin"
    serverPassword := "tcs123"
    serverAddress := deviceIP
    OutputCommand1 := OutputCommand

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

    router1Output, err := router1Session.CombinedOutput(OutputCommand1)
    if err != nil {
        fmt.Printf("Failed to execute ping command on Router 1: %v", err)
    }
    time.Sleep(30 * time.Second)
   // fmt.Printf("Output:\n%s\n", router1Output)
    fmt.Printf("Output ----:\n%s\n", router1Output)
    return nil
}
func Configure_ospf_multiarea(deviceIP string, intf string, ospf_name string) error{
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
        //int {intf}\n ip router ospf 2 area 1\n end
            "conf t\n",
            fmt.Sprintf("interface %s\n", intf),
            fmt.Sprintf("ip router ospf %s area 1\n", ospf_name),
            "end\n",
    }
    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil
}
func Configure_ospf_loopback(deviceIP string, intf string, ospf_name string) error{
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
        //int {intf}\n ip router ospf 2 area 2\n end'''
            "conf t\n",
            fmt.Sprintf("interface %s\n", intf),
            fmt.Sprintf("ip router ospf %s area 2\n", ospf_name),
            "end\n",
    }
    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil
}
func Unconfigure_ospf_multiarea(deviceIP string, intf string, ospf_name string) error{
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
            fmt.Sprintf("interface %s\n", intf),
            fmt.Sprintf("no ip router ospf %s area 1\n", ospf_name),
            fmt.Sprintf("no ip router ospf %s area 0\n", ospf_name),
            "end\n",
    }
    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil
}
func Unconfigure_ospf_loopback(deviceIP string, intf string, ospf_name string) error{
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
        //int {intf}\n ip router ospf 2 area 2\n end'''
            "conf t\n",
            fmt.Sprintf("interface %s\n", intf),
            fmt.Sprintf("no ip router ospf %s area 2\n", ospf_name),
            "end\n",
    }
    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil
}
func Configure_ospf_authentication(deviceIP string, intf string, ospf_name string) error{
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
        //int {intf}\n ip router ospf 2 area 2\n end'''
            "conf t\n",
            fmt.Sprintf("router ospf %s\n", ospf_name),
            fmt.Sprintf("area 0.0.0.0 authentication message-digest\n"),
            fmt.Sprintf("exit\n"),
            fmt.Sprintf("int %s", intf),
            fmt.Sprintf("ip ospf authentication-key 0 cisco\n"),
            fmt.Sprintf("ip ospf message-digest-key 21 md5 0 cisco\n"),
            "end\n",
    }
    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil
}
func Unconfigure_ospf_authentication(deviceIP string, intf string, ospf_name string) error{
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
        //int {intf}\n ip router ospf 2 area 2\n end'''
            "conf t\n",
            fmt.Sprintf("router ospf %s\n", ospf_name),
            fmt.Sprintf("no area 0.0.0.0 authentication message-digest\n"),
            fmt.Sprintf("exit\n"),
            fmt.Sprintf("int %s", intf),
            fmt.Sprintf("no ip ospf authentication-key 0 cisco\n"),
            fmt.Sprintf("no ip ospf message-digest-key 21 md5 0 cisco\n"),
            "end\n",
    }
    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil
}
func Configure_ospf_timers(deviceIP string, intf string) error{
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
    //int {intf}\n ip ospf hello-interval 30\n ip ospf dead-interval 100\n end'''.format(intf = intf))
    commands := []string{
            "conf t\n",
            fmt.Sprintf("interface %s\n", intf),
            fmt.Sprintf("ip ospf hello-interval 30\n"),
            fmt.Sprintf("ip ospf dead-interval 100\n"),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil
}
func Unconfigure_ospf_timers(deviceIP string, intf string) error{
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
    //int {intf}\n ip ospf hello-interval 30\n ip ospf dead-interval 100\n end'''.format(intf = intf))
    commands := []string{
            "conf t\n",
            fmt.Sprintf("interface %s\n", intf),
            fmt.Sprintf("no ip ospf hello-interval 30\n"),
            fmt.Sprintf("no ip ospf dead-interval 100\n"),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil
}
func Configure_ospf_passive_interface(deviceIP string, intf string) error {
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
            fmt.Sprintf("interface %s\n", intf),
            fmt.Sprintf("ip ospf passive-interface\n"),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil 
}
func Unconfigure_ospf_passive_interface(deviceIP string, intf string) error {
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
            fmt.Sprintf("interface %s\n", intf),
            fmt.Sprintf("no ip ospf passive-interface\n"),
	    fmt.Sprintf("exit\n"),
	    fmt.Sprintf("exit\n"),
	    fmt.Sprintf("no feature ospf\n"),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil 
}
func Configure_ospf_DR_BDR(deviceIP string, intf string) error {
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
    //int {intf}\n ip ospf priority 100\n end'''.format(intf = intf))
    commands := []string{
            "conf t\n",
            fmt.Sprintf("interface %s\n", intf),
            fmt.Sprintf("ip ospf priority 100\n"),
            fmt.Sprintf("no feature ospf\n"),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil
}
func Configure_ospf_clear_process(deviceIP string) error {
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
    //int {intf}\n ip ospf priority 100\n end'''.format(intf = intf))
    commands := []string{
            "conf t\n",
            fmt.Sprintf("clear ip ospf neighbor *\n"),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil
}
func Unconfigure_ospf_DR_BDR(deviceIP string, intf string) error {
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
        //int {intf}\n no ip ospf priority 100\n end'''
            "conf t\n",
            fmt.Sprintf("interface %s\n", intf),
            fmt.Sprintf("no ip ospf priority 100\n"),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil
}
func Configure_ospf_graceful_shutdown(deviceIP string, ospf_name string) error {
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
        //router ospf 2\n shutdown\n end
        "conf t\n",
         fmt.Sprintf("router ospf %s\n",ospf_name),
         fmt.Sprintf("shutdown\n"),
         "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil
}
func Unconfigure_ospf_graceful_shutdown(deviceIP string, ospf_name string) error {
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
        //router ospf 2\n shutdown\n end
        "conf t\n",
         fmt.Sprintf("router ospf %s\n",ospf_name),
         fmt.Sprintf("no shutdown\n"),
         "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }
    return nil
}
func Configure_ospf_vlinks(deviceIP string, ospf_name string)error {
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
    //router ospf 2\n area 0.0.0.2 virtual-link 4.4.4.4\n end''')
    commands := []string{
            "conf t\n",
            fmt.Sprintf("router ospf %s\n", ospf_name),
            fmt.Sprintf("area 0.0.0.2 virtual-link 192.168.29.1\n"),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }

    return nil
}
func Unconfigure_ospf_vlinks(deviceIP string, ospf_name string)error{
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
    //router ospf 2\n area 0.0.0.2 virtual-link 4.4.4.4\n end''')
    commands := []string{
            "conf t\n",
            fmt.Sprintf("router ospf %s\n", ospf_name),
            fmt.Sprintf("no area 0.0.0.2 virtual-link 192.168.29.1\n"),
            "end\n",
    }

    for _,cmd := range commands {
        sshShell.Write([]byte(cmd))
        time.Sleep(1 * time.Second)
    }

    return nil
}

func Validation(pattern string, data string) {
        match, _ := regexp.MatchString(pattern, data)
        if match {
                fmt.Println("Test case validation is successfully passed")
         
        } else {
                fmt.Println("Test case validation is failed")
        //      panic(match)
        }
}
