package detail

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"proxy/base"
	"regexp"
	"strconv"
	"strings"
)

type _66ip struct {
	addressVector
	ipRegex *regexp.Regexp
}

func init() {
	p := new(_66ip)
	p.addressVector.init()
	p.ipRegex, _ = regexp.Compile("\\d{1,3}.\\d{1,3}.\\d{1,3}.\\d{1,3}:\\d{1,5}")
	p.Register(base.GlobalProxyManger)
}

func (p *_66ip) run(url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Add("User-Agent", base.GlobalUserAgent.Next())
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8")
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}
	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		data := p.ipRegex.FindAll([]byte(s.Text()), -1)
		address := make([]*base.Address, 0, len(data))
		for _, e := range data {
			ipPort := strings.Split(string(e), ":")
			if len(ipPort) == 2 {
				ipSplit := strings.Split(string(ipPort[0]), ".")
				if len(ipSplit) != 4 {
					continue
				}
				for i := 0; i < 4; i++ {
					num, err := strconv.Atoi(ipSplit[i])
					if err != nil || num <= 0 || num > 255 {
						continue
					}
				}

				num, err := strconv.Atoi(ipPort[1])
				if err != nil || num <= 0 || num > 65535 {
					continue
				}
				addr := new(base.Address)
				addr.IP = ipPort[0]
				addr.Port = num
				addr.Level = base.Unknown
				address = append(address, addr)
			}
		}
		p.append(address)
	})
}

func (p *_66ip) Run() error {
	go p.run("http://www.66ip.cn/mo.php?sxb=&tqsl={}&port=&export=&ktip=&sxa=&submit=%CC%E1++%C8%A1&textarea=")
	p.run("http://www.66ip.cn/nmtq.php?getnum={}&isp=0&anonymoustype=0&s")
	return nil
}

func (p *_66ip) AddressList() []*base.Address {
	return p.extract()
}

func (p *_66ip) Register(m *base.Manager) {
	m.Register("66ip", p)
}
