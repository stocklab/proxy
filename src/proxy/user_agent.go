package proxy

import (
	"bufio"
	"io"
	"math/rand"
	"os"
)

var GlobalUserAgent *UserAgent = nil

func init() {
	var err error
	GlobalUserAgent, err = NewUserAgent("conf/user_agent.conf")
	if err != nil {
		panic(err)
	}
}

type UserAgent struct {
	agents []string
	index int
}


func newUserAgent() *UserAgent {
	ua := &UserAgent{index : 0}
	ua.agents = make([]string, 0, 4)
	return ua
}

func NewUserAgent(config string) (*UserAgent, error) {
	file, err := os.Open(config)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ua := newUserAgent()
	br := bufio.NewReader(file)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		ua.agents = append(ua.agents, string(line))
	}
	return ua, nil
}

func (ua *UserAgent) Next() (agent string) {
	if ua.index >= len(ua.agents) {
		ua.index = 0
	}
	agent = ua.agents[ua.index]
	ua.index++
	return
}

func (ua *UserAgent) Random() string {
	return ua.agents[rand.Intn(len(ua.agents))]
}