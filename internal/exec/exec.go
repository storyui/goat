package exec

import (
	"fmt"
	"os"
	"time"

	"github.com/pterm/pterm"
	"github.com/storyui/goat/config"
	"golang.org/x/crypto/ssh"
	"golang.org/x/exp/rand"
)

type Exec struct{}

func NewExec() *Exec {
	return &Exec{}
}

func (d *Exec) Run(environment string, goatConfig config.GoatConfig) {
	serverConfig := goatConfig.Servers[0][environment]

	config := &ssh.ClientConfig{
		User: serverConfig.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(serverConfig.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	sshServer := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	client, err := ssh.Dial("tcp", sshServer, config)

	if err != nil {
		pterm.FgRed.Printf("Dial error: %v\n", err)
		return
	}

	defer client.Close()

	if serverConfig.Exec != nil {
		if len(serverConfig.Exec.Steps) > 0 {
			for _, step := range serverConfig.Exec.Steps {
				session, err := client.NewSession()
				if err != nil {
					pterm.FgRed.Printf("New session error: %v\n", err)
					return
				}
				pterm.FgBlue.Printf("step: %s\n", step.Name)

				session.Stdout = os.Stdout
				session.Stderr = os.Stderr
				err = session.Run(step.Run)
				if err != nil {
					pterm.FgRed.Printf("Server run exec error: %v\n", err)
					return
				}

				if step.Delay != 0 {
					time.Sleep(time.Duration(rand.Int31n(step.Delay)) * time.Millisecond)
				}

				session.Close()
			}
		}
	}

}
